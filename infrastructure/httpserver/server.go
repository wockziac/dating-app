package httpserver

import (
	"datingapp/application/core"
	"datingapp/infrastructure/httpserver/jwt"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Dependency
}

type Dependency struct {
	CredentialRepo       core.ICredentialRepository
	CredentialUseCase    core.ICredentialUseCase
	SessionRepo          core.ISessionRepository
	SessionUseCase       core.ISessionUseCase
	SessionTokenProvider jwt.JWTProvider
}

func NewHttpServer(dep Dependency) *HttpServer {
	return &HttpServer{
		Dependency: dep,
	}
}

func (s *HttpServer) Init() {

	r := gin.Default()
	SetupRoute(s.Dependency, r)
	r.Run()
}
