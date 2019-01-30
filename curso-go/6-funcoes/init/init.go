package main

import "fmt"

//funão que será executada sempre no começo da execução independente da função ter sido chamada
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main")
}
