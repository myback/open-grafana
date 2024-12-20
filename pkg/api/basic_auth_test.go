package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/macaron.v1"
)

func TestBasicAuthenticatedRequest(t *testing.T) {
	const expectedUser = "prometheus"
	const expectedPass = "password"

	t.Run("Given a valid set of basic auth credentials", func(t *testing.T) {
		httpReq, err := http.NewRequest("GET", "http://localhost:8080/metrics", nil)
		require.NoError(t, err)
		req := macaron.Request{
			Request: httpReq,
		}
		encodedCreds := encodeBasicAuthCredentials(expectedUser, expectedPass)
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
		authenticated := BasicAuthenticatedRequest(req, expectedUser, expectedPass)

		assert.True(t, authenticated)
	})

	t.Run("Given an invalid set of basic auth credentials", func(t *testing.T) {
		httpReq, err := http.NewRequest("GET", "http://localhost:8080/metrics", nil)
		require.NoError(t, err)
		req := macaron.Request{
			Request: httpReq,
		}
		encodedCreds := encodeBasicAuthCredentials("invaliduser", "invalidpass")
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
		authenticated := BasicAuthenticatedRequest(req, expectedUser, expectedPass)

		assert.False(t, authenticated)
	})
}

func encodeBasicAuthCredentials(user, pass string) string {
	creds := fmt.Sprintf("%s:%s", user, pass)
	return base64.StdEncoding.EncodeToString([]byte(creds))
}
