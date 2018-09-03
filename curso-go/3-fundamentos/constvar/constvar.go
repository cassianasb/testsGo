package main

import (
	"fmt"
	"math" //colocar um underline antes desativa o pacote sem retira-lo do import
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64 inferido pelo compilador

	//forma reduzida
	area := PI * math.Pow(raio, 2) //declarando e inicializando a variavel

	fmt.Println("A área é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
