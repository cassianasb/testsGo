package main

import (
	"fmt"
	"strings"
)

var valPtr *float32
var countPtr *int
var person *struct {
	name string
	age  int
}
var matrix *[1024]int
var row []*int64

func main() {
	fmt.Println(valPtr, countPtr, person, matrix, row)

	var a = 1024
	var aptr = &a

	fmt.Printf("a=%v\n", a)
	fmt.Printf("aptr=%v\n", aptr)

	structPtr := &struct{ x, y int }{44, 55}
	pairPtr := &[2]string{"A", "B"}

	fmt.Printf("struct = %#v, type = %T\n", structPtr, structPtr)
	fmt.Printf("pairPtr = %#v, type = %T\n", pairPtr, pairPtr)

	intptr := new(int)
	*intptr = 44

	p := new(struct{ first, last string })
	p.first = "Samuel"
	p.last = "Pierre"

	fmt.Printf("Value %d, type %T\n", *intptr, intptr)
	fmt.Printf("Person %+v\n", p)

	a1 := 3
	double(&a1)
	fmt.Println(a1)

	p1 := &struct{ first, last string }{"Max", "Planck"}
	cap(p1)
	fmt.Println(p1)
}

func double(x *int) {
	*x = *x * 2
}

func cap(p *struct{ first, last string }) {
	p.first = strings.ToUpper(p.first)
	p.last = strings.ToUpper(p.last)
}
