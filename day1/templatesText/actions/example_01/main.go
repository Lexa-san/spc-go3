package main

import (
	"os"
	"text/template"
)

type User struct {
	Name   string
	Female bool
}

// {{- end}} закрывает блок {{if}}
const Msg = `
{{if .Female}}Дорогая,{{else}}Дорогой,{{end}} {{.Name}},
Ваш заказ готов.
Спасибо за обращение!
`

func main() {
	u1 := User{"Петр", false}
	u2 := User{"Анна", true}

	t := template.Must(template.New("msg").Parse(Msg))
	err := t.Execute(os.Stdout, u1)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, u2)
	if err != nil {
		panic(err)
	}
}
