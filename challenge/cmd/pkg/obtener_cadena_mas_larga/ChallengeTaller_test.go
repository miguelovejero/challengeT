package obtener_cadena_mas_larga

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

var (
	cadena = ""
	resultado = 0
)

func TestObtenerCadenaMasLarga(t *testing.T) {
	dadoUnaCadenaDeTexto("abcdefg")
	cuandoBuscoLaSubCadenaMasLarga()
	verificoQueLaSubCadenaVale(t,7)
}

func dadoUnaCadenaDeTexto(s string) {
	cadena = s
}

func cuandoBuscoLaSubCadenaMasLarga() {
	resultado = ObtenerSubCadenaMasLarga(cadena)
}

func verificoQueLaSubCadenaVale(t *testing.T, valorEsperado int) {
	assert.Equal(t, valorEsperado, resultado)
}



