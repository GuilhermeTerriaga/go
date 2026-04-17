package main

import "fmt"

func main() {
	soma := 1 + 1
	subtracao := 1 - 2
	divisao := 10 / 2
	mod := 10 % 2
	fmt.Println(soma, subtracao, divisao, mod)

	var num1 int16 = 1
	var num2 int8 = 2
	// sum := num1 + num2 não funciona pois apesar de serem numeros, eles não são do mesmo tipo
	sum := num1 + int16(num2)
	fmt.Println(sum)

	// atribuições
	var varia1 string = "string"
	varia2 := "string2"
	fmt.Println(varia1, varia2)

	// relacionais
	fmt.Println(1 > 2, 1 < 2, 1 == 2, 1 <= 2, 1 >= 2, 1 != 2)

	// operadores lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// unários
	num := 10
	num++
	fmt.Println(num)
	num += 10
	fmt.Println(num)
	num -= 10
	num *= 5
	num /= 2
	num %= 2
	fmt.Println(num)

	// ternário
	// não tem ternário no Go
	var txt string
	if num > 5 {
		txt = "Maior que 5"
	} else {
		txt = "Menor que 5"
	}

	fmt.Println(txt)
}
