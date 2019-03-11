package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//<- chan - Somente leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls { //índice ignorado
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)</title>")
			rs := r.FindStringSubmatch(string(html))
			if len(rs) != 0 {
				c <- rs[1]
			}
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://twitter.com/", "https://www.google.com/")
	t2 := titulo("https://www.google.com", "https://www.youtube.com")

	fmt.Println("Primeiros: ", <-t1, " | ", <-t2)
	fmt.Println("Segundos: ", <-t1, " | ", <-t2)
}
