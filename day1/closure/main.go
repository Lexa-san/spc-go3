package main

// Замыкания в Golng
import "fmt"

func scope(value int) func() int {
	outerVar := 10
	foo := func() int {
		return outerVar + 2
	}
	bar := func() int {
		return outerVar * 2
	}
	if value > 10 {
		return foo
	} else {
		return bar
	}
}

// Прееопределяем переменную outerVar внутри функции inner
func outer() (func() int, int) {
	outerVar := 2
	inner := func() int {
		outerVar += 99
		return outerVar
	}
	return inner, outerVar // => 101, 2
}

// не скомпилируется, потому что эти переменные здесь недоступны
//func anotherScope() int {
//	outerVar = 123;
//	return foo
//}

func main() {

	resFunc := scope(5)
	fmt.Println(resFunc()) // => 20

	//fmt.Println(scope()())

	funcRes, res := outer()
	fmt.Println(funcRes(), res)

}
