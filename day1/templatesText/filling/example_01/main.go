package main

import (
	"os"
	"text/template"
)

type User struct {
	Name   string
	UserId string
	Email  string
}

const Msg = `Здравствуйте, {{.Name}},
Вы зарегистрировались с id {{.UserId}}
и e-mail {{.Email}}.
`

func main() {
	u := User{"User", "0001", "user@mail.com"}
	t := template.Must(template.New("msg").Parse(Msg))
	err := t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
