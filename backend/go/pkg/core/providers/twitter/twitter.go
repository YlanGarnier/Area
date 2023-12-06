package twitter

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
	clientId := os.Getenv("TWITTER_CLIENT_ID")
	if len(clientId) == 0 {
		return nil, fmt.Errorf("TWITTER_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("TWITTER_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("TWITTER_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("TWITTER_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("TWITTER_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("TWITTER_API_URL")
	if len(urlApi) == 0 {
		return nil, fmt.Errorf("TWITTER_API_URL is not set")
	}
	return &Provider{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApi:       urlApi,
	}, nil
}

func (p *Provider) GetUserTokenCredentials(code string, redirectUri string, _ models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", p.ClientID)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("code_verifier", "challenge")
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+"/token", "POST", data.Encode(), []utils.Header{
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
		{
			Key:   "Authorization",
			Value: "Basic " + base64.StdEncoding.EncodeToString([]byte(p.ClientID+":"+p.ClientSecret)),
		},
	}, "")
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token credentials: %w", err)
	}
	return tokenCredentials, nil
}

type MeResponse struct {
	Data struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"data"`
}

func (p *Provider) GetUserIdentifier(accessToken string) (string, error) {
	response, err := utils.GetFromJsonReq[MeResponse](p.UrlApi+"/users/me", "GET", "", []utils.Header{
		{
			Key:   "Authorization",
			Value: "Bearer " + accessToken,
		},
	}, "")
	if err != nil {
		return "", fmt.Errorf("failed to get user username: %w", err)
	}
	return response.Data.Username, nil
}

func (p *Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	panic("implement me")
}
