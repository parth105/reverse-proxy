package main

import (
	"github.com/parth105/reverse-proxy/internal/httpserver"
	"github.com/parth105/reverse-proxy/internal/proxy"

	"fmt"
)

const HttpServerPort = "8991"

func main() {
	fmt.Println("Starting the http server..")
	go httpserver.StartHttpServer()

	fmt.Println("Starting the proxy server..")
	proxy.StartProxy()
}
