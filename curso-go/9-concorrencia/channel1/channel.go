package main

import "fmt"

func main() {
	ch := make(chan int, 1) //canal com valores inteiros com 1 buffer
	//Canal é um ponto de sincronismo.
	//Enviar informação pelo canal
	ch <- 1 //Enviando para o canal o valor 1

	//Receber informação pelo canal
	<-ch //Recebendo o valor do canal (não é obrigatório uma variavel para recebimento)

	ch <- 2
	fmt.Println(<-ch)
}
