package main

import "fmt"

//func* + nome da função* (parametro1 tipo do parametro1, parametro2 tipo do parametro2) (tipo de retorno1, tipo de retorno2)
//*Obrigatórios
func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s \n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("home", "run")
	fmt.Println(f3())
	fmt.Println(f4("spider", "man"))
	fmt.Println(f5())

	r3, r4 := f3(), f4("Param 1", "Param 2")
	fmt.Println(r3)
	fmt.Println(r4)

	//_, r52 := f5()
	//r51, _ := f5()
	r51, r52 := f5()
	fmt.Println("F5:", r51)
	fmt.Println("F5:", r52)
}
