package main

import (
	"challenge/cmd/pkg/obtener_cadena_mas_larga"
	"fmt"
)

func main() {
	var cadena = "" //25
	resultado := obtener_cadena_mas_larga.ObtenerSubCadenaMasLarga(cadena)
	fmt.Println(resultado)
}


//Como se declara que un metodo o funcion es exportable o que se lo puede invocar desde otra clase? El metodo tiene que empezar con mayuscula
//Como se importa un paquete y como se invoca al metodo? Se declara import(<<ruta del paquete>>)
//Que pasa si importamos algo y no lo usamos? Go no compila.
//Como se soluciona? 2 Formas, agregando un _ (guion bajo) adelante del import ---> import ( _ "<<ruta del paquete>>"), eliminando el import
//Que pasa si una variable queda sin usar? Go no compila.
//Como se arma un test, regla que tiene que seguir el nombre de la clase y los metodos.
