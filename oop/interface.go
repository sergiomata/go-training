package oop

import "fmt"

type MyInterface interface {
	Method1()
	Method2(a int)
	Method3(a, b float32) float32
}

// cualqiuer objecto puede ser implementado en metodo interface
// solo tienes que castearlo
interface{

}

type MyType struct{}

func (MyType) Method1() {
	fmt.Println("Hola,mundo")
}

func (MyType) Method2(a int) {
	fmt.Println("Hola,mundo 2")
}

func (MyType) Method3(a, b float32) float32 {
	fmt.Println("Hola,mundo 3")
	return 0
}

func IntImplementation(i MyInterface) {
	i.Method1()
	intExample(10)
}

func intExample (i interface{})

func run() {
	IntImplementation(MyType{})
}


func intExample (i interface{}){
	v := i.Method1()
}