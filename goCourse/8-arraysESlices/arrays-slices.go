package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [10]int
	fmt.Println(array1) // arrays têm a posição inicial em 0
	array1[0] = 42
	fmt.Printf("array1: %v\n", array1)
	arr2 := [5]uint8{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2)

	arr3 := [...]uint8{6, 7, 8, 9, 10, 11}
	fmt.Printf("arr3: %v\n", arr3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Printf("slice: %v\n", slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))
	slice = append(slice, 18)
	fmt.Println(reflect.TypeOf(slice))

	slice2 := arr2[1:3]
	fmt.Printf("slice2: %v\n", slice2)

	arr2[1] = 42
	fmt.Printf("slice2: %v\n", slice2) // ele é um ponteiro!
}
