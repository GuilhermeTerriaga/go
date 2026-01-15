package main

import "fmt"

func closure() func(){ // recebe nada como parametro, retorna uma funcão o escopo da variavel texto fica dentro da funcao
	texto := "dentro da closure"
	
	funcao := func()  {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("Closures")
	texto := "dentro da funcao main"
	fmt.Println(texto)
	funcaoNova := closure()
	funcaoNova() // sempre que chamarmos essa referencia da closure, ela vai utilizar a variavel texto do escopo dela
}
