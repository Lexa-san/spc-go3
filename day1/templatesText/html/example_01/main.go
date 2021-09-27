package main

import (
	"html/template"
	"os"
)

const Page = `
<html>
<head>
    <title>Языки программирования, которые знает {{.Name}}: </title>
</head>
<body>
	<h3>Языки программирования, которые знает {{.Name}}: </h3>
    <ul>
    {{range .Languages}}<li>{{print .}}</li>{{end}}
    </ul>
</body>
</html>
`

type UserExperience struct {
	Name      string
	Languages []string
}

func main() {
	languages := []string{"Go", "C++", "C#"}
	u := UserExperience{"User", languages}

	t := template.Must(template.New("web").Parse(Page))

	err := t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
