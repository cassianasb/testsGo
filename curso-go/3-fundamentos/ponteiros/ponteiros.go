package main

import (
	"fmt"
)

func main() {
	i := 1 //Inteiro = 8 bites na memoria

	var p *int = nil //nome do ponteiro *tipo do ponteiro (64bits de memoria - int) = nulo
	p = &i           //endereço da vairiavel i
	*p++             //desrererenciando  (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)

	// GO não tem aritmeetica de ponteiros
	//p++ nao funciona
}
