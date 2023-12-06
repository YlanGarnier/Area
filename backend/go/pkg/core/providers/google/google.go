package google

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
	UrlApi       string
}

func NewGoogleProvider() (*Provider, error) {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	if len(clientId) == 0 {
		return nil, fmt.Errorf("failed to get GOOGLE_CLIENT_ID")
	}
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("failed to get GOOGLE_CLIENT_SECRET")
	}
	urlOauth := os.Getenv("GOOGLE_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("failed to get GOOGLE_OAUTH_URL")
	}
	urlApi := os.Getenv("GOOGLE_API_URL")
	if len(urlApi) == 0 {
		return nil, fmt.Errorf("failed to get GOOGLE_API_URL")
	}
	return &Provider{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApi:       urlApi,
	}, nil
}

func (p *Provider) GetUserTokenCredentials(code string, redirectUri string, _ models.Platform) (providers.TokenCredentials, error) {
	if len(code) == 0 {
		return providers.TokenCredentials{}, fmt.Errorf("no code provided")
	} else if len(redirectUri) == 0 {
		return providers.TokenCredentials{}, fmt.Errorf("no redirectUri provided")
	}
	data := url.Values{}
	data.Set("client_id", p.ClientID)
	data.Set("client_secret", p.ClientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+"/token",
		"POST", data.Encode(), []utils.Header{
			{
				Key:   "Content-Type",
				Value: "application/x-www-form-urlencoded",
			},
			{
				Key:   "Accept",
				Value: "application/json",
			},
		}, "")
	fmt.Println(tokenCredentials)
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token: %v", err)
	}
	return tokenCredentials, nil
}

func (p *Provider) GetUserIdentifier(accessToken string) (string, error) {
	identifier, err := utils.GetFromJsonReq[string](p.UrlApi+"/oauth2/v2/userinfo?access_token="+accessToken,
		"GET", "", []utils.Header{}, "email")
	if err != nil {
		return "", fmt.Errorf("failed to get user identifier: %v", err)
	}
	return identifier, nil
}

func (p *Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	data.Set("client_id", p.ClientID)
	data.Set("client_secret", p.ClientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+"/token", "POST", data.Encode(), []utils.Header{
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
		{
			Key:   "Accept",
			Value: "application/json",
		},
	}, "access_token")
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to refresh token: %v", err)
	}
	return tokenCredentials, nil
}
