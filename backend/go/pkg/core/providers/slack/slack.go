package slack

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
	clientId := os.Getenv("SLACK_CLIENT_ID")
	if len(clientId) == 0 {
		return nil, fmt.Errorf("SLACK_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("SLACK_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("SLACK_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("SLACK_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("SLACK_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("SLACK_API_URL")
	if len(urlApi) == 0 {
		return nil, fmt.Errorf("SLACK_API_URL is not set")
	}
	return &Provider{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApi:       urlApi,
	}, nil
}

type IdentifierResponse struct {
	User struct {
		ID string `json:"id"`
	}
}

func (p Provider) GetUserIdentifier(accessToken string) (string, error) {
	identifier, err := utils.GetFromJsonReq[IdentifierResponse](p.UrlApi+"/users.info", "GET", "", []utils.Header{
		{
			Key:   "Accept",
			Value: "application/json",
		},
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", accessToken),
		},
	}, "")
	if err != nil {
		return "", fmt.Errorf("failed to get user identifier: %v", err)
	}
	return identifier.User.ID, nil
}

func (p Provider) GetUserTokenCredentials(code string, redirectUri string, platform models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	data.Set("client_id", p.ClientID)
	data.Set("client_secret", p.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOauth+".access", "POST", data.Encode(), []utils.Header{
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
		return providers.TokenCredentials{}, fmt.Errorf("failed to get user token credentials: %v", err)
	}
	return tokenCredentials, nil
}

func (p Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	//TODO implement me
	panic("implement me")
}
