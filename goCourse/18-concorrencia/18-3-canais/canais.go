package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for range 5 {
		canal <- texto
		time.Sleep(time.Second)
	}

	// com o close, ele fecha o canal
	close(canal)
}

func main() {
	// usa make chan tipo
	canal := make(chan string)
	// agora uso a go routine
	go escrever("olá mundo!", canal)
	for { // aqui acontecerá um deadlock, pois o canal vai ficar esperando para sempre chegar uma msg
		mensagem, aberto := <- canal // aguarda receber o valor da mensagem. é um await mais ou menos com o aberto evita deadlock
		fmt.Printf("%v\n", mensagem)
		if !aberto {
			break
		}
	}
	canal = make(chan string)
	go escrever("olá 2", canal)
	for mensagem := range canal { // dessa forma não há chance de deadlock
		fmt.Println(mensagem)
	}
}
