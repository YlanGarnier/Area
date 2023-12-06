package notion

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
	clientId := os.Getenv("NOTION_CLIENT_ID")
	if len(clientId) == 0 {
		return nil, fmt.Errorf("NOTION_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("NOTION_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return nil, fmt.Errorf("NOTION_CLIENT_SECRET is not set")
	}
	urlOauth := os.Getenv("NOTION_OAUTH_URL")
	if len(urlOauth) == 0 {
		return nil, fmt.Errorf("NOTION_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("NOTION_API_URL")
	if len(urlApi) == 0 {
		return nil, fmt.Errorf("NOTION_API_URL is not set")
	}
	return &Provider{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		UrlOauth:     urlOauth,
		UrlApi:       urlApi,
	}, nil
}

type IdentifierResponse struct {
	Results []struct {
		Person struct {
			Email string `json:"email"`
		}
	}
}

func (p Provider) GetUserIdentifier(accessToken string) (string, error) {
	identifier, err := utils.GetFromJsonReq[IdentifierResponse](p.UrlApi+"/users", "GET", "", []utils.Header{
		{
			Key:   "Accept",
			Value: "application/json",
		},
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", accessToken),
		},
		{
			Key:   "Notion-Version",
			Value: "2022-06-28",
		},
	}, "")
	if err != nil || len(identifier.Results[0].Person.Email) == 0 {
		return "", fmt.Errorf("failed to get user identifier: %v", err)
	}
	return identifier.Results[0].Person.Email, nil
}

func (p Provider) GetUserTokenCredentials(code string, redirectUri string, platform models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
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
		return providers.TokenCredentials{}, fmt.Errorf("failed to get user token credentials: %v", err)
	}
	return tokenCredentials, nil
}

func (p Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	//TODO implement me
	panic("implement me")
}
