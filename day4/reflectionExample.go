package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
	Quarter
}

type Quarter int

func main() {
	n := Quarter(2)
	p1 := Point{X: 1, Y: -1, Quarter: n}

	//r1 := reflect.ValueOf(&p1).Elem()
	r1 := reflect.ValueOf(p1)
	fmt.Println(r1.Type(), r1.NumField(), r1.NumMethod())
	fmt.Println(reflect.ValueOf(&p1).Elem())

	for i := 0; i < r1.NumField(); i++ {
		fmt.Printf("Field name: %v,\n", r1.Type().Field(i).Name)
		fmt.Printf("Field type: %v,\n", r1.Field(i).Type())
		// panic wil be accured if a field is not exported (private)
		fmt.Printf("Field value: %v\n", r1.Field(i).Interface())
	}

}
