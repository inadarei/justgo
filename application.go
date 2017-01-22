package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "3737"
		log.Print("WARNING: no server port supplied. Defaulting to " + serverPort)
	}

	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		log.Fatal("ERROR: couldn't start server: ", err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")

	fmt.Fprintf(w, "Hello world! This was easy!")
}
