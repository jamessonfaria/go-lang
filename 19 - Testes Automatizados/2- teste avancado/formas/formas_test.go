package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T){
		ret := Retangulo{ 10, 12 }
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada) // falta para aqui, o error continua os testes
		}

	})

	t.Run("Circulo", func(t *testing.T){
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada) // falta para aqui, o error continua os testes
		}
	})
	
}