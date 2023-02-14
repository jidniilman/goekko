// Package main is the entry point of the application goekko
package main

import (
	"jidniilman/goekko/configs"
	"jidniilman/goekko/internal/server"
	"jidniilman/goekko/internal/server/routes"
	"log"
	// "jidniilman/goekko"
)

// entry point of goekko application
func main() {
	// get the configs and start an instance of echo web server
	cfg := configs.GetConfig()
	goekkoServer := server.NewGoekkoServer(cfg)

	routes.ConfigureRoutes(goekkoServer)

	err := goekkoServer.Start(cfg.GoekkoAppConfig.AppName, "1330")
	if err != nil {
		log.Fatal(err)
	}

}
