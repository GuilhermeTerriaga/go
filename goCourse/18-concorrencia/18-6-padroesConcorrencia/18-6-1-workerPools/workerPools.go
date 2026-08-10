package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	results := make(chan int, 45)
	go worker(tarefas, results)
	go worker(tarefas, results)
	go worker(tarefas, results)
	go worker(tarefas, results)
	go worker(tarefas, results)
	go worker(tarefas, results)
	for i := range 45 {
		tarefas <- i
	}
	close(tarefas)

	for result := range results {
		// result := <-results
		fmt.Println(result)
	}
}

func worker(tarefas <-chan int, results chan<- int) {
	defer close(results)
	for msg := range tarefas {
		results <- fibonacci(msg)
	}
}

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}
