package providers

import "github.com/lenismtho/area/pkg/core/models"

type Scopes uint

const (
	Email Scopes = iota
	Basic
)

type TokenCredentials struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Provider interface {
	GetUserIdentifier(accessToken string) (string, error)
	GetUserTokenCredentials(code string, redirectUri string, platform models.Platform) (TokenCredentials, error)
	RefreshToken(refreshToken string, platform models.Platform) (TokenCredentials, error)
}
