package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.1415

	//fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x) //Gera uma variavel do tipo String do parmetro
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f. ", x)
	fmt.Printf("O valor de x é %f. ", x)

	fmt.Printf("O valor de x é %s", xs)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	//%v imprime quase todos os tipos de forma correta
	fmt.Printf("\n %d \n%f \n%.1f \n%t \n%s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
