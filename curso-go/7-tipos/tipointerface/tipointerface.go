package main

import "fmt"

type curso struct {
	nome string
}

//tipo interface é dinamico, object Java
func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(3)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do Google"}
	fmt.Println(coisa2)
}
