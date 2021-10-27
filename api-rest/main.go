package main

import (
	"api-rest/config"
	"fmt"
	"log"
	"net/http"
)

// go mod init api rest crea un nuevo go mod  = package.json
// go get https://github.com/joho/godotenv para instalar paquetes
func main() {

	conf := config.NewConfiguration()

	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello there")
	})

	fmt.Println("Server listening on ", fmt.Sprintf(":%d", conf.Port))
	// abre un puerto para escuchar
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil),
	)
	fmt.Println("hello word 2")
}
