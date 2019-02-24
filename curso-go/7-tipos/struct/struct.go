package main

import "fmt"

type produto struct { //criando um tipo e seus atributos
	nome     string
	preco    float64
	desconto float64
}

//Método: função com receiver (informando o produto)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var prod1 produto
	prod1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(prod1, prod1.precoComDesconto())

	prod2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(prod2, prod2.precoComDesconto())
}
