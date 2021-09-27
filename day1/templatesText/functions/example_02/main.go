package main

import (
	"os"
	"strings"
	"text/template"
)

const Msg = `
Мушкетеры короля это:
{{join . ", "}}
`

func main() {
	musketeers := []string{"Атос", "Портос", "Арамис", "Д`Артаньян"}

	funcs := template.FuncMap{"join": strings.Join}

	t, err := template.New("msg").Funcs(funcs).Parse(Msg)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, musketeers)
	if err != nil {
		panic(err)
	}
}
