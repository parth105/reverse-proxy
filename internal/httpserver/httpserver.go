package httpserver

import (
	"fmt"
	"net/http"
)

const HttpServerPort = "8991"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've reached the HTTP server!\n")
}

func StartHttpServer() {
	httpServer := http.NewServeMux()        // Using NewServeMux allows us to overcome the multiple registrations
	httpServer.HandleFunc("/", httpHandler) // recieved error when using defaultservemux (http.HandleFunc)
	serverAddr := fmt.Sprintf(":%s", HttpServerPort)
	http.ListenAndServe(serverAddr, httpServer)
}
