package obtener_cadena_mas_larga

func ObtenerSubCadenaMasLarga(cadena string) int {
	var contador = 0
	var mapa = make(map[string] int)

	for _, l := range cadena {
		_, exist := mapa[string(l)]
		if exist {
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