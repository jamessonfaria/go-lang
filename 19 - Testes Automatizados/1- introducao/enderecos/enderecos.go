package enderecos

import (
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se o endereço tem um tipo valido e o retorna
func TipoDeEndereco(end string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavra := strings.Split(end, " ")[0]
	
	primeiraPalavaMinuscula := strings.ToLower(primeiraPalavra)
	if slices.Contains(tiposValidos, primeiraPalavaMinuscula) {
		return caser.String(primeiraPalavaMinuscula)
	}

	return "Tipo invalido"
}