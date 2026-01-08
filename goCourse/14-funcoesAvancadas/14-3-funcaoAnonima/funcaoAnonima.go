package main

import "fmt"

func main(){
	result := func(texto string) string{
		return fmt.Sprintf("Recebido -> %v", texto)
	}("passa parametro aqui")
	fmt.Println(result)
	// será que é possível ter uma funcao anonima com retorno nomeado e variádica?
	result2 := func(nums ...int) (soma int){
		for _, num := range nums{
			soma += num
		}
		return
	}(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(result2)
}
