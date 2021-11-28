package zapper

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestLogMiddleware(t *testing.T) {

	fooHandler := func(w http.ResponseWriter, r *http.Request) {
		Error("LogMiddleware", zap.String("net/http middleware", "net/http的handler打印"))
		w.Write([]byte("I'm OK"))
	}

	http.Handle("/foo", LogMiddleware(http.HandlerFunc(fooHandler)))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
