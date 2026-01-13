package main

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil{ // importante para evitar que o programa seja recuperado sem necessidade
		fmt.Println("execucao recuperada")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media:=(n1+n2)/2
	if media > 6 {
		fmt.Println("Aluno aprovado")
		return true
	} else if media < 6{
		fmt.Println("Aluno reprovado")
		return false
	}

	panic("Média exatamente 6!") // panic vai chamar todos os defers antes de matar de fato a execução
}


func main() {
	fmt.Println(alunoAprovado(6,6))	
}
