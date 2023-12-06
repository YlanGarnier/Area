package area

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/lenismtho/area/pkg/core/action"
	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/reaction"
	"github.com/lenismtho/area/pkg/protogen"
)

type dbArea interface {
	CreateArea(area models.Area) (uint, error)
	GetUserByToken(token string) (*models.User, error)
	GetAServiceAccessTokenByUserID(userID uint, service string) (string, error)
	DeleteAreaByID(areaID string) error
}

type Controller struct {
	db dbArea

	cHTTP     action.HTTPServiceAction
	cDiscord  action.DiscordServiceAction
	cEthereum action.EthereumServiceAction
	cSpotify  action.SpotifyServiceAction
	cGMAIL    action.GmailServiceAction
}

func NewController(
	db dbArea,
	httpClient protogen.HTTPServiceActionClient,
	discordClient protogen.DiscordServiceActionClient,
	ethereumClient protogen.EthereumServiceActionClient,
	spotifyClient protogen.SpotifyServiceActionClient,
	gmailClient protogen.GmailServiceActionClient,
) (*Controller, error) {
	return &Controller{
		db:        db,
		cHTTP:     action.HTTPServiceAction{HTTPServiceActionClient: httpClient},
		cDiscord:  action.DiscordServiceAction{DiscordServiceActionClient: discordClient},
		cEthereum: action.EthereumServiceAction{EthereumServiceActionClient: ethereumClient},
		cSpotify:  action.SpotifyServiceAction{SpotifyServiceActionClient: spotifyClient},
		cGMAIL:    action.GmailServiceAction{GmailServiceActionClient: gmailClient},
	}, nil
}

type CreateAreaRequest struct {
	Name   string `json:"name"`
	Action struct {
		Service string         `json:"service"`
		Route   string         `json:"route"`
		Params  map[string]any `json:"params"`
	} `json:"action"`
	Reaction struct {
		Service string `json:"service"`
		Route   string `json:"route"`
		Target  string `json:"target"`
	} `json:"reaction"`
}

// DeleteArea 		godoc
// @Summary			Delete an area.
// @Tags			area
// @Param			id	path	string	true	"Area ID"
// @Accept			json
// @Success			204	"no content"
// @Failure			400 "one or more parameters are missing or invalid"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/area/{id} [delete]
func (c *Controller) DeleteArea(ctx *gin.Context) {
	areaID := ctx.Param("id")
	if len(areaID) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "area id is empty",
		})
		return
	}
	err := c.db.DeleteAreaByID(areaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// CreateArea 		godoc
// @Summary			Create an area.
// @Tags			area
// @Param			CreateAreaRequest	body	CreateAreaRequest	true	"Fill requestr"
// @Accept			json
// @Success			204	"no content"
// @Failure			400 "one or more parameters are missing or invalid"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/area/new [post]
func (c *Controller) CreateArea(ctx *gin.Context) {
	var req CreateAreaRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(req.Name) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "name is empty",
		})
		return
	}
	headerTab := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(headerTab) == 0 || headerTab[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
	}
	token := headerTab[1]
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: get Action token

	reactionAccessToken, err := c.db.GetAServiceAccessTokenByUserID(user.ID, req.Reaction.Service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	base := &protogen.Base{
		Token:  reactionAccessToken,
		Target: req.Reaction.Target,
	}
	areaID, err := c.db.CreateArea(models.Area{
		Name:        req.Name,
		ActService:  req.Action.Service,
		RouteAction: req.Action.Route,
		ReaServices: req.Reaction.Service,
		Route:       req.Reaction.Route,
		Base:        base,
		UserID:      user.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	respType := reaction.GetRespType(reaction.ReaService(req.Reaction.Service), req.Reaction.Route)
	log.Println(respType)

	switch models.ActService(req.Action.Service) {
	case models.HttpService:
		err := c.cHTTP.ActionHTTPRouter(ctx, action.HTTPAct(req.Action.Route), req.Action.Params, respType, uint32(areaID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	case models.EthereumService:
		err := c.cEthereum.ActionEthereumRouter(ctx, action.EthereumAct(req.Action.Route), req.Action.Params, respType, uint32(areaID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	case models.DiscordService:
		discordAccessToken, err := c.db.GetAServiceAccessTokenByUserID(user.ID, req.Action.Service)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = c.cDiscord.ActionDiscordRouter(
			ctx,
			action.DiscordAct(req.Action.Route),
			req.Action.Params,
			respType,
			uint32(areaID),
			discordAccessToken,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	case models.SpotifyService:
		spotifyAccessToken, err := c.db.GetAServiceAccessTokenByUserID(user.ID, req.Action.Service)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = c.cSpotify.ActionSpotifyRouter(
			ctx,
			action.SpotifyAct(req.Action.Route),
			req.Action.Params,
			respType,
			uint32(areaID),
			spotifyAccessToken,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	case models.GmailService:
		gmailAccessToken, err := c.db.GetAServiceAccessTokenByUserID(user.ID, req.Action.Service)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = c.cGMAIL.ActionGmailRouter(
			ctx,
			action.GmailAct(req.Action.Route),
			req.Action.Params,
			respType,
			uint32(areaID),
			gmailAccessToken,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid service",
		})
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
