package main

import "fmt"

func main() {
	s := make([]int, 10) //criou um array interno para referenciar
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // slice com 10 elementos mas um array interno de 20 elementos

	fmt.Println(s, len(s), cap(s)) // slice, tamanho do slice, capacidade do slice

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) // mesmo nro da capacidade do array interno

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // a capacidade dobrou de tamanho quando atinge a capacidade do array interno
}
