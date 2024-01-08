package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
		if err != nil {
			return
		}
	})

	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			_, err := fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}
	})

	err := r.Run(":9999")
	if err != nil {
		return
	}
}
