package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	var s1 []int

	//a1=append(a1, 1,4,5,6)
	s1 = append(s1, 1, 4, 5, 6) //append só funciona para slice
	fmt.Println(a1, s1)

	s2 := make([]int, 2)
	copy(s2, s1) //copia a quantidade possivel de elementos (s2 com 2 posições copia os 2 primeiros elementos de s1)
	fmt.Println(s2)

}
