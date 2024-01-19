package jwt

import (
	"datingapp/application/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {

	t.Run("success generate token", func(t *testing.T) {
		provider := JWTProvider{}
		session := core.Session{
			Credential: core.Credentials{
				UserID: 1,
			},
			SessionID: 12,
		}
		sessionToken, err := provider.GenerateToken(session)
		assert.Nil(t, err)
		assert.NotEqual(t, "", sessionToken.Token)
	})
}

func TestIsTokenValid(t *testing.T) {
	t.Run("token invalid", func(t *testing.T) {
		provider := JWTProvider{}
		valid := provider.IsTokenValid("ABCDEasdlkjfe")
		assert.Equal(t, false, valid)

	})
}
