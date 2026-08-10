package main

import "fmt"

func main() {
	canal := make(
		chan string,
		2,
	) // aqui tem um canal com buffer, é um limite para as mensagens num canal
	canal <- "Olá mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
