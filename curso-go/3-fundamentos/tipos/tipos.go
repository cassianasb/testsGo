package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 100)
	//TypeOf apresenta o tipo da variável passada como parametro
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//Sem sinal (positivos)... unit8, unit16, unit32 e unit64 (unsignal int: byte, short, int, long)
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Com sinal... int8, int16, int32 e int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo da variavel i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representação na tabela unicode (int 32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("A represnetação na tabela unicode de i2 é", i2)

	//numeros reais float32 e float64
	var x float32 = 49.99 // ou var x = float32(49.99)
	fmt.Println("O tipo x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) //por padrão é float64

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Olá, meu nome é Cassiana"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de s1 é", len(s1))

	//String com muitas linhas
	s2 := `Olá,
	 meu
	 nome
	 é
	 Cassiana`
	fmt.Println(s2)

	//Char não existe em Go, é um int32, apesar de permitir um caracter entre aspas simples
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}
