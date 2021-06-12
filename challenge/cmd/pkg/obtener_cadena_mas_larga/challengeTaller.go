package obtener_cadena_mas_larga

func ObtenerCadenaMasLarga(cadena string) int {
	var contador = 0
	var mapa = make(map[string] int)

	for _, l := range cadena {
		_, err := mapa[string(l)]
		if err {
			actualizarContador( &contador, len(mapa) )
			mapa = make(map[string] int)
		}
		mapa[string(l)] = 1
	}
	actualizarContador( &contador, len(mapa) )
	return contador
}

func actualizarContador(contador *int, lenMapa int) {
	if lenMapa > *contador {
		*contador = lenMapa
	}
}