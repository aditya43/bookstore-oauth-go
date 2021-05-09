package oauth

import "net/http"

const (
	headerXPublic = "X-Public"
)

type oauthClient struct{}

type oauthInterface interface{}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}
