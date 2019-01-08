package main

import (
	"fmt"
)

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Aprovado com nota", grade)
	} else {
		fmt.Println("Reprovado com nota", grade)
	}
}

func main() {
	printResult(7.4)
	printResult(9)
	printResult(5.1)
}
