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

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("bye"))

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oups", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "goodbye %s\n", d)
}
