package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta o tamanho

	for i, numero := range numeros { //index, variavel do elemento do array
		fmt.Printf("%d) %d \n", i+1, numero)
	}

	for num := range numeros { //apenas o index e não o elemento
		fmt.Println(num)
	}

	for _, num := range numeros { //sem nessidade do index (variavel)
		fmt.Println(num)
	}
}
