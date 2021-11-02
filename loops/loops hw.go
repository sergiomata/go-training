package loops

import (
	"errors"
	"fmt"
)

func semaforo(status int) string {

	if status == 1 {
		return "go"
	} else if status == 2 {
		return "warning"
	} else {
		return "not move"
	}

}

type CustomError struct {
	message string
}

type User struct {
	id   string
	name string
}

func GetUser(url string) (interface{}, error) {

	if url == "" {
		return nil, errors.New("url not found")
	}

	res := &User{
		id:   "qwerqwe",
		name: "paquito",
	}
	return res, nil
}

type cakeIngredients struct {
	typeCake string
}

type cake struct {
	typeCake string
}

func makeCake(isIndividual bool, flavour string, ingredients interface{}) interface{} {

	res := &cake{
		typeCake: "noCake",
	}

	switch flavour {
	case "string":
		res.typeCake = "tres leches"
		break
	case "int":
		res.typeCake = "hojuelas"
		break
	default:
		res.typeCake = "chocolate"
		break
	}
	return res
}

func PrintAllFields(arg ...string) {
	for _, elem := range arg {
		fmt.Println(elem)
	}
}
