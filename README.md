# reverse-proxy
Exploring the implementation of a reverse proxy by running a `SimpleReverProxy` running in front of a simple http server.

### Deployment
```
go run cmd/reverse-proxy/main.go
Starting the http server..
Starting the proxy server..
```

Proxy forwarding connections to the backend HTTP server (and relaying the result back to the client):
```
curl "http://127.0.0.1:8990/"
You've reached the HTTP server!
```

Proxy intercepting connections to "/proxy":
```
curl "http://127.0.0.1:8990/proxy"
You've reached the proxy!
```
