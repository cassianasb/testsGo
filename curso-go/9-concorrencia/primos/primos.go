package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) //fechar o canal
}

func main() {
	ch := make(chan int, 30)
	go primos(cap(ch), ch) //cap=capacidade do buffer do canal

	for primo := range ch {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
