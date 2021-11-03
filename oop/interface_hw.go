package oop

import "fmt"

type PlayStation interface {
	TurnOn()
	Videogame(name string)
	Movie(func(movie string))
}

func PlayMovie(movie string) {
	if movie != "" {
		fmt.Printf("playing: %s", movie)
	}
}

type PS struct{}

func (PS) TurnOn() {
	fmt.Println("turn on, please stand by...")

}

func (PS) Videogame(name string) {
	if name != "" {
		fmt.Printf("playing: %s", name)
	}
}

func (PS) Movie(func(movie string)) {

}

func MyPS3(ps3 PS) {
	ps3.TurnOn()
	ps3.Videogame("God Of War")

}

func MyPS4(ps4 PS) {
	ps4.TurnOn()
	ps4.Videogame("God Of War Remake")
}

// anonymous function

func rappi(sendShopper func(thing string) string) string {
	res := sendShopper("milk")
	return res
}

func run() {
	rappi(func(thing string) string {

		// todo: verify if thing exist in mall
		if thing != "" {
			return "Sucess"
		} else {
			return "Failure"
		}
	})
}
