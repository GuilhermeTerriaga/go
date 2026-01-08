package main

import "fmt"
// nesse caso é variádica e com retorno nomeado
func soma(nums ...int) (total int) { // funcao variadica é aquela que recebe um parametro com tamanho não especificado
	for _, num := range nums {
		total += num
	}
	return
}
// não pode-se ter mais de um parametro variadico por funcao
func escrever (texto string, nums ...int) {
	fmt.Println(texto)
	for _, num := range nums {
		fmt.Println(num)
	}
}
func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8))
	escrever("texto uhul!", 1,2,3,4,5,6,7,8)
}
