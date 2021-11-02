package arrays

import "fmt"

func defineArrays() {
	var a [2]string // siempre que declaramos un array le ponemos el tamaño
	a[0] = "ab"
	a[1] = "cd"
	fmt.Println(a[0])

	primos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primos)

	// slice es un wrapper para un array,  son dinamicos y flexibles

	var s []string

	s2 := primos[1:4]
	// si quiers clonar todo el array
	s3 := primos[:]

	// slice siempre refencia al primer elemento, si no le indicamos cual

	// length of slice
	len(s3) // length
	cap(s3) // capacity

	var s4 []int
	fmt.Println(s4)

	if s4 == nil {
		fmt.Println("nil")
	}

	// make crea algo , primero tipo y luego longitud
	a1 := make([]int, 5)
	// primero tipo, luego longitud, luego capacidad
	a2 := make([]int, 0, 5)

	//array de arrays

	matrix := [][]float32{
		[]float32{1.1, 1.2, 1.3},
		[]float32{2.1, 2.2, 2.3},
	}

	s4 = append(s4, 2)
}
