package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalPonteiro(num *int) {
	*num = *num * -1
}

func main() {
	fmt.Println("Ponteiros")
	numero := 42
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)
	novoNum := 422
	fmt.Println(novoNum)
	inverterSinalPonteiro(&novoNum) // passa o endereco de memória da variável
	fmt.Println(novoNum)

}
