package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	canal := multiplexerVariadico(
		escrever("olá mundo!"),
		escrever("go golang! go! go!"),
		escrever("variadic!"),
	)
	fmt.Println(canal)
	for msg := range canal {
		fmt.Println(msg)
	}
}

func multiplexerVariadico(chans ...<-chan string) <-chan string {
	canalDeSaida := make(chan string)
	var wg sync.WaitGroup
	for _, chanell := range chans {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for mensagem := range c {
				canalDeSaida <- mensagem
			}
		}(chanell)
	}

	go func() {
		wg.Wait()
		close(canalDeSaida)
	}()
	return canalDeSaida
}

func multiplexer(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem

			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}

		}
	}()
	fmt.Println(canalDeSaida)
	return canalDeSaida
}
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(2000)))
		}
	}()

	return canal
}
