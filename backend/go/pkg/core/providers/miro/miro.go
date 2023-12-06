package miro

import (
	"fmt"
	"net/url"
	"os"

	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/providers"
	"github.com/lenismtho/area/pkg/core/utils"
)

type Provider struct {
	ClientID     string
	ClientSecret string
	UrlOauth     string
	UrlApiV1     string
	urlApiV2     string
}

func NewProvider() (*Provider, error) {
	clientId := os.Getenv("MIRO_CLIENT_ID")
	if len(clientId) == 0 {
		return nil, fmt.Errorf("MIRO_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("MIRO_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("MIRO_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("MIRO_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("MIRO_OAUTH_URL is not set")
	}
	urlApiV1 := os.Getenv("MIRO_API_V1_URL")
	if len(urlApiV1) == 0 {
		return nil, fmt.Errorf("MIRO_API_V1_URL is not set")
	}
	urlApiV2 := os.Getenv("MIRO_API_V2_URL")
	if len(urlApiV2) == 0 {
		return nil, fmt.Errorf("MIRO_API_V2_URL is not set")
	}
	return &Provider{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApiV1:     urlApiV1,
		urlApiV2:     urlApiV2,
	}, nil
}

func (p Provider) GetUserIdentifier(accessToken string) (string, error) {
	identifier, err := utils.GetFromJsonReq[string](p.UrlApiV1+"/users/me", "GET", "", []utils.Header{
		{
			Key:   "Accept",
			Value: "application/json",
		},
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", accessToken),
		},
	}, "email")
	if err != nil {
		return "", fmt.Errorf("failed to get user identifier: %v", err)
	}
	return identifier, nil
}

func (p Provider) GetUserTokenCredentials(code string, redirectUri string, platform models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", p.ClientID)
	data.Set("client_secret", p.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+"/token", "POST", data.Encode(), []utils.Header{
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
		{
			Key:   "Accept",
			Value: "application/json",
		},
	}, "")
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token credentials: %w", err)
	}
	return tokenCredentials, nil
}

func (p Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	//TODO implement me
	panic("implement me")
}
