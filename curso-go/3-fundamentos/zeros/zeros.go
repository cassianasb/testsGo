package main

import (
	"fmt"
)

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int //ponteriro de inteiro, valor nulo (nil) por padrão

	fmt.Printf("%v, %v, %v, %q, %v", a, b, c, d, e)
}
