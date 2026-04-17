package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"joao", "sobrenome", 20, 178}
	fmt.Println(p1)
	e1 := estudante{pessoa{"zé", "sobrenome do zé", 20, 150}, "curso1", "faculdade 1"}
	e2 := estudante{p1, "curso1", "faculdade 1"}
	fmt.Println(e1, e2)
}
