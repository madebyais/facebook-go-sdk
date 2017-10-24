// Package facebook implements methods for saving developer
// when they want to use Facebook Graph API using Golang
package facebook

import (
	"encoding/json"
	"errors"
	"io"
)

type (
	// Facebook is the base schema
	Facebook struct {
		AppID       string
		AppSecret   string
		Version     string
		AccessToken string
	}

	// Interface is the map interface for facebook-go-sdk
	Interface interface {
		SetAppID(appID string) Interface
		SetAppSecret(appSecret string) Interface
		SetVersion(version string) Interface
		SetAccessToken(accessToken string) Interface

		GetVersion() string
		GetAppID() string
		GetAppSecret() string

		GenerateAccessToken(redirectURI string, code string) (map[string]interface{}, error)

		Raw(method string, path string, queryStr string, body io.Reader) (interface{}, error)
		API(path string) APIInterface
	}
)

// New is use to initalize facebook-go-sdk
func New() Interface {
	return &Facebook{
		Version: `v2.10`,
	}
}

// SetAppID is use to set your application id
func (fb *Facebook) SetAppID(appID string) Interface {
	fb.AppID = appID
	return fb
}

// SetAppSecret is use to set your application secret
func (fb *Facebook) SetAppSecret(appSecret string) Interface {
	fb.AppSecret = appSecret
	return fb
}

// SetVersion is use to set the version of graph api
func (fb *Facebook) SetVersion(version string) Interface {
	fb.Version = version
	return fb
}

// SetAccessToken is use to set current access token
func (fb *Facebook) SetAccessToken(accessToken string) Interface {
	fb.AccessToken = `Bearer ` + accessToken
	return fb
}

// GetVersion returns the current used API version
func (fb *Facebook) GetVersion() string {
	return fb.Version
}

// GetAppID returns the current used app id
func (fb *Facebook) GetAppID() string {
	return fb.AppID
}

// GetAppSecret returns the current used app secret
func (fb *Facebook) GetAppSecret() string {
	return fb.AppSecret
}

// GenerateAccessToken returns access_token data
func (fb *Facebook) GenerateAccessToken(redirectURI string, code string) (map[string]interface{}, error) {
	if fb.GetAppID() == `` {
		return map[string]interface{}{}, errors.New(`Please set your application id.`)
	}

	if fb.GetAppSecret() == `` {
		return map[string]interface{}{}, errors.New(`Please set your application secret.`)
	}

	qs := `?client_id=` + fb.GetAppID() + `&client_secret=` + fb.GetAppSecret() + `&redirect_uri=` + redirectURI + `&code=` + code
	urlRaw := new(Graph).
		SetVersion(fb.Version).
		GenerateRawURL(`/oauth/access_token`, qs)

	data, err := new(Call).SetMethod(`GET`).SetURL(urlRaw).Submit()
	if err != nil {
		return map[string]interface{}{}, err
	}

	var dataAccessToken map[string]interface{}

	dataByte := []byte(data.(string))
	errJSONUnmarshal := json.Unmarshal(dataByte, &dataAccessToken)

	return dataAccessToken, errJSONUnmarshal
}

// Raw is use to call Facebook Graph API in raw method
// @method 		string				e.g. `GET`, `POST`
// @path 			strin					e.g. `/me`
// @queryStr 	string 				e.g. `?fields=id,name`
// @body			io.Reader			e.g. ... (optional)
func (fb *Facebook) Raw(method string, path string, queryStr string, body io.Reader) (interface{}, error) {
	urlRaw := new(Graph).SetVersion(fb.Version).GenerateRawURL(path, queryStr)
	return new(Call).
		SetMethod(method).
		SetURL(urlRaw).
		SetBody(body).
		AddHeader(`Authorization`, fb.AccessToken).
		Submit()
}

// API represents the Facebook Graph API that can be use by developers.
// It helps to simplify things
func (fb *Facebook) API(path string) APIInterface {
	return &API{
		FB:             fb,
		Path:           path,
		SelectedFields: ``,
		Message:        ``,
		ShareURL:       ``,
	}
}
