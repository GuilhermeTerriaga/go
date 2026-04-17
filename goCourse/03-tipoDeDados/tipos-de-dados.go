package main

import (
	"errors"
	"fmt"
)

func main() {
	// inteiros
	// int8, int16, int32, int64
	var numero int = 1000000000000
	fmt.Println(numero)

	//uint - unsigned integer
	var num2 uint = 10

	fmt.Println(num2)
	// rune e int32 são alias
	var num3 rune = 42
	fmt.Printf("rune %v, %T \n", num3, num3)

	// BYTE e uint8 são alias
	var num4 byte = 42
	fmt.Printf("byte %v, %T \n", num4, num4)

	// float32 e float64 para ponto flutuante
	var num5 float32 = 1230.56
	fmt.Printf("float %v, %T \n", num5, num5)

	num6 := 12356.78
	fmt.Printf("float %v, %T \n", num6, num6)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'A'
	fmt.Printf("char %v, %T \n", char, char)

	// NIL
	var str3 string

	fmt.Printf("string vazia, %v, %T \n", str3, str3)

	// bool default value é false
	var booleano bool
	fmt.Println(booleano)

	var erro error

	fmt.Printf("erro %v, %T \n", erro, erro)
	
	var err error = errors.New("Erro interno")
	fmt.Printf("erro %v, %T \n", err, err)


}
