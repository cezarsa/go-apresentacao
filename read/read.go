package main

import "io/ioutil"

func main() {
	file, _ := ioutil.ReadFile("arquivo.txt")
	println(string(file))
}
