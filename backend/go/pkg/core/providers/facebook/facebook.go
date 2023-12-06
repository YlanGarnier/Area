package facebook

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
	clientID := os.Getenv("FACEBOOK_CLIENT_ID")
	if len(clientID) == 0 {
		return nil, fmt.Errorf("FACEBOOK_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("FACEBOOK_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("FACEBOOK_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("FACEBOOK_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("FACEBOOK_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("FACEBOOK_API_URL")
	if len(urlApi) == 0 {
		return nil, fmt.Errorf("FACEBOOK_API_URL is not set")
	}
	return &Provider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApi:       urlApi,
	}, nil
}

func (p Provider) GetUserIdentifier(accessToken string) (string, error) {
	identifier, err := utils.GetFromJsonReq[string](p.UrlApi+"/me?access_token="+accessToken, "GET", "", []utils.Header{}, "id")
	if err != nil {
		return "", fmt.Errorf("failed to get user identifier: %v", err)
	}
	return identifier, nil
}

func (p Provider) GetUserTokenCredentials(code string, redirectUri string, platform models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("client_id", p.ClientID)
	data.Set("client_secret", p.ClientSecret)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+"/access_token", "POST", data.Encode(), []utils.Header{
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
	}, "")
	if err != nil || tokenCredentials == (providers.TokenCredentials{}) {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get token credentials: %v", err)
	}
	return tokenCredentials, nil
}

func (p Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	panic("implement me")
}
