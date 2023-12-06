package spotify

import (
	"encoding/base64"
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

func NewProvider() (*Provider, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if len(clientID) == 0 {
		return nil, fmt.Errorf("SPOTIFY_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("SPOTIFY_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("SPOTIFY_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("SPOTIFY_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("SPOTIFY_API_URL")
	if len(urlApi) == 0 {
		return nil, fmt.Errorf("SPOTIFY_API_URL is not set")
	}
	return &Provider{
		ClientID:     clientID,
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
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+"/api/token",
		"POST", data.Encode(), []utils.Header{
			{
				Key:   "Content-Type",
				Value: "application/x-www-form-urlencoded",
			},
			{
				Key:   "Accept",
				Value: "application/json",
			},
			{
				Key:   "Authorization",
				Value: "Basic " + base64.StdEncoding.EncodeToString([]byte(p.ClientID+":"+p.ClientSecret)),
			},
		}, "")
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token: %v", err)
	}
	return tokenCredentials, nil
}

func (p *Provider) GetUserIdentifier(accessToken string) (string, error) {
	email, err := utils.GetFromJsonReq[string](p.UrlApi+"/me", "GET", "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", accessToken),
		},
	}, "email")
	if err != nil {
		return "", fmt.Errorf("failed to get email: %w", err)
	}
	return email, nil
}

func (p *Provider) RefreshToken(refreshToken string, _ models.Platform) (providers.TokenCredentials, error) {
	if len(refreshToken) == 0 {
		return providers.TokenCredentials{}, fmt.Errorf("no refresh token provided")
	}
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](fmt.Sprintf("%s/api/token", p.UrlOauth), "POST",
		data.Encode(), []utils.Header{
			{
				Key:   "Content-Type",
				Value: "application/x-www-form-urlencoded",
			},
			{
				Key:   "Accept",
				Value: "application/json",
			},
			{
				Key:   "Authorization",
				Value: "Basic " + base64.StdEncoding.EncodeToString([]byte(p.ClientID+":"+p.ClientSecret)),
			},
		}, "")
	if err != nil {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token credentials by refresh_token")
	}
	return tokenCredentials, nil
}
