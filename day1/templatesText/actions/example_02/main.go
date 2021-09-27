package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Score uint32
}

const Msg = `
{{.Name}}, ваш результат: {{.Score}}
и ваш уровень: {{if le .Score 50}}Низкий
{{else if le .Score 80}}Средний
{{else}}Высокий
{{end}}
`

func main() {
	u1 := User{"Петр", 30}
	u2 := User{"Анна", 80}

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
