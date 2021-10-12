package main

//struct una manera de tener los datos agrupados

type myType struct {
	myString string
	myBool   bool
	myInt    int
}

func (myType) myFunction() {

}

func main() {
	// var myBool bool
	// myBool2 := false
	// tiene el bit mas significativo como signo

	// Signed
	// var myInt int8   //-128 ~ 127
	// var myInt2 int16 // -32768 ~32767
	// var myInt3 int32
	// var myInt3 int // could be int32 or int64
	// var myInt3 int64

	// Unsigned
	// var myUInt uint8 // 0-255
	// var myByte byte  // is unit8 , is like char using ASCII
	// var myUInt2 uint32
	// var myRune rune // char UNICODE
	// var myUInt3 uint64

	// float
	// var myFloat float32
	// var myFloat2 float64

	//String
	// var myString string = "una cadena de caracteres"
	// var myString2 string = `

	//multilinea
	// `

	// var myDeclare myType
	// myDeclare.myBool
}
