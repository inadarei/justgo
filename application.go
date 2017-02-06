package main

import (
	"log"
	"os"

	"github.com/inadarei/justgo/server"
)

func main() {

	// isDevelopment := os.Getenv("APP_ENV") == "development"

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "3737"
		log.Print("WARNING: no server port supplied in the environment. Defaulting to " + serverPort)
	}

	server.StartServer(serverPort)
}

// @see: https://golang.org/doc/code.html
