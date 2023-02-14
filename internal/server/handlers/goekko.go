package handlers

import (
	"github.com/labstack/echo/v4"
	"jidniilman/goekko/internal/server"
	"jidniilman/goekko/internal/services/hello"
	"net/http"
)

type HelloHandlers struct {
	server *server.Server
}

func NewHelloHandlers(server *server.Server) *HelloHandlers {
	return &HelloHandlers{server: server}
}

func (h *HelloHandlers) Hello(c echo.Context) error {
	hello := hello.NewHelloService()
	hello.Hello()
	return c.String(http.StatusOK, "Goekko is Online!")
}
