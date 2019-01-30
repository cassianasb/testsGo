package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por *
func inc2(n *int) {
	//* é usado para desreferenciar, ou seja, ter acesso ao valor que o ponteiro aponta
	*n++ //valor no ponteiro n e o incremento no mesmo
}

func main() {
	n := 1

	inc1(n) //por valor, assim não altera o n do escopo da função main
	fmt.Println(n)

	//& usado para obter o endereço
	inc2(&n) //por referencia, altera o n do escopo da funçao main
	fmt.Println(n)
}
