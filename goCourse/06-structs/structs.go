package main

import "fmt"

// structs são uma coleção de campos -- lembra do C? saudades
type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	endereco := endereco{"rua n", 1000}
	var usr1 usuario
	fmt.Println(usr1)
	usr1.nome = "Nome do usuário"
	usr1.idade = 42
	fmt.Println(usr1)

	usr2 := usuario{"nome do segundo usuario", 52, endereco}
	fmt.Println(usr2)

	usr3 := usuario{idade: 21}
	fmt.Println(usr3)
}
