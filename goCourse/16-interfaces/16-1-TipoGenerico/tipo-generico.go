package main

import "fmt"

func generica(interf interface{}){
	fmt.Println(interf)
}

func main() {
	generica("teste")
	generica(1)
	generica(true)

	fmt.Println(true, false, 1, 2, 42.5, "teste") // Println é uma func que recebe interfaces

	mapa := map[interface{}]interface{}{ // cuidado ao utilizar isso, pode levar a uma grande bagunça
		1: "String",
		float32(123): 2,
		"String": 3,
		true: float64(523),
	}

	fmt.Println(mapa)
}
