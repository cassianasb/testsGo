package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func notaParaConceitoSwitch(nota float64) string {
	var n = int(nota)
	switch {
	case n < 10 && n > 8:
		return "A"
	case n <= 8 && n > 6:
		return "B"
	case n <= 6 && n > 4:
		return "C"
	case n <= 4 && n > 2:
		return "D"
	case n <= 2 && n >= 0:
		return "E"
	default:
		return "Nota Inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceitoSwitch(9.8))
	fmt.Println(notaParaConceitoSwitch(6.9))
	fmt.Println(notaParaConceitoSwitch(2.1))
}
