package main

import (
	"fmt"
	"strings"
)

func main() {

	// --------------- Ejercicio 6 ----------
	// Ejecucion de una funcion anonima
	func(nombre string) {
		fmt.Println("Hola", nombre, "como estas")
	}("Ian")

	// -------------- Ejercicio 7 ---------------
	// Creacion de una funcion y asignacion de dicha funcion a una variable
	suma := func(x ...int) int {
		suma := 0
		for _, i := range x {
			suma += i
		}

		return suma
	}

	// Pasando como argumentos un slide de ints
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(suma(numeros...))

	// Llamada a la funcion
	// Como OtraFuncion regresa una funcion, se le puede asignar a una variable,
	// esta variable almacena una funcion
	llamada := otraFuncion("ian")

	// si mandamos a llamar a la funcion devuelta entonces regresamos un string
	fmt.Println(llamada())

}

// ------------ Ejercicio 8 -------------------
// Crear una funcion que retorne una funcion
// Esta funcion llamada otraFuncion retorna una funcion que a su vez retorna un string
func otraFuncion(nombre string) func() string {
	return func() string {
		return "Hola " + strings.ToUpper(nombre) + " como estas"
	}
}
