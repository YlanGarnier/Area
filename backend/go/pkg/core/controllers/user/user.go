package user

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lenismtho/area/pkg/core/authentication"
	"github.com/lenismtho/area/pkg/core/cache"
	"github.com/lenismtho/area/pkg/core/database"
	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/providers"
	"github.com/lenismtho/area/pkg/core/providers/discord"
	"github.com/lenismtho/area/pkg/core/providers/dropbox"
	"github.com/lenismtho/area/pkg/core/providers/facebook"
	"github.com/lenismtho/area/pkg/core/providers/github"
	"github.com/lenismtho/area/pkg/core/providers/google"
	"github.com/lenismtho/area/pkg/core/providers/linkedin"
	"github.com/lenismtho/area/pkg/core/providers/miro"
	"github.com/lenismtho/area/pkg/core/providers/notion"
	"github.com/lenismtho/area/pkg/core/providers/slack"
	"github.com/lenismtho/area/pkg/core/providers/spotify"
	"github.com/lenismtho/area/pkg/core/providers/twitch"
	"github.com/lenismtho/area/pkg/core/providers/twitter"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type dbUser interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserToken(email string, token string) error
	GetUserByToken(token string) (*models.User, error)
	GetServicesByUserID(userID uint) ([]models.Service, error)
	GetUserAreasByUserID(userID uint) ([]models.Area, error)
	SetUserService(service models.Service) error
	UpdateUserUsername(username string, userID uint) error
	UpdateUserEmail(email string, userID uint) error
	UpdateUserPassword(password string, userID uint) error
	UpdateUserFirstName(firstName string, userID uint) error
	UpdateUserLastName(lastName string, userID uint) error
	DeleteUser(userID uint) error
	DeleteServicesByUserID(userId uint) error
	DeleteServiceByUserIDAndName(userId uint, name string) error
}

type Controller struct {
	db                dbUser
	cacheDB           cache.Cache
	providers         map[models.Provider]providers.Provider
	redirectUriMobile string
}

func NewController(db dbUser, cacheDb cache.Cache) (*Controller, error) {
	githubP, err := github.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create github provider: %w", err)
	}
	discordP, err := discord.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create discord provider: %w", err)
	}
	spotifyP, err := spotify.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failted to create spotify provider: %w", err)
	}
	googleP, err := google.NewGoogleProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create google provider: %w", err)
	}
	twitterP, err := twitter.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create twitter provider: %w", err)
	}
	facebookP, err := facebook.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create facebook provider: %w", err)
	}
	miroP, err := miro.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create miro provider: %w", err)
	}
	twitchP, err := twitch.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create twitch provider: %w", err)
	}
	notionP, err := notion.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create notion provider: %w", err)
	}
	slackP, err := slack.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create slack provider: %w", err)
	}
	dropboxP, err := dropbox.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create dropbox provider: %w", err)
	}
	linkedinP, err := linkedin.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create linkedin provider: %w", err)
	}
	redirectUriMobile := os.Getenv("REDIRECT_URI_MOBILE_CALLBACK")
	if len(redirectUriMobile) == 0 {
		return nil, errors.New("failed to get REDIRECT_URI_MOBILE_CALLBACK env")
	}
	return &Controller{
		db:                db,
		cacheDB:           cacheDb,
		redirectUriMobile: redirectUriMobile,
		providers: map[models.Provider]providers.Provider{
			models.Discord:  discordP,
			models.Github:   githubP,
			models.Spotify:  spotifyP,
			models.Google:   googleP,
			models.Twitter:  twitterP,
			models.Facebook: facebookP,
			models.Miro:     miroP,
			models.Twitch:   twitchP,
			models.Notion:   notionP,
			models.Slack:    slackP,
			models.Dropbox:  dropboxP,
			models.Linkedin: linkedinP,
		},
	}, nil
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login 		godoc
// @Summary			Login as a user.
// @Tags			users
// @Param			LoginRequest	body	LoginRequest	true	"Login as a user"
// @Accept			json
// @Success			200 {object} LoginResponse
// @Failure			400 "Bad request: email or password is missing"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind LoginRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.db.GetUserByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.WithField("error", err).Error("user not found")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		log.WithField("error", err).Error("failed to get user by email")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ok, err := authentication.CompareHashPassword(user.Password.String, req.Password)
	if err != nil {
		log.WithField("error", err).Error("failed to compare password")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		log.WithField("error", err).Error("wrong password")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "wrong password"})
		return
	}
	token, err := authentication.GenerateHashPassword(user.Email)
	if err != nil {
		log.WithField("error", err).Error("failed to generate OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err = authentication.GenerateOAT(token)
	if err != nil {
		log.WithField("error", err).Error("failed to generate OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithField("token", token).Info("token")
	hashToken, err := authentication.HashOAT(token)
	if err != nil {
		log.WithField("error", err).Error("failed to hash OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithField("hashToken", hashToken).Info("hashToken")
	err = c.db.UpdateUserToken(user.Email, hashToken)
	if err != nil {
		log.WithField("error", err).Error("failed to update user token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUp 		godoc
// @Summary			Create a new user.
// @Tags			users
// @Param			SignUpRequest	body	SignUpRequest	true	"Created a user object"
// @Accept			json
// @Success			201 {token} string
// @Failure         409 "User already exists"
// @Failure			400 "Bad request: email or password is missing or invalid"
// @Failure 		500 "Internal server error"
// @Router			/users/signup [post]
func (c *Controller) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind SignUpRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.Email) == 0 || len(req.Password) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	hashPassword, err := authentication.GenerateHashPassword(req.Password)
	if err != nil {
		log.WithField("error", err).Error("failed to hash password")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser := models.User{
		Email:     req.Email,
		FirstName: sql.NullString{String: "", Valid: false},
		LastName:  sql.NullString{String: "", Valid: false},
		Password:  sql.NullString{String: hashPassword, Valid: true},
		Username:  sql.NullString{String: "", Valid: false}, // Username is not required
		Kind:      "password",
		Token:     sql.NullString{String: "", Valid: false},
	}
	err = c.db.CreateUser(newUser)
	if err != nil {
		if database.DbIsError(err, database.ErrUniqueConstraintFailed) {
			log.WithField("error", err).Error("user already exists")
			ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		} else {
			log.WithField("error", err).Error("failed to create user")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	token, err := authentication.GenerateHashPassword(req.Email)
	if err != nil {
		log.WithField("error", err).Error("failed to generate OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err = authentication.GenerateOAT(token)
	if err != nil {
		log.WithField("error", err).Error("failed to generate OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithField("token", token).Info("token")
	hashToken, err := authentication.HashOAT(token)
	if err != nil {
		log.WithField("error", err).Error("failed to hash OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithField("hashToken", hashToken).Info("hashToken")
	err = c.db.UpdateUserToken(req.Email, hashToken)
	if err != nil {
		log.WithField("error", err).Error("failed to update user token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

type ResetUserPasswordRequest struct {
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

// ResetUserPassword godoc
// @Summary			Reset a user password.
// @Tags			user
// @Param			ResetUserPasswordRequest	body	ResetUserPasswordRequest	true	"Reset a user password"
// @Accept			json
// @Success			204
// @Failure 	    400 "Bad request: password or token is missing"
// @Failure         404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/reset_password [put]
func (c *Controller) ResetUserPassword(ctx *gin.Context) {

}

type MeResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Kind      string `json:"kind"`
}

// Me godoc
// @Summary			Get my information as a user.
// @Tags			users
// @Accept			json
// @Success			200 {object} MeResponse
// @Failure			401 "Unauthorized"
// @Failure         404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/me [get]
func (c *Controller) Me(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		log.WithField("error", "no authorization provided").Error("failed to get authorization")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization provided"})
	}
	token := authorization[7:]
	if len(token) == 0 {
		log.WithField("error", "no token provided").Error("failed to get token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		return
	}
	user, err := c.db.GetUserByToken(token)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.WithField("error", err).Error("user not found")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		log.WithField("error", err).Error("failed to get user by authorization")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Kind:      user.Kind,
	})
}

type UpdateMeRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UpdateMe godoc
// @Summary			Update my information as a user.
// @Tags			users
// @Accept			json
// @Param			UpdateMeRequest	body	UpdateMeRequest	true	"Update my information as a user"
// @Success			204
// @Failure 	    400 "Bad request: username, first_name and last_name is missing"
// @Failure 	    401 "Unauthorized"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/me [put]
func (c *Controller) UpdateMe(ctx *gin.Context) {
	var req UpdateMeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind UpdateMeRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.LastName) == 0 && len(req.FirstName) == 0 && len(req.Username) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		log.WithField("error", "no authorization provided").Error("failed to get authorization")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization provided"})
	}
	token := authorization[7:]
	if len(token) == 0 {
		log.WithField("error", "no token provided").Error("failed to get token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		return
	}
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		log.WithField("error", err).Error("failed to get user by authorization")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(req.Username) != 0 {
		err = c.db.UpdateUserUsername(req.Username, user.ID)
		if err != nil {
			log.WithField("error", err).Error("failed to update user username")
		}
	}
	if len(req.FirstName) != 0 {
		err = c.db.UpdateUserFirstName(req.FirstName, user.ID)
		if err != nil {
			log.WithField("error", err).Error("failed to update user first name")
		}
	}
	if len(req.LastName) != 0 {
		err = c.db.UpdateUserLastName(req.LastName, user.ID)
		if err != nil {
			log.WithField("error", err).Error("failed to update user last name")
		}
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

type UpdateMeEmailRequest struct {
	NewEmail string `json:"email"`
	Password string `json:"password"` // Password to confirm the change
}

// UpdateMeEmail godoc
// @Summary			Update my email as a user.
// @Tags			users
// @Accept			json
// @Param			UpdateMeEmailRequest	body	UpdateMeEmailRequest	true	"Update my email as a user"
// @Success			204
// @Failure 	    400 "Bad request: email is missing"
// @Failure 	    401 "Unauthorized"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/me/email [put]
func (c *Controller) UpdateMeEmail(ctx *gin.Context) {
	var req UpdateMeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind UpdateMeEmailRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		log.WithField("error", "no authorization provided").Error("failed to get authorization")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization provided"})
	}
	token := authorization[7:]
	if len(token) == 0 {
		log.WithField("error", "no token provided").Error("failed to get token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		return
	}
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		log.WithField("error", err).Error("failed to get user by authorization")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ok, err := authentication.CompareHashPassword(user.Password.String, req.Password)
	if err != nil {
		log.WithField("error", err).Error("failed to compare password")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		log.WithField("error", err).Error("wrong password")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}
	err = c.db.UpdateUserEmail(req.NewEmail, user.ID)
	if err != nil {
		log.WithField("error", err).Error("failed to update user email")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

type UpdateMePasswordRequest struct {
	PreviousPassword string `json:"previous_password"` // PreviousPassword to confirm the change
	NewPassword      string `json:"new_password"`
}

// UpdateMePassword  godoc
// @Summary			Update my password as a user .
// @Tags			users
// @Param			UpdateMePasswordRequest	body	UpdateMePasswordRequest	true	"Update my password as a user"
// @Accept			json
// @Success			200
// @Failure 	    400 "Bad request: password is missing"
// @Failure 	    401 "Unauthorized"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/me/password [put]
func (c *Controller) UpdateMePassword(ctx *gin.Context) {
	var req UpdateMePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind UpdateMeEmailRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		log.WithField("error", "no authorization provided").Error("failed to get authorization")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization provided"})
	}
	token := authorization[7:]
	if len(token) == 0 {
		log.WithField("error", "no token provided").Error("failed to get token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		return
	}
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		log.WithField("error", err).Error("failed to get user by authorization")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ok, err := authentication.CompareHashPassword(user.Password.String, req.PreviousPassword)
	if err != nil {
		log.WithField("error", err).Error("failed to compare password")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		log.WithField("error", err).Error("wrong password")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}
	hashPassword, err := authentication.GenerateHashPassword(req.NewPassword)
	if err != nil {
		log.WithField("error", err).Error("failed to hash password")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.db.UpdateUserPassword(hashPassword, user.ID)
	if err != nil {
		log.WithField("error", err).Error("failed to update user email")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// DeleteMe godoc
// @Summary			Delete my account as a user.
// @Tags			users
// @Accept			json
// @Success			204
// @Failure 	    401 "Unauthorized"
// @Failure 	    404 "User not found"
// @Failure 		500 "Internal server error"
// @Router			/users/me [delete]
func (c *Controller) DeleteMe(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		log.WithField("error", "no authorization provided").Error("failed to get authorization")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization provided"})
	}
	token := authorization[7:]
	if len(token) == 0 {
		log.WithField("error", "no token provided").Error("failed to get token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		return
	}
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		log.WithField("error", err).Error("failed to get user by authorization")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.db.DeleteUser(user.ID)
	if err != nil {
		log.WithField("error", err).Error("failed to delete user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.db.DeleteServicesByUserID(user.ID)
	if err != nil {
		log.WithField("error", err).Error("failed to delete user services")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

type GetUserServicesResponse struct {
	Name string `json:"name"`
}

// GetUserServices godoc
// @Summary			Get my services as a user.
// @Tags			users
// @Accept			json
// @Success			200 {object} []GetUserServicesResponse
// @Failure 	    401 "Unauthorized"
// @Failure 		500 "Internal server error"
// @Router			/users/me/services [get]
func (c *Controller) GetUserServices(ctx *gin.Context) {
	headerTab := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(headerTab) < 2 || headerTab[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
	}
	token := headerTab[1]
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	services, err := c.db.GetServicesByUserID(user.ID)
	response := make([]GetUserServicesResponse, 0)
	for _, service := range services {
		response = append(response, GetUserServicesResponse{Name: service.Name})
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

type GetUserAreasResponse struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	ActionService        string `json:"action_service"`
	RouteActionService   string `json:"route_action_service"`
	ReactionService      string `json:"reaction_service"`
	RouteReactionService string `json:"route_reaction_service"`
}

// GetUserAreas godoc
// @Summary			Get my areas as a user.
// @Tags			users
// @Accept			json
// @Success			200 {object} []GetUserAreasResponse
// @Failure 	    401 "Unauthorized"
// @Failure 		500 "Internal server error"
// @Router			/users/me/areas [get]
func (c *Controller) GetUserAreas(ctx *gin.Context) {
	headerTab := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(headerTab) == 0 || headerTab[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
	}
	token := headerTab[1]
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	areas, err := c.db.GetUserAreasByUserID(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(areas)
	response := make([]GetUserAreasResponse, 0)
	for _, area := range areas {
		response = append(response, GetUserAreasResponse{
			ID:                   area.ID,
			Name:                 area.Name,
			ActionService:        area.ActService,
			RouteActionService:   area.RouteAction,
			ReactionService:      area.ReaServices,
			RouteReactionService: area.Route,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

type OauthRequest struct {
	Code        string `json:"code"`
	Provider    string `json:"provider"`
	RedirectURI string `json:"redirect_uri"`
	Platform    string `json:"platform"`
}
type OauthResponse struct {
	Token string `json:"token"`
}

// SignUpOauth godoc
// @Summary			Create a new user with oauth.
// @Tags			users
// @Param 		 OauthRequest	body	OauthRequest	true	"Create a new user with oauth"
// @Accept			json
// @Success			201 {OauthResponse} string "User created"
// @Success 	    302 {OauthResponse} string "User already exists, login"
// @Failure 	    400 "Bad request: email or code is missing"
// @Failure 	    500 "Internal server error"
// @Router			/users/signup/oauth [post]
func (c *Controller) SignUpOauth(ctx *gin.Context) {
	var req OauthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind OauthRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.Code) == 0 || len(req.Provider) == 0 || len(req.RedirectURI) == 0 || len(req.Platform) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	if models.Provider(req.Provider) != models.Github && models.Provider(req.Provider) != models.Discord {
		log.Error("provider provided not available")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "provider not available"})
		return
	}
	provider := c.providers[models.Provider(req.Provider)]
	if provider == nil {
		log.Error("provider not found")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong provider"})
		return
	}
	tokenCredentials, err := provider.GetUserTokenCredentials(req.Code, req.RedirectURI, models.Platform(req.Platform))
	if err != nil {
		log.WithField("error", err).Error("failed to get access token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithField("token Credentials", tokenCredentials).Info("test")
	email, err := provider.GetUserIdentifier(tokenCredentials.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("failed to get email")
		ctx.JSON(http.StatusInternalServerError, "failed to get email")
		return
	}
	token, err := authentication.GenerateHashPassword(email)
	if err != nil {
		log.WithField("error", err).Error("failed to generate OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err = authentication.GenerateOAT(token)
	if err != nil {
		log.WithField("error", err).Error("failed to generate OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hashToken, err := authentication.HashOAT(token)
	if err != nil {
		log.WithField("error", err).Error("failed to hash OAT")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser := models.User{
		Email:    email,
		Password: sql.NullString{String: "", Valid: false},
		Username: sql.NullString{String: "", Valid: false}, // Username is not required
		Kind:     req.Provider,
		Token:    sql.NullString{String: hashToken, Valid: true},
	}
	err = c.db.CreateUser(newUser)
	if err != nil {
		if database.DbIsError(err, database.ErrUniqueConstraintFailed) {
			log.WithField("error", err).Error("user already exists")
			err = c.db.UpdateUserToken(email, hashToken)
			if err != nil {
				log.WithField("error", err).Error("failed to update user token")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusFound, gin.H{"token": token})
			return
		} else {
			log.WithField("error", err).Error("failed to create user")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

// UserConnectService godoc
// @Summary			 Connect user to a service.
// @Tags			users
// @Param 		 OauthRequest	body	OauthRequest	true	"Connect a user to a service"
// @Accept			json
// @Success			201
// @Failure 	    400 "Bad request: parameters are missing"
// @Failure 	    500 "Internal server error"
// @Router			/users/me/services [post]
func (c *Controller) UserConnectService(ctx *gin.Context) {
	var req OauthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.WithField("error", err).Error("failed to bind OauthRequest")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.Code) == 0 || len(req.Provider) == 0 || len(req.RedirectURI) == 0 || len(req.Platform) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	headerTab := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(headerTab) == 0 || headerTab[0] != "Bearer" {
		log.Error("no authorization token")
		ctx.JSON(http.StatusForbidden, gin.H{"error": "no authorization token provided"})
		return
	}
	token := headerTab[1]
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		log.WithField("error", err.Error()).Info("failed to retrieve user by token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	providerReq := req.Provider
	if strings.Contains(providerReq, "google") {
		providerReq = "google"
	}
	provider := c.providers[models.Provider(providerReq)]
	if provider == nil {
		log.Error("provider not found")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong provider"})
		return
	}
	tokenCredentials, err := provider.GetUserTokenCredentials(req.Code, req.RedirectURI, models.Platform(req.Platform))
	if err != nil {
		log.WithField("error", err).Info("failed to get token credentials")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(tokenCredentials)
	identifier, err := provider.GetUserIdentifier(tokenCredentials.AccessToken)
	if err != nil {
		log.WithField("error", err.Error()).Info("failed to get identifier")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.db.SetUserService(models.Service{
		Name:         req.Provider,
		Identifier:   identifier,
		AccessToken:  tokenCredentials.AccessToken,
		RefreshToken: tokenCredentials.RefreshToken,
		UserID:       user.ID,
	})
	if err != nil {
		log.WithField("error", err.Error()).Info("failed to set the service to the user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

// UserDisconnectService godoc
// @Summary			 Disconnect user to a service.
// @Tags			users
// @Param 		    serviceName	path	string	true	"Service Name"
// @Success			204
// @Failure 	    400 "Bad request: parameters are missing"
// @Failure 	    404 "Service not found"
// @Failure 	    500 "Internal server error"
// @Router			/users/me/services/{serviceName} [delete]
func (c *Controller) UserDisconnectService(ctx *gin.Context) {
	serviceName := ctx.Param("serviceName")
	if len(serviceName) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	headerTab := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(headerTab) == 0 || headerTab[0] != "Bearer" {
		log.Error("no authorization token")
		ctx.JSON(http.StatusForbidden, gin.H{"error": "no authorization token provided"})
		return
	}
	token := headerTab[1]
	user, err := c.db.GetUserByToken(token)
	if err != nil {
		log.WithField("error", err.Error()).Info("failed to retrieve user by token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.db.DeleteServiceByUserIDAndName(user.ID, serviceName)
	if err != nil {
		log.WithField("error", err.Error()).Info("failed to delete user service")
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *Controller) MobileCallback(ctx *gin.Context) {
	state := strings.Split(ctx.Query("state"), " ")
	if len(state) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	method := ""
	if len(state) == 1 {
		method = "login"
	} else if len(state) == 2 {
		method = "connect_service"
	} else {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	code := ctx.Query("code")
	if len(code) == 0 {
		log.Error("failed to get parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}
	providerReq := state[0]
	if strings.Contains(providerReq, "google") {
		providerReq = "google"
	}
	provider := c.providers[models.Provider(providerReq)]
	if provider == nil {
		log.Error("provider not found")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong provider"})
		return
	}
	tokenCredentials, err := provider.GetUserTokenCredentials(code, "https://stirred-kit-alive.ngrok-free.app/v1/mobile/callback", models.Mobile)
	if err != nil {
		log.WithField("error", err).Info("failed to get token credentials")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	identifier, err := provider.GetUserIdentifier(tokenCredentials.AccessToken)
	if err != nil {
		log.WithField("error", err.Error()).Info("failed to get identifier")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if method == "login" {
		token, err := authentication.GenerateHashPassword(identifier)
		if err != nil {
			log.WithField("error", err).Error("failed to generate OAT")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		token, err = authentication.GenerateOAT(token)
		if err != nil {
			log.WithField("error", err).Error("failed to generate OAT")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		hashToken, err := authentication.HashOAT(token)
		if err != nil {
			log.WithField("error", err).Error("failed to hash OAT")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		newUser := models.User{
			Email:    identifier,
			Password: sql.NullString{String: "", Valid: false},
			Username: sql.NullString{String: "", Valid: false}, // Username is not required
			Kind:     state[0],
			Token:    sql.NullString{String: hashToken, Valid: true},
		}
		err = c.db.CreateUser(newUser)
		if err != nil {
			if database.DbIsError(err, database.ErrUniqueConstraintFailed) {
				log.WithField("error", err).Error("user already exists")
				err = c.db.UpdateUserToken(identifier, hashToken)
				if err != nil {
					log.WithField("error", err).Error("failed to update user token")
					ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?error=failed to update user token")
					return
				}
				ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?token="+token+"&state="+state[0]+"&status=302")
				return
			} else {
				log.WithField("error", err).Error("failed to create user")
				ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?error=failed to create user")
				return
			}
		}
		ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?token="+token+"&state="+state[0])
	} else if method == "connect_service" {
		user, err := c.db.GetUserByToken(state[1])
		if err != nil {
			log.WithField("error", err.Error()).Info("Token must be valid")
			ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?error=Token must be valid")
			return
		}
		err = c.db.SetUserService(models.Service{
			Name:         state[0],
			Identifier:   identifier,
			AccessToken:  tokenCredentials.AccessToken,
			RefreshToken: tokenCredentials.RefreshToken,
			UserID:       user.ID,
		})
		if err != nil {
			log.WithField("error", err.Error()).Info("failed to set the service to the user")
			ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?error=failed to set the service to the user")
			return
		}
		ctx.Redirect(http.StatusPermanentRedirect, c.redirectUriMobile+"?token="+state[1])
	}
}
