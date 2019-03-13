package main

import (
	"fmt"

	"github.com/cassianasb/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar: misturar (mensagens) em um canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.globo.com/", "https://www.google.com/"),
		html.Titulo("https://twitter.com/", "https://www.google.com/"),
	)
	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
