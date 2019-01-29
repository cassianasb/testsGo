package main

import "fmt"

func obterValorAprovado(nro int) int {
	defer fmt.Println("Fim!") //defer atrasa a execução até o final do metodo.
	if nro > 5000 {
		fmt.Println("Valor Alto...")
		return 5000
	}
	fmt.Println("Valor Baixo...")
	return nro

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
