package middlewares

import (
	"datingapp/infrastructure/httpserver/jwt"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AUTH_KEY = "Authorization"
	BEARER   = "Bearer "
)

var URL_EXCEPTION_LIST = []string{
	"/v1/credentials/register",
	"/v1/auth/initiate-login",
	"/v1/auth/complete-login",
}

type MessageResponse struct {
	Message string `json:"message"`
}

func AuthMiddleware(ctx *gin.Context) {
	uri := ctx.Request.RequestURI
	isExceptionURI := slices.Contains(URL_EXCEPTION_LIST, uri)
	if isExceptionURI {
		ctx.Next()
		return
	}

	bearerToken := ctx.Request.Header.Get(AUTH_KEY)
	fmt.Println(bearerToken)
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, MessageResponse{
			Message: "missing token",
		})
		ctx.Abort()
		return
	}

	jwtProvider := jwt.JWTProvider{}
	token, _ := strings.CutPrefix(bearerToken, BEARER)
	if !jwtProvider.IsTokenValid(token) {
		ctx.JSON(http.StatusUnauthorized, MessageResponse{
			Message: "invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
