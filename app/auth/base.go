package auth

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/smbody/common-server/app"
)

type auth struct {
	auth_url  string
	token_url string
	tokens    map[string]string
}

var instance *auth = &auth{
	auth_url:  "http://localhost:8400/auth",
	token_url: "http://localhost:8400/auth/token"}

func Instance() *auth {
	return instance
}

func (auth *auth) GetUID(appId string, token string) (string, error) {
	uid, err := auth.validate_token(appId, token)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func appUrl(urlPath string, appId string) string {
	u, _ := url.Parse(urlPath)
	q := u.Query()
	q.Add("appId", appId)
	u.RawQuery = q.Encode()

	return u.String()
}

func (auth *auth) validate_token(appId string, token string) (string, error) {
	// спросим у периметра
	req, err := http.NewRequest(http.MethodGet, appUrl(auth.auth_url, appId), nil)
	if err != nil {
		return "", err
	}

	req.URL.Query().Add("appId", appId)
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	user := &User{}
	app.Read(resp.Body, user)

	return user.Id, nil
}
