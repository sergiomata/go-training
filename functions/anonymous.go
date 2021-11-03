package functions

import "fmt"

func calculate(fun func(a, b int) int) {
	res := fun(0, 1)
	fmt.Println(res)
}

func run() {
	calculate(func(a, b int) int {
		return a + b
	})

	myFunc := returnFunc()
	myFunc(0, 5)

	func() {

	}()
}

func returnFunc() func(a, b int) int {
	return func(a, b int) int {
		return 0
	}
}
