package main

import "fmt"

func main() {
	var variavel1 string = "hello variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "asd"
		variavel4 string = "dsa"
	)
	fmt.Println(variavel4, variavel3)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante1"
	fmt.Println(constante1)
	
	//inverter os valores das variáveis
	variavel6, variavel5 = variavel5, variavel6

	fmt.Println(variavel5, variavel6)
}
