package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	//return segundo, primeiro
	return //retorno limpo, uma vez que o retorno j á foi declarado na declaração da função
}

func main() {
	//fmt.Println(troca(1, 2))
	x, y := troca(2, 3)
	fmt.Println(x, y)
}
