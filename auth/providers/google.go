package providers

import (
	"encoding/json"
	"fmt"
	"github.com/carrizoaagustin/cv-online/auth"
	"io"
	"net/http"
	"net/url"
)

type GoogleProvider struct {
	config oauth2Config
}

func NewGoogleProvider(p auth.OAuthConfigParams) auth.OAuthProvider {
	return &GoogleProvider{
		config: oauth2Config{
			ClientID:  p.ClientID,
			SecretKey: p.SecretKey,
			Name:      "google",
			Endpoints: Endpoints{
				AuthorizationURL: "https://accounts.google.com/o/oauth2/v2/auth",
				TokenURL:         "https://oauth2.googleapis.com/token",
				UserInfoURL:      "https://www.googleapis.com/oauth2/v2/userinfo",
			},
			Scopes: p.Scopes,
		},
	}
}

func (o GoogleProvider) GetName() string {
	return o.config.Name
}

func (o GoogleProvider) GetUserInformation(accessToken string) (*auth.UserInformation, error) {
	req, err := http.NewRequest(http.MethodGet, o.config.Endpoints.UserInfoURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decodificar el JSON de la respuesta
	var userData map[string]interface{}
	if err := json.Unmarshal(body, &userData); err != nil {
		return nil, err
	}

	mappedUserInfo := auth.UserInformation{
		ID:    userData["id"].(string),    //nolint:errcheck
		Email: userData["email"].(string), //nolint:errcheck
		Name:  userData["name"].(string),  //nolint:errcheck
	}

	return &mappedUserInfo, nil
}

func (o GoogleProvider) GetAuthorizationURL() string {
	return o.config.Endpoints.AuthorizationURL + "?scope=https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile&response_type=code&client_id=240652747137-s3krn73b2qagrojnftun4cirpipr5pc2.apps.googleusercontent.com&redirect_uri=http://localhost:8000/api/resources/callback"
}

func (o GoogleProvider) RedeemToken(authorizationCode string) (string, error) {
	data := url.Values{}
	data.Set("code", authorizationCode)
	data.Set("client_id", o.config.ClientID)
	data.Set("client_secret", o.config.SecretKey)
	data.Set("redirect_uri", "http://localhost:8000/api/resources/callback")
	data.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(o.config.Endpoints.TokenURL, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse map[string]interface{}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return accessToken, nil
}
