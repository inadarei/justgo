package server

import (
	"fmt"
	"log"
	"net/http"
)

/**
* Main entrance into the server.
 */
func StartServer(serverPort string) {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		log.Fatal("ERROR: couldn't start server: ", err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")

	fmt.Fprintf(w, "Hello world! This was pretty awesome! Hot-reloading is the thing!")
}
