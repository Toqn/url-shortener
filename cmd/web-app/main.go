package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ShowHomPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

func ShowHomPage(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
