package main

import (
	"github.com/Toqn/url-shortener/internal/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/shorten", controllers.ShowShorten)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
