package main

import (
	"log"
	"net/http"
	"os"
	"github.com/lyteabovenyte"
)

func main() {

	// creating log for the handlers
	l := log.New(os.Stdout, "product-api", log.Lmsgprefix)

	// creating the handlers
	helloHandler := handler.newHello(l)
	goodbyeHandler := handler.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", handler.Hello)
	sm.Handle("/goodbye", handler.Goodbye)

	http.ListenAndServe(":9000", nil)
}
