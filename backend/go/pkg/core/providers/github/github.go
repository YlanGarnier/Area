package github

import (
	"fmt"
	"net/url"
	"os"

	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/providers"
	"github.com/lenismtho/area/pkg/core/utils"
	log "github.com/sirupsen/logrus"
)

type Provider struct {
	ClientIDWeb        string
	ClientSecretWeb    string
	ClientIDMobile     string
	ClientSecretMobile string
	UrlOAuth           string
	UrlApi             string
}

func NewProvider() (*Provider, error) {
	clientIDWeb := os.Getenv("GITHUB_CLIENT_ID_WEB")
	if clientIDWeb == "" {
		return nil, fmt.Errorf("GITHUB_CLIENT_ID is not set")
	}
	clientSecretWeb := os.Getenv("GITHUB_CLIENT_SECRET_WEB")
	if clientSecretWeb == "" {
		return nil, fmt.Errorf("GITHUB_CLIENT_SECRET is not set")
	}
	urlOAuth := os.Getenv("GITHUB_OAUTH_URL")
	if urlOAuth == "" {
		return nil, fmt.Errorf("GITHUB_OAUTH_URL is not set")
	}
	urlApi := os.Getenv("GITHUB_API_URL")
	if urlApi == "" {
		return nil, fmt.Errorf("GITHUB_API_URL is not set")
	}
	clientIDMobile := os.Getenv("GITHUB_CLIENT_ID_MOBILE")
	if len(clientIDMobile) == 0 {
		return nil, fmt.Errorf("GITHUB_CLIENT_ID_WEB is not set")
	}
	clientSecretMobile := os.Getenv("GITHUB_CLIENT_SECRET_MOBILE")
	if len(clientSecretMobile) == 0 {
		return nil, fmt.Errorf("GITHUB_CLIENT_SECRET_WEB")
	}
	return &Provider{
		ClientIDWeb:        clientIDWeb,
		ClientSecretWeb:    clientSecretWeb,
		ClientIDMobile:     clientIDMobile,
		ClientSecretMobile: clientSecretMobile,
		UrlOAuth:           urlOAuth,
		UrlApi:             urlApi,
	}, nil
}

type UserEmailResponse []struct {
	Email      string `json:"email"`
	Verified   bool   `json:"verified"`
	Primary    bool   `json:"primary"`
	Visibility string `json:"visibility"`
}

func (p *Provider) GetUserIdentifier(accessToken string) (string, error) {
	emails, err := utils.GetFromJsonReq[UserEmailResponse](p.UrlApi+"/user/emails", "GET", "", []utils.Header{
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
		{
			Key:   "Accept",
			Value: "application/vnd.github+json",
		},
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", accessToken),
		},
	}, "")
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}
	for _, email := range emails {
		if email.Primary && email.Verified {
			return email.Email, nil
		}
	}
	return "", fmt.Errorf("no primary email found")
}

func (p *Provider) GetUserTokenCredentials(code string, redirectUri string, platform models.Platform) (providers.TokenCredentials, error) {
	data := url.Values{}
	log.WithField("platform", platform).Info("here's the platform")
	if platform == models.Mobile {
		data.Set("client_id", p.ClientIDMobile)
		data.Set("client_secret", p.ClientSecretMobile)
	} else if platform == models.Web {
		data.Set("client_id", p.ClientIDWeb)
		data.Set("client_secret", p.ClientSecretWeb)
	}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOAuth+"/access_token",
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
		return providers.TokenCredentials{}, fmt.Errorf("failed to get github access_token: %v", err)
	}
	return tokenCredentials, nil
}

func (p *Provider) RefreshToken(refreshToken string, platform models.Platform) (providers.TokenCredentials, error) {
	if len(refreshToken) == 0 {
		return providers.TokenCredentials{}, fmt.Errorf("no refreshToken provided")
	}
	data := url.Values{}
	if platform == models.Mobile {
		data.Set("client_id", p.ClientIDMobile)
		data.Set("client_secret", p.ClientSecretMobile)
	} else if platform == models.Web {
		data.Set("client_id", p.ClientIDWeb)
		data.Set("client_secret", p.ClientSecretWeb)
	}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	tokenCredentials, err := utils.GetFromJsonReq[providers.TokenCredentials](p.UrlOAuth+"/access_token",
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
		return providers.TokenCredentials{}, fmt.Errorf("failed to get github access_token: %v", err)
	}
	return tokenCredentials, nil
}
