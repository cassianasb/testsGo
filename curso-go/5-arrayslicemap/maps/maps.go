package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[11122233344] = "Maria"
	aprovados[77722233355] = "João"
	aprovados[97988836312] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[97988836312])
	delete(aprovados, 97988836312)
	fmt.Println(aprovados[97988836312])
}
