package main

import "fmt"
import "math/rand"
import "time"

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	//iniciar o i dentro do bloco if-else
	if i := randomNumber(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
