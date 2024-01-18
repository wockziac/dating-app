package jwt

import (
	"datingapp/application/core"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TOKEN_PERIOD = 24 // HOUR
	ISSUER       = "dating-app"
)

var (
	SIGNING_KEY = []byte("signingKey")
)

type JWTProvider struct {
}

type CustomClaims struct {
	UserID    int `json:"userID"`
	SessionID int `json:"sessionID"`
	jwt.RegisteredClaims
}

func (j JWTProvider) GenerateToken(session core.Session) (core.SessionToken, error) {
	expiredAt := time.Now().Add(TOKEN_PERIOD * time.Hour)

	jwtClaims := CustomClaims{
		UserID:    session.Credential.UserID,
		SessionID: session.SessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ISSUER,
			Subject:   ISSUER,
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Audience:  []string{fmt.Sprint(session.Credential.UserID)},
			ID:        "1",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	signedToken, err := token.SignedString(SIGNING_KEY)
	if err != nil {
		return core.SessionToken{}, fmt.Errorf("error generate token")
	}

	return core.SessionToken{
		Token:      signedToken,
		ExpiryTime: expiredAt,
	}, nil
}
