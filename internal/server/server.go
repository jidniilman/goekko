package server

import (
	"fmt"
	"jidniilman/goekko/configs"
	"jidniilman/goekko/pkg/gql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/machinebox/graphql"
)

type Server struct {
	Echo    *echo.Echo
	GraphQL *graphql.Client
	Config  *configs.Config
}

func NewGoekkoServer(cfg *configs.Config) *Server {
	return &Server{
		Echo:    echo.New(),
		GraphQL: gql.Init(cfg),
		Config:  cfg,
	}
}

func (server *Server) Start(appName string, port string) error {
	// Middleware
	server.Echo.Use(middleware.CORS())
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Use(middleware.Gzip())
	server.Echo.HideBanner = false

	fmt.Println("Starting Server " + appName)
	return server.Echo.Start(":" + port)
}
