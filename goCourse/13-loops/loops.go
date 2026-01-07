package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i ", i)
		// time.Sleep(time.Second)
	}
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		// time.Sleep(time.Second)
	}
	for k := range 10 {
		fmt.Println("valor de k ", k)
	}
	fmt.Println("array")
	array := [3]string{"Nome1", "Nome2", "Nome3"}
	for _, nome := range array {
		fmt.Println(nome)
	}
	fmt.Println("slice")
	slice := []string{"Nome1", "Nome2", "Nome3"}
	for _, nome := range slice {
		fmt.Println(nome)
	}
	fmt.Println("por sobre um string")
	for _, letra := range "PALAVRA" {
		fmt.Println(letra) // ->> aqui vai trazer a rune da letra
		fmt.Println(string(letra))
	}
	users := map[string]string{
		"nome":      "zé",
		"sobrenome": "sobrenome lol",
	}
	for chave, valor := range users {
		println(chave, valor)
	}
	type userStruct struct {
		nome      string
		sobrenome string
	}
	// não é possível iterar sobre um struct
	// usuario := userStruct{"zé", "sobrenome do zé"}
	// for chave, valor := range usuario{
	// 	fmt.Println(chave, valor)
	// }
	for {
		fmt.Println("Infinito")
		time.Sleep(time.Second)
	}
}
