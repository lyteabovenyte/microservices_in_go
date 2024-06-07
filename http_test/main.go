package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listening...")

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("hello world")
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}