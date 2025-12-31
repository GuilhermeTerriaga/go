package main

import "fmt"

var x bool // -> Verdadeiro ou Falso -> 0

func main() {
	fmt.Println(x) // -> Falso
	x = true       // -> Verdadeiro
	fmt.Println(x) // -> Verdadeiro
	x = 10 < 100   // -> Verdadeiro
	fmt.Println(x) // -> Verdadeiro

}
