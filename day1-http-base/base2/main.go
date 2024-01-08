package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

// package http
//
// type Handler interface {
//     ServeHTTP(w ResponseWriter, r *Request)
// }
//
// func ListenAndServe(address string, h Handler) error

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
