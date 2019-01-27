package main

import "fmt"

func media(numeros ...float64) float64 { //... indica funcção variatica, numero variavel de parametros (será um array)
	total := 0.0
	for _, num := range numeros { //_ = indice, num sã os numeros
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.2, 10.0, 5.6))
}
