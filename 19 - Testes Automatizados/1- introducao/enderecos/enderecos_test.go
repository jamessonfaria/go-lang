// TESTE UNITARIO

package enderecos_test

import (
	. "introducao-testes/enderecos" // esse ponto na frente faz um alias no documento para nao ter que mencionar enderecos antes das funcoes
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T)  {

	t.Parallel() // roda os testes em paralelo

	cenariosDeTeste := []cenarioDeTeste {
		{ "Rua ABC", "Rua" },
		// { "Avenida Paulista", "Avenida" },
		// { "Rodovia dos Imigrantes", "Rodovia" },
		// { "Estrada do Arraial", "Estrada" },
		// { "Praça as Rosas", "Tipo invalido" },
		// { "RUA DOS BOBOS", "Rua" },
		// { "AV das Alegras", "Tipo invalido" }, 
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s!", retornoRecebido, cenario.retornoEsperado)
		}
	}
}