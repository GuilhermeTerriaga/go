package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosM(n1, n2 int) (int, int) {
	return (n1 + n2), (n1 * n2)
}

func main() {
	soma := somar(40, 2)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Eu sou uma funcao f")
	}
	f()

	var f2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	result := f2("ola")
	fmt.Printf("%v \n", result)
	fmt.Printf("funcao f, %T \n", f)
	fmt.Printf("funcao f2, %T \n", f2)

	resultado1, resultado2 := calculosM(40, 2)
	fmt.Println(resultado1, resultado2)
	resultado, _ := calculosM(40, 2)
	fmt.Println(resultado)

}
