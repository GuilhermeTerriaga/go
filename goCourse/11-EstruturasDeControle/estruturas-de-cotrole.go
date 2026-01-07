package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")
	num := 10
	if num > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	if num2 := num + 5; num2 > 0 { // if init -> variável com escopo fechado dentro do condicional
		fmt.Println("Número é maior que zero")
	} else if num2 < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
