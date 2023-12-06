package services

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/providers"
	"github.com/lenismtho/area/pkg/core/providers/discord"
	"github.com/lenismtho/area/pkg/core/providers/github"
)

type dbService interface {
	SetUserService(service models.Service) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserServiceTokenCredentialsByUserID(userID uint) (providers.TokenCredentials, error)
	UpdateServiceTokenCredentialsByUserID(credentials providers.TokenCredentials, userID uint) error
}

type Service struct {
	db                              dbService
	clientCallBackAuth              string
	clientCallBackAuthServiceMobile string
	providers                       map[models.Provider]providers.Provider
}

func NewService(db dbService) (*Service, error) {
	discordP, err := discord.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create discord provider: %w", err)
	}
	githubP, err := github.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to create github provider: %w", err)
	}
	clientCallBackAuth := os.Getenv("CLIENT_CALLBACK_AUTH_SERVICE")
	if len(clientCallBackAuth) == 0 {
		return nil, fmt.Errorf("CLIENT_CALLBACK_AUTH_SERVICE is not set")
	}
	clientCallBackAuthServiceMobile := os.Getenv("CLIENT_CALLBACK_AUTH_SERVICE_MOBILE")
	return &Service{
		db:                              db,
		clientCallBackAuth:              clientCallBackAuth,
		clientCallBackAuthServiceMobile: clientCallBackAuthServiceMobile,
		providers: map[models.Provider]providers.Provider{
			models.Discord: discordP,
			models.Github:  githubP,
		},
	}, nil
}

func (s *Service) RefreshToken(ctx *gin.Context) {
	//providerID := ctx.Query("provider")
	//if len(providerID) == 0 {
	//	log.Error("no provider provided")
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no provider provided"})
	//	return
	//}
	//provider := s.providers[models.Provider(providerID)]
	//if provider == nil {
	//	log.Error("invalid provider")
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid provider"})
	//	return
	//}
	//userID := ctx.Query("user_id")
	//if len(userID) == 0 {
	//	log.Error("no user_id provided")
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no user_id provided"})
	//}
	//userIDUint, err := strconv.ParseUint(userID, 10, 32)
	//if err != nil {
	//	log.WithField("error", err).Error("failed to parse userID")
	//	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//user, err := s.db.GetUserByID(uint(userIDUint))
	//if err != nil {
	//	log.WithField("error", err).Error("failed to get user")
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//tokenCredentials, err := s.db.GetUserServiceTokenCredentialsByUserID(user.ID)
	//if err != nil {
	//	log.WithField("error", err).Error("failed to get token credentials")
	//	ctx.AbortWithStatusJSON(http.StatusTooEarly, gin.H{"error": err.Error()})
	//	return
	//}
	//newTokenCredentials, err := provider.RefreshToken(tokenCredentials.RefreshToken)
	//if err != nil {
	//	log.WithField("error", err).Error("failed to refresh token")
	//	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//err = s.db.UpdateServiceTokenCredentialsByUserID(newTokenCredentials, user.ID)
	//if err != nil {
	//	log.WithField("error", err).Error("failed to update token credentials")
	//	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
}
