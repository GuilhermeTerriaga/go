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

	// arrays internos
	slice3 := make([]float32, 10, 11) // cria um array e referencia ele internamente
	fmt.Printf("slice3 len: %v, cap: %v\n%v\n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 10)
	fmt.Printf("slice3 len: %v, cap: %v\n%v\n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 10)
	fmt.Printf("slice3 len: %v, cap: %v\n%v\n", len(slice3), cap(slice3), slice3)
	// conforme crescimento do slice, a lang se ajusta para aumentar o tamanho, o que evita uma quebra mas
	// pode levar a um slice enorme e risco de memory leak!!! usar com cuidado
	slice4 := make([]float32, 5) // nesse momento o slice tem cap 5
	fmt.Printf("slice4 len: %v, cap: %v\n%v\n", len(slice4), cap(slice4), slice4)
	slice4 = append(slice4, 10) // aqui foi para 12 pois é sempre o len * 2
	fmt.Printf("slice4 len: %v, cap: %v\n%v\n", len(slice4), cap(slice4), slice4)



}
