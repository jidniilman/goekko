package routes

import (
	"jidniilman/goekko/internal/server"
	"jidniilman/goekko/internal/server/handlers"
)

func ConfigureRoutes(server *server.Server) {

	helloHandler := handlers.NewHelloHandlers(server)

	// default route
	server.Echo.GET("/", helloHandler.Hello)
}
