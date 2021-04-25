package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const DefaultPort = "8990"
const BackendUrl = "http://127.0.0.1:8991/"

type SimpleProxy struct {
	RevProxy *httputil.ReverseProxy
}

type ProxyHandlers interface {
	handler(http.ResponseWriter, *http.Request)
	proxyHandler(http.ResponseWriter, http.Request)
}

func NewSimpleProxy() *SimpleProxy {
	bkendUrl, _ := url.Parse(getEnviron("BackendURL", BackendUrl))
	p := httputil.NewSingleHostReverseProxy(bkendUrl)
	return &SimpleProxy{
		RevProxy: p,
	}
}

func (p *SimpleProxy) handler(w http.ResponseWriter, r *http.Request) {
	p.RevProxy.ServeHTTP(w, r)
}

func getEnviron(env string, def string) string {
	value := os.Getenv(env)
	if value == "" {
		value = def
	}
	return value
}

func (p *SimpleProxy) proxyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've reached the proxy!\n")
}

func StartProxy() {
	proxyAddr := fmt.Sprintf(":%s", DefaultPort)
	p := NewSimpleProxy()
	proxyServer := http.NewServeMux()
	proxyServer.HandleFunc("/", p.handler)
	proxyServer.HandleFunc("/proxy", p.proxyHandler)

	http.ListenAndServe(proxyAddr, proxyServer)
}
