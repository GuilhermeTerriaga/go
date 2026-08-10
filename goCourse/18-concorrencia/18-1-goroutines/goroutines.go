package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("olá mundo!") // aqui é uma go routine
	escrever("olá outro mundo!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
