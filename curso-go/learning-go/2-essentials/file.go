package main

import (
	"fmt"
	"math/rand"
	"time"
)

var greetings = [][]string{
	{"Olá Mundo", "Portuguese"},
	{"Hello World", "English"},
	{"Salut Monde", "Franch"},
	{"qo' vIvan", "Klingon"},
	{"Wapendwa Dunia", "Swahili"},
	{"Hola Mundo", "Spanish"},
	{"Merhaba Dünya", "Turkish"},
}

func greeting() []string {
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	return greetings[rnd.Intn(len(greetings))]
}

func main() {
	g := greeting()
	fmt.Printf("%s (%s)\n", g[0], g[1])
}
