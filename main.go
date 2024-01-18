package main

import (
	"datingapp/application/usecase"
	"datingapp/infrastructure/db/inmemory"
	"datingapp/infrastructure/httpserver"
	"datingapp/infrastructure/httpserver/jwt"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	credentialRepo := inmemory.NewCredentialRepository()
	sessionRepo := inmemory.NewSessionRepository()

	jwtProvider := jwt.JWTProvider{}

	sessionUseCase := usecase.NewSessionUseCase(sessionRepo, credentialRepo, jwtProvider)
	credentialUseCase := usecase.NewCredentialsUsecase(credentialRepo, sessionUseCase)

	httpServer := httpserver.NewHttpServer(httpserver.Dependency{
		CredentialRepo:       credentialRepo,
		SessionRepo:          sessionRepo,
		SessionTokenProvider: jwtProvider,
		SessionUseCase:       sessionUseCase,
		CredentialUseCase:    credentialUseCase,
	})
	httpServer.Init()
}
