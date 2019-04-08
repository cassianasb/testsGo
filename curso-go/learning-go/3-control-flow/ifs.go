package main

import "fmt"

type Current struct {
	Name    string
	Country string
	Number  int
}

var CAD = Current{
	Name:    "Canadian Dollar",
	Country: "Canada",
	Number:  124,
}

var FJD = Current{
	Name:    "Fiji Dollar",
	Country: "Fiji",
	Number:  242,
}

var JMD = Current{
	Name:    "Jamaican Dollar",
	Country: "Jamaica",
	Number:  388,
}

var USD = Current{
	Name:    "US Dollar",
	Country: "USA",
	Number:  840,
}

func main() {
	num0 := 242
	if num0 > 100 || num0 < 900 {
		fmt.Println("Currency: ", num0)
		printCurr(num0)
	} else {
		fmt.Println("Current unknown")
	}

	if num1 := 388; num1 > 100 || num1 < 900 {
		fmt.Println("Currency: ", num1)
		printCurr(num1)
	}
}

func printCurr(number int) {
	if CAD.Number == number {
		fmt.Printf("Found: %v\n", CAD)
	} else if FJD.Number == number {
		fmt.Printf("Found %v\n", FJD)
	} else if JMD.Number == number {
		fmt.Printf("Found %v\n", JMD)
	} else if USD.Number == number {
		fmt.Printf("Found: %v\n", USD)
	} else {
		fmt.Println("No currency found with number", number)
	}
}
