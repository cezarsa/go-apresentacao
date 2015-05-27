package main

import "fmt"

func contarLetra(str string, letra rune, ch chan int) {
	counter := 0
	for _, c := range str {
		if c == letra {
			counter++
		}
	}
	ch <- counter
}

func main() {
	str := "zzzzAzzzzzzAzzzzzAzzAzzzzAz"
	ch := make(chan int)
	go contarLetra(str, 'A', ch)
	fmt.Printf("%d\n", <-ch)
}
