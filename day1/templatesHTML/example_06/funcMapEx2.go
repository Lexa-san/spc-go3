package main

/*
Делаем наши функции глобальными во всем приложении

Функция hasPermission, которую мы определили в последнем разделе, отлично работает, но одна проблема
с ней заключается в том, что мы можем использовать ее только тогда, когда у нас есть доступ к объекту User.
На первый взгляд, это не так уж плохо, но по мере роста приложения у него будет много шаблонов,
и довольно легко забыть передать объект User в шаблон или пропустить его во вложенном шаблоне.

Наша функция была бы намного проще, если бы мы могли ее упростить, и нам нужно было только передать
имя функции, поэтому давайте продолжим и обновим наш код, чтобы это произошло.

Первое, что нам нужно сделать, это создать функцию, котора не принимает объект User.
Мы установим это в template.FuncMap перед синтаксическим анализом нашего шаблона,
чтобы у нас не было ошибок синтаксического анализа, и чтобы убедиться, что у нас есть
некоторая логика в случае, если пользователь недоступен.
*/
import (
	"html/template"
	"net/http"
)

var templTwo *template.Template

type ViewDataUsers struct {
	User UserType
}

type UserType struct {
	ID    int
	Email string
}

func main() {
	var err error
	templTwo, err = template.New("tmplEx6.html").Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			return false
		},
	}).ParseFiles("tmplEx6.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handlerTwo)
	http.ListenAndServe(":3000", nil)
}

/*

Далее нам нужно определить нашу функцию, которая использует замыкание.
Мы собираемся определить динамическую функцию, которая имеет доступ к переменным,
которые не обязательно передаются в нее, но доступны, когда мы определяем функцию.
В нашем случае эта переменная будет объектом User.

Несмотря на то, что мы определили функцию hasPermission в нашей функции main(),
мы перезаписываем ее внутри нашего обработчика, когда имеем доступ к объекту User,
но перед тем, как выполнить шаблон. Это действительно мощно, потому что теперь мы можем
использовать функцию hasPermission в любом шаблоне, не беспокоясь о том, был ли передан
объект User в шаблон или нет.
*/

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	user := UserType{
		ID:    1,
		Email: "user@email.io",
	}
	vd := ViewDataUsers{user}
	err := templTwo.Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
