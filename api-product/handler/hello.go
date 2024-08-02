package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func newHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) serveHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("received from hello handler.")
	d, err := io.ReadAll(r.Body); 
	if err != nil {
		h.l.Println("couldn't read the request's body")
		http.Error(rw, "request's body is invalid", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "it received by the hello handler. MR.%s", d)
}
