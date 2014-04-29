package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
)

func main() {
    listento := ":8080"

    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true

    println("HTTP proxy server listening on " + listento)

    log.Fatal( http.ListenAndServe(listento, proxy ) )
}
