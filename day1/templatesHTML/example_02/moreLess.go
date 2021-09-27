package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

var tmpl *template.Template

type Status struct {
	Usage   int
	Limit   int
	Warning int
}

func main() {
	var err error
	tmpl, err = template.ParseFiles("tmplEx2.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := tmpl.Execute(w, &Status{Usage: rand.Intn(30), Limit: 30, Warning: 20})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
