package functions

import (
	"errors"
	"fmt"
)

type MyType struct {
	I int
	S string
}

type MyCustomError struct{}

func (MyCustomError) Error() string {
	return "Custom error"
}

func MyFunction(arg1 int, arg2 string, arg3 interface{}) {

}

func myFunc() {
	MyFunction(2, "3", MyType{
		I: 1,
		S: "sdf",
	})
	variadic(1, 2, 3, 4, 5, 6)
	a1 := []int{1, 2, 3, 4, 5}
	variadic(a1...)

	m1, _ := multipleReturn()

	r, err := funcError()

	if err != nil {
		panic(r)
	}

}

func variadic(arg ...int) {
	for _, el := range arg {
		fmt.Println(el)
	}
}

func returnFunction() int {
	return 0
}

func multipleReturn() (int, string) {
	return 0, ""
}

func funcError() (interface{}, error) {
	return nil, errors.New("new error")
}
