package oauth

import (
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
	restClient           = resty.New()
	userLoginAPIEndpoint = "http://localhost:8080/users/login"
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
}
