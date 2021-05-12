package oauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOauthConstants(t *testing.T) {
	assert.EqualValues(t, "X-Public", headerXPublic)
	assert.EqualValues(t, "X-Client-Id", headerXClientId)
	assert.EqualValues(t, "X-User-Id", headerXUserId)
	assert.EqualValues(t, "access_token", paramAccessToken)
}
