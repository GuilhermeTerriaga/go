package enderecos

import (
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoEndereco verifica se um endereco é do tipo valido
func TipoEndereco(endereco string) string {
	caser := cases.Title(language.Portuguese)
	tiposValidos := []string{"rua", "avenida", "estrada", "viela", "travessa"}
	endereco = strings.ToLower(endereco)
	primeiraPalavraDoEndereco, _, _ := strings.Cut(endereco, " ")
	if slices.Contains(tiposValidos, primeiraPalavraDoEndereco) {
		return caser.String(primeiraPalavraDoEndereco)
	}
	return "Tipo inválido"
}
