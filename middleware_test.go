package zapper

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"testing"
)

func TestLogMiddleware(t *testing.T) {
	

	fooHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am ok"))
	}

	http.Handle("/foo", LogMiddleware(http.HandlerFunc(fooHandler)))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
