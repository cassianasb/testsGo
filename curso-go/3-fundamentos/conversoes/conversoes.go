package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4                    //float64
	y := 2                      //inteiro
	fmt.Println(x / float64(y)) //converter o inteiro para float64

	nota := 6.9
	notaFinal := int(nota) //converter o float64 para inteiro
	fmt.Println(notaFinal)

	//concatenar String
	fmt.Println("Teste " + string(123)) //123 é codigo na tabela ASCII por isso será uma chave

	//int para string
	fmt.Println("Teste " + strconv.Itoa(234))

	//string para int
	num, _ := strconv.Atoi("123") //retorna 2 valores, primeiro um nro, segundo é um erro; utilizar o _ para ignorar a variavel
	fmt.Println(num - 122)

	//string para boolean
	b, _ := strconv.ParseBool("true")
	//b, _ := strconv.ParseBool("1") apenas o valor 1 é verdadeiro
	if b {
		fmt.Println("Verdadeiro")
	}
}
