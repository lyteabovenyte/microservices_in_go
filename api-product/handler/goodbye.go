package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func newGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) serveHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("received by the goodbye handler")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		g.l.Println("couldn't read request body")
		http.Error(rw, "request body in awful.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "received by goodbye handle, MR.%v", d)
}
