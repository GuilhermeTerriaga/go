package main

import "fmt"

func main() {
	var1 := 10
	var2 := var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)
	var pont *int
	pont = &var1

	fmt.Printf("pont: %v\n", pont) // aqui mostra apenas o endereco de memória do ponteiro
	var1++
	fmt.Printf("pont: %v\n", *pont) // aqui usando o * a gente manda o ponteiro buscar o valor do endereço
	pont2 := &pont
	fmt.Printf("pont2: %v\n", **pont2) // ponteiro de ponteiro? hahahaha
}
