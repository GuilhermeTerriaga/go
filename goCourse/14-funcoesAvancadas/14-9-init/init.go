package main

import "fmt"

var n int
func init(){ // pode ter uma por arquivo! útil para fazer inicialações e setup de variável
	fmt.Println("Rodando a init")
	n = 10
}

func main() {
	fmt.Println("Rodando a main")
	fmt.Println(n)
}
