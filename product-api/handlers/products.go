package handler

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/lyteabovenyte/microservices_in_go/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// implementing Handler interface with the ServeHTTP method
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "cannot convert to json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
