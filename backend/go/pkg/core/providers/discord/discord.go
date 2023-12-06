package discord

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

func NewProvider() (*Provider, error) {
	clientId := os.Getenv("DISCORD_CLIENT_ID")
	if clientId == "" {
		return nil, fmt.Errorf("DISCORD_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, fmt.Errorf("DISCORD_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("DISCORD_OAUTH_URL")
	if urlOauth == "" {
		return nil, fmt.Errorf("DISCORD_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("DISCORD_API_URL")
	if urlApi == "" {
		return nil, fmt.Errorf("DISCORD_API_URL is not set")
	}
	return &Provider{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApi:       urlApi,
	}, nil
}

func (p *Provider) GetUserIdentifier(accessToken string) (string, error) {
	email, err := utils.GetFromJsonReq[string](p.UrlApi+"/users/@me", "GET", "", []utils.Header{
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
			Value: fmt.Sprintf("Bearer %s", accessToken),
		},
	}, "email")
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}
	return email, err
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
	//response, err := utils.HttpReq(p.UrlApi+"/oauth2/token", "POST", data.Encode(), []utils.Header{
	//	{
	//		Key:   "Content-Type",
	//		Value: "application/x-www-form-urlencoded",
	//	},
	//	{
	//		Key:   "Accept",
	//		Value: "application/json",
	//	},
	//})
	//log.WithField("redirect_uri = ", redirectUri)
	//fmt.Println("redirect uri = ", redirectUri)
	//log.WithField("response", string(response)).Info("response")
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlApi+"/oauth2/token",
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
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token: %v", err)
	}
	return tokenCredentials, nil

}

func (p *Provider) RefreshToken(refreshToken string, _ models.Platform) (providers.TokenCredentials, error) {
	if len(refreshToken) == 0 {
		return providers.TokenCredentials{}, fmt.Errorf("no refreshToken provided")
	}
	data := url.Values{}
	data.Set("client_id", p.ClientID)
	data.Set("client_secret", p.ClientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlApi+"/oauth2/token",
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
	if err != nil {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token: %v", err)
	}
	return tokenCredentials, nil
}
