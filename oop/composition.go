package oop

import "fmt"

type product struct {
	Name  string
	Price float32
	productDetails
}

type productDetails struct {
	genre   string
	reviews string
	color   string
}

func (p productDetails) showDetails() {
	fmt.Println("Color: ", p.color)
}

func run() {
	v := product{
		Name: "Coca",
	}
	v.reviews = "Hola"
	v.showDetails()
}

type qoutation interface {
	SendQuotation()
}

type demostration interface {
	DemostrateProduct()
}

type isalesMan interface {
	qoutation
	demostration
}

type Salesman struct {
}

func (s Salesman) SendQuotation() {

}

func (s Salesman) DemostrateProduct() {

}

func hireSalesman(salesman isalesMan) {
	salesman.SendQuotation()
	salesman.DemostrateProduct()
}

func getQuotation(quote qoutation) {

}

func run2() {
	s := Salesman{}
	hireSalesman(s)
	getQuotation(s)
}
