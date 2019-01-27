package main

import "fmt"

var soma = func(a, b int) int { //função anonima mas já com variável declarada com o retorno da função
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int { //possível fazer dentro da main
		return a - b
	}

	fmt.Println(sub(2, 3))
}
