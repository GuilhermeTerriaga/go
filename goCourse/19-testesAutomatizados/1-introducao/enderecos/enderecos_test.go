package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Praca ABC", "Tipo inválido"},
		{"AVENIDA ABC", "Avenida"},
		{"	", "Tipo inválido"},
		{"RUA ABC", "Rua"},
	}
	// enderecoParaTeste := "Avenida Paulista"
	for _, cenario := range cenariosDeTeste {
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if cenario.enderecoEsperado != tipoEnderecoRecebido {
			t.Errorf(
				"O tipo recebido foi %s e é diferente do tipo esperado, %s",
				tipoEnderecoRecebido,
				cenario.enderecoEsperado,
			)
		}

	}
}
