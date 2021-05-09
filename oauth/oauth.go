package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aditya43/bookstore-oauth-go/oauth/errors"
	"github.com/go-resty/resty/v2"
)

const (
	headerXPublic   = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXUserId   = "X-User-Id"

	paramAccessToken = "access_token"
)

var (
	restClient             = resty.New()
	getAccessTokenEndpoint = "http://localhost:8080/oauth/access_token/%s"
)

type accessToken struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	ClientId string `json:"client_id"`
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func AuthenticateRequest(request *http.Request) *errors.RESTErr {
	if request == nil {
		return nil
	}

	accessToken := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessToken == "" {
		return nil
	}
	return nil
}

func getAccessToken(accessTokenId string) (*accessToken, *errors.RESTErr) {
	var at accessToken
	response, err := restClient.R().
		SetResult(&at).
		Get(fmt.Sprintf(getAccessTokenEndpoint, accessTokenId))

	if response == nil || err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal") {
			return nil, errors.InternalServerErr("Invalid JSON response received")
		}
		return nil, errors.InternalServerErr(err.Error())
	}

	if response.StatusCode() == 200 && at.ClientId != "" {
		return &at, nil
	}

	var restErr errors.RESTErr
	if err := json.Unmarshal(response.Body(), &restErr); err != nil {
		return nil, errors.InternalServerErr("Invalid error interface")
	}

	return nil, &restErr
}
