package main

import "fmt"

type hotdog int

var b hotdog = 42

func main() {
	x := 10

	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", x, x)
	x = int(b)
	fmt.Printf("%v, %T \n", x, x)

}
