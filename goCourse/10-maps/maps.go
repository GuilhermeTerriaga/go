package main

import "fmt"

func main() {
	fmt.Println("Maps!")
	usuarios := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	usuarios2 := map[string]map[string]string{
		"nomeCompleto": {"nome": "zé", "sobrenome": "silva"},
		"curso": {"nome": "Engenharia", "campus": "campus 1"},
	}
	fmt.Println(usuarios)
	fmt.Println(usuarios["nome"])
	fmt.Println(usuarios2)
	fmt.Println(usuarios2["nomeCompleto"]["nome"])
	for chave := range usuarios2 {
		fmt.Println(chave)
	}
}
