package main

import "fmt"

var x uint16

func main() {
	x = 65535
	fmt.Printf("%v\n", x)
	// x = 65536 ->> cannot use 65536 (untyped int constant) as uint16 value in assignment (overflows)
	// fmt.Printf("%v\n", x)
	x++ // 0
	fmt.Printf("%v\n", x)
	x++ // 1
	fmt.Printf("%v\n", x)
	x = 65536

}
