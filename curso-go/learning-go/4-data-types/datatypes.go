package main

import "fmt"

func main() {
	vals := []int{
		1024,
		0x0FF1CE,
		0x8BADF00D,
		0xBEEF,
		077,
	}

	for _, i := range vals {
		if i == 0xBEEF {
			fmt.Printf("Got %d\n", i)
			break
		}
	}

	p := 3.1415926535
	e := .5772156649
	x := 7.2E-5
	y := 1.616199e-35
	z := .416833e32

	fmt.Println(p, e, x, y, z)

	a := -3.5 + 2i
	fmt.Printf("%v\n", a)
	fmt.Printf("%g, %+g\n", real(a), imag(a))

	bksp := '\b'
	tab := '\t'
	nwln := '\n'
	char1 := '\u0369'
	char2 := '\xFA'
	char3 := '\045'

	fmt.Println(bksp)
	fmt.Println(tab)
	fmt.Println(nwln)
	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(char3)
}
