package main

import "fmt"

func contarLetra(str string, letra rune) int {
	counter := 0
	for _, c := range str {
		if c == letra {
			counter++
		}
	}
	return counter
}

func main() {
	str := "zzzzAzzzzzzAzzzzzAzzAzzzzAz"
	n := contarLetra(str, 'A')
	fmt.Printf("%d\n", n)
}
