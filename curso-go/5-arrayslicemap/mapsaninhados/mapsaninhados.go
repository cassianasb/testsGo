package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15090.09,
			"Gustavo Perez":  12786.00,
		},
		"J": {
			"Jayme Mendonça": 17235.76,
		},
		"P": {
			"Paula Campos": 11870.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		//fmt.Println(letra, funcs) //Aula
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}

	}

}
