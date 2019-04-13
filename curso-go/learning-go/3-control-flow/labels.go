package main

import "fmt"

var words = [][]string{
	{"break", "lake", "go", "right", "strong", "kite", "hello"},
	{"fix", "river", "stop", "left", "weak", "flight", "bye"},
	{"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},
}

func search(w string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if w == words[i][k] {
				fmt.Println("Found", w)
				break DoSearch
				//continue DoSearch
			}
		}
	}
}

func main() {
	search("high")
	search("fix")

	var a string
Start:
	for {
		switch {
		case a < "aaa":
			goto A
		case a >= "aaa" && a < "aaabbb":
			goto B
		case a == "aaabbb":
			break Start
		}
	A:
		a += "a"
		continue Start
	B:
		a += "b"
		continue Start
	}
	fmt.Println(a)
}
