package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		escrever("olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("olá mundo 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
