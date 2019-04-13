package main

import (
	"fmt"
	"strings"
)

//Curr...
type Curr struct {
	Currency string
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

	find("Real")
	fmt.Println(assertEuro(currencies[2]))
	fmt.Println(assertEuro(currencies[8]))

	findAny("Peso")
	findAny(978)
	findAny("BRL")
	findAny(false)

	fmt.Println(currencies)
	sortByNumber()
	fmt.Println(currencies)
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

//Switch sem expressão
func find(name string) {
	for i := 0; i < len(currencies); i++ {
		c := currencies[i]
		switch {
		case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name):
			fmt.Println("Found", c)
		}
	}
}

//inicializando a variavel no switch
func assertEuro(cur Curr) bool {
	switch name, curr := "Euro", "EUR"; {
	case cur.Name == name:
		return true
	case cur.Currency == curr:
		return true
	}
	return false
}

func findNumber(num int) {
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("Found", curr)
		}
	}
}

//comparando por tipo da variavel
func findAny(val interface{}) {
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}

func sortByNumber() {
	N := len(currencies)
	for i := 0; i < N-1; i++ {
		currMin := 1
		for k := i + 1; k < N; k++ {
			if currencies[k].Number < currencies[currMin].Number {
				currMin = k
			}
		}
		//swap
		if currMin != i {
			temp := currencies[i]
			currencies[i] = currencies[currMin]
			currencies[currMin] = temp
		}
	}
}
