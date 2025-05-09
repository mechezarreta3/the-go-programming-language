package main

import (
	"log"
	"net/http"

	"github.com/mechezarreta3/the-go-programming-language/chapter1/examples/lissajous/lissajous"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}
