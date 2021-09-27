package main

/*
call - это функция, уже предоставленная пакетом html/template, которая вызывает первый
переданный ей аргумент (в нашем случае - функцию .User.HasPermission),
используя остальные аргументы в качестве аргументов вызова функции.
*/

import (
	"html/template"
	"net/http"
)

var tstTemplate *template.Template

type UserData struct {
	User UserAccount
}

type UserAccount struct {
	ID            int
	Email         string
	HasPermission func(string) bool
}

func main() {
	var err error
	tstTemplate, err = template.ParseFiles("tmplEx4.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", Handler4)
	http.ListenAndServe(":3000", nil)
}

func Handler4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	ud := UserData{
		User: UserAccount{
			ID:    1,
			Email: "user@email.io",
			HasPermission: func(feature string) bool {
				if feature == "feature-b" {
					return true
				}
				return false
			},
		},
	}
	err := tstTemplate.Execute(w, ud)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
