package main

import "fmt"

//struct una manera de tener los datos agrupados

type structure struct {
	chasisSize int8
	hasBumper  bool
	pieces     uint32
}

type motor struct {
	cilinders uint16
	sparkPlug int
	hasBand   bool
}

type aesthetic struct {
	mainColor      string
	secondaryColor string
	wheelsColor    string
}

type car struct {
	chasis    structure
	mainMotor motor
	color     aesthetic
}

func monsterTruck() *structure {

	monsterTruckInstance := &structure{
		chasisSize: 20,
		hasBumper:  true,
		pieces:     200,
	}
	return monsterTruckInstance
}

func V8() *motor {

	motorInstance := &motor{
		cilinders: 8,
		sparkPlug: 8,
		hasBand:   true,
	}
	return motorInstance
}

func punk() *aesthetic {

	aestheticInstance := &aesthetic{
		mainColor:      "#000099",
		secondaryColor: "#595959",
		wheelsColor:    "#000000",
	}
	return aestheticInstance
}

func torettoCar() *car {

	newCar := &car{
		chasis:    *monsterTruck(),
		mainMotor: *V8(),
		color:     *punk(),
	}
	return newCar
}

func main() {
	var myCar = torettoCar()
	fmt.Printf("%v", myCar.chasis.chasisSize)
}
