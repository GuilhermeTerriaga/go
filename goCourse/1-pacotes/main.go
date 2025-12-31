package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("guilherme.terriagagmail.com")
	fmt.Println(err)
}
