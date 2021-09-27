package main

import (
	"html/template"
	"net/http"
)

var testTmpl *template.Template

type Data struct {
	User UserAcc
}

type UserAcc struct {
	ID    int
	Email string
}

func (u UserAcc) HasPermission(feature string) bool {
	if feature == "feature-a" {
		return true
	} else {
		return false
	}
}

func main() {
	var err error
	testTmpl, err = template.ParseFiles("tmplEx3.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", HandlerF)
	http.ListenAndServe(":3000", nil)
}

func HandlerF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	vd := Data{
		User: UserAcc{1, "user@email.io"},
	}
	err := testTmpl.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
