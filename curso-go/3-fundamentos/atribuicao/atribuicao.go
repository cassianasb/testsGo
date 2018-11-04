package main

import (
	"fmt"
)

func main() {
	var b byte = 3 // atribuir variavel informando o tipo
	fmt.Println(b)

	i := 3 //inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 3
	i *= 2 // i = i * 3
	i %= 2 // i = i % 3

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x //inverter os valores entre x e y
	fmt.Println(x, y)
}
