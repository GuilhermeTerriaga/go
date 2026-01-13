package main

import (
	"fmt"
)

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

func main() {
	fmt.Println("recursivas")
	posicao := uint(15)
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
