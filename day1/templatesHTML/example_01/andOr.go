package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type User struct {
	Admin bool
}

type ViewData struct {
	*User
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("tmplEx1.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{&User{true}}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
