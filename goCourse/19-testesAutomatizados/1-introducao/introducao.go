package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("rua gaspar conqueiro")
	fmt.Print(tipoEndereco)
}
