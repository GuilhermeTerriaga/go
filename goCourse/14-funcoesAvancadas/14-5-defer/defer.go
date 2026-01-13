package main

import "fmt"

func funcao1(){
	fmt.Println("executando a func 1")
}

func funcao2(){
	fmt.Println("executando a func 2")
}

func calcularMedia(notas ...int) bool{
	defer fmt.Println("média calculada, resultado será retornado") // muito útil para evitar repetição de código
	soma := 0
	for _, i := range(notas){
		soma += i
	}
	if soma/len(notas) >= 6{
		fmt.Println("Aluno aprovado!")
		return true
	}
	fmt.Println("Aluno não aprovado")
	return false
}

func main(){
	defer funcao1() // DEFER = adiar, ele vai adiar a execução daquela linha até a linha anterior a do retorno.
	funcao2()
	// fmt.Println(calcularMedia(0,2,3,4,5,4,3))
	fmt.Println(calcularMedia(7,7,8,9,10,10))
}
