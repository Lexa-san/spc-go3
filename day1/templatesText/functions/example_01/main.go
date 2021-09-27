package main

import (
	"os"
	"text/template"
)

//то же самое, что и x[3]
const Msg = `
Мушкетеры короля это:
{{slice . 3}} 
`

func main() {
	musketeers := []string{"Атос", "Портос", "Арамис", "Д`Артаньян"}

	t := template.Must(template.New("msg").Parse(Msg))

	err := t.Execute(os.Stdout, musketeers)
	if err != nil {
		panic(err)
	}
}
