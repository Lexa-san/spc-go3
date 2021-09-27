package main

/*
Создание пользовательских функций с помощью template.FuncMap

Последний способ вызова наших собственных функций - это создание пользовательских функций с шаблоном template.FuncMap.
Это наиболее полезный и мощный способ определения функций, поскольку он позволяет нам создавать глобальные вспомогательные методы,
которые можно использовать в нашем приложении.

Чтобы начать, сначала посмотрим на template.FuncMap.
Первое, на что нужно обратить внимание, это то, что этот тип выглядит как map[string]interface{}, но ниже есть примечание,
что каждый интерфейс должен быть функцией с одним возвращаемым значением или функцией с двумя возвращаемыми значениями,
где первое, это данные, к которым вам нужно получить доступ в шаблоне,
а второе, ошибка, которая прекратит выполнение шаблона, если он не равен нулю.
*/

import (
	"html/template"
	"net/http"
)

var templOne *template.Template

type ViewData struct {
	User UserT
}

type UserT struct {
	ID    int
	Email string
}

/*
Определите функции перед разбором шаблонов

В предыдущих примерах мы создавали наш шаблон, вызывая функцию template.ParseFiles,
предоставляемую пакетом html/template. Это функция уровня пакета, которая возвращает
шаблон после анализа файлов. Теперь мы вызываем метод ParseFiles для типа template.Template,
который имеет те же возвращаемые значения, но применяет изменения к существующему шаблону
(а не к новому) и затем возвращает результат.

В этой ситуации нам нужно использовать метод, потому что нам нужно сначала определить
любые пользовательские функции, которые мы планируем использовать в наших шаблонах,
и как только мы сделаем это с пакетом шаблонов, он вернет *template.Template.
После определения этих пользовательских функций мы можем приступить к анализу шаблонов,
которые используют функции. Если бы мы сначала проанализировали шаблоны, вы бы увидели ошибку,
связанную с неопределенной функцией, вызываемой в вашем шаблоне.
*/

func main() {
	var err error
	templOne, err = template.New("tmplEx5.html").Funcs(template.FuncMap{
		"hasPermission": func(user UserT, feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).ParseFiles("tmplEx5.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handlerOne)
	http.ListenAndServe(":3000", nil)
}

func handlerOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	user := UserT{
		ID:    1,
		Email: "user@email.io",
	}
	vd := ViewData{user}
	err := templOne.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
