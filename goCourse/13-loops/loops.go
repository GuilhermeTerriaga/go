package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i:=0
	for i < 10 {
		i++
		fmt.Println("Incrementando i ", i)
		time.Sleep(time.Second)
	}
	for j := 0; j < 10; j++{
		fmt.Println(j)
		time.Sleep(time.Second)
	}
	for k := range 10{
		fmt.Println("valor de k ", k)
	}
	array := [3]string{"Nome1", "Nome2", "Nome3"}
	for _, nome := range array{
		fmt.Println(nome)
	}
}
