package obtener_cadena_mas_larga

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

var (
	cadena = ""
	resultado = 0
)

func TestDadoUnaCadenaDeTextoObtenerSubCadenaMasLarga(t *testing.T) {
	dadoUnaCadena("abcdefgabc")
	cuandoBuscoLaSubCadenaMasLarga()
	verificoQueLaSubCadenaVale(t,7)
}

func TestDadoUnaCadenaDeNumerosObtenerSubCadenaMasLarga(t *testing.T) {
	dadoUnaCadena("1234123423567899")
	cuandoBuscoLaSubCadenaMasLarga()
	verificoQueLaSubCadenaVale(t,7)
}

func TestDadoUnaCadenaDeTextoYNumeroObtenerSubCadenaMasLarga(t *testing.T) {
	dadoUnaCadena("abcd1234ab")
	cuandoBuscoLaSubCadenaMasLarga()
	verificoQueLaSubCadenaVale(t,8)
}

func TestDadoUnaCadenaVaciaObtenerSubCadenaMasLarga(t *testing.T) {
	dadoUnaCadena("")
	cuandoBuscoLaSubCadenaMasLarga()
	verificoQueLaSubCadenaVale(t,0)
}

func dadoUnaCadena(s string) {
	cadena = s
}

func cuandoBuscoLaSubCadenaMasLarga() {
	resultado = ObtenerSubCadenaMasLarga(cadena)
}

func verificoQueLaSubCadenaVale(t *testing.T, valorEsperado int) {
	assert.Equal(t, valorEsperado, resultado)
}

