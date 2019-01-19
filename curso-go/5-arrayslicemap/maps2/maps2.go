package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"José João":     11234.00,
		"Ana Silva":     9876.89,
		"Luíza Pereira": 14564.10,
	}

	funcsESalarios["Paulo Borges"] = 8234.77
	delete(funcsESalarios, "Luciana Souza")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
