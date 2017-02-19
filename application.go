// Note: the comment below after 'main' is a hint to go-wrapper and is meaningful
// Replace with the package path for your project
package main // import "github.com/inadarei/justgo"

import (
	"log"
	"os"

	"github.com/inadarei/justgo/server"
)

func main() {

	log.Print("Application starting up...")

	// isDevelopment := os.Getenv("APP_ENV") == "development"

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "3737"
		log.Print("WARNING: no server port supplied in the environment. Defaulting to " + serverPort)
	}

	server.StartServer(serverPort)
}

// @see: https://golang.org/doc/code.html
