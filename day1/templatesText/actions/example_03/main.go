package main

import (
	"os"
	"text/template"
)

const msg = `
Мушкетеры короля это:
{{range .}}{{print .}} {{end}}
`

func main() {
	musketeers := []string{"Атос", "Портос", "Арамис", "Д`Артаньян"}

	t := template.Must(template.New("msg").Parse(msg))
	err := t.Execute(os.Stdout, musketeers)
	if err != nil {
		panic(err)
	}
}
