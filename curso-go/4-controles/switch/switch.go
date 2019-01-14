package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough //continuar executando, no Java sempre executa os cases; 9 e 10 são A
	case 9:
		return "A"
	case 8, 7: //Dois casos onde o retorno é igual
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: //Caso não seja nenhum dos cases
		return "Nota Inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7.6))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(11))
}
