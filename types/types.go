package types

import "fmt"

type motor struct {
	power uint16
}

type car struct {
	motor motor
	color string
}

func newCar() *car {
	// var myCar car
	// myCar.color = "white"
	// myCar := car{
	// 	color: "white",
	// 	motor: motor{
	// 		power: 160,
	// 	},
	// }

	// declaras una nueva direccion de memoria
	myCar := new(car)
	fmt.Printf("%p", myCar)
	fmt.Printf("%v", myCar.color)
	return myCar
}

func newCar2() *car {
	//nueva direccion de memoria no es una instancia
	myCar := &car{
		color: "white",
		motor: motor{
			power: 160,
		},
	}
	return myCar
}

func createCar() {

	// del puntero new card asigna un color red
	myCar := newCar2()
	myCar.color = "red"
}
