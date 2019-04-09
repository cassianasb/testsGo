package main

import "fmt"

//Curr...
type Curr struct {
	currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Alegria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344},
	Curr{"USD", "US Dollar", "United States", 840},
	Curr{"BRL", "Real", "Brazil", 986},
	Curr{"CAD", "Canadian Dollar", "Canadian", 124},
}

//Current:...
type Current struct {
	Name    string
	Country string
	Number  int
}

//CAD:...
var CAD = Current{
	Name:    "Canadian Dollar",
	Country: "Canada",
	Number:  124,
}

//FJD:...
var FJD = Current{
	Name:    "Fiji Dollar",
	Country: "Fiji",
	Number:  242,
}

//JMD:...
var JMD = Current{
	Name:    "Jamaican Dollar",
	Country: "Jamaica",
	Number:  388,
}

//USD:...
var USD = Current{
	Name:    "US Dollar",
	Country: "USA",
	Number:  840,
}

func isDollar(curr Curr) bool {
	var result bool
	switch curr {
	default:
		result = false
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		result = true
	case Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344}:
		result = true
	case Curr{"USD", "US Dollar", "United States", 840}:
		result = true
	}
	return result
}

func isDollar2(curr Curr) bool {
	dollars := []Curr{currencies[2], currencies[4], currencies[9]}
	switch curr {
	default:
		return false
	case dollars[0]:
		fallthrough
	case dollars[1]:
		fallthrough
	case dollars[2]:
		return true
	}
}

func isEuro(curr Curr) bool {
	switch curr {
	case currencies[2], currencies[4]:
		return true
	default:
		return false
	}
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

	curr := Curr{"EUR", "Euro", "Italy", 978}
	if isDollar(curr) {
		fmt.Printf("%+v  is Dollar currency \n", curr)
	} else if isEuro(curr) {
		fmt.Printf("%+v  is Euro currency \n", curr)
	} else {
		fmt.Println("Currency is not Dollar or Euro")
	}

	dol := Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344}
	if isDollar2(dol) {
		fmt.Println("Dollar currency found: ", dol)
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
