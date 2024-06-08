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

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// implementing the Handler interface for Hello handler
// Handler interface just define a single method called ServeHTTP
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world") // for the server

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oups", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s\n", d) // for the client
}
