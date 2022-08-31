package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hi")
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Image")
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Thumbnail Image")
		if err != nil {
			return
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
