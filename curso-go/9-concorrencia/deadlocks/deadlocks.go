package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação de bloqueio
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) //canal sem buffer
	go rotina(c)

	fmt.Println(<-c) //operação de bloqueio
	fmt.Println("Foi lido")
	//fmt.Println(<-c)//não existe canal para ser recebido ---> Deadlock
	//fmt.Println("Fim!")
}
