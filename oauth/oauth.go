package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func GetClientId(request *http.Request) int64 {
	if request == nil {
		return 0
	}

	callerId, err := strconv.ParseInt(request.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0
	}

	return callerId
}

func GetUserId(request *http.Request) int64 {
	if request == nil {
		return 0
	}
	return 0
}

func AuthenticateRequest(request *http.Request) *errors.RESTErr {
	if request == nil {
		return nil
	}

	cleanRequest(request)
	accessTokenId := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessTokenId == "" {
		return nil
	}

	at, err := getAccessToken(accessTokenId)
	if err != nil {
		return err
	}

	request.Header.Add(headerXClientId, fmt.Sprintf("%v", at.ClientId))
	request.Header.Add(headerXUserId, fmt.Sprintf("%v", at.UserId))

	return nil
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}

	request.Header.Del(headerXClientId)
	request.Header.Del(headerXUserId)
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
