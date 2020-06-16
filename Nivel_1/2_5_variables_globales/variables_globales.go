package main

import "fmt"

// Creacion de variables globales, el operador de asignacion corta no funciona
var x = 42
var y = "James Bond"
var z = true

func main() {
	// Creacion de un nuevo string a partir de las variables globales
	s := fmt.Sprintf("Valor de X: %d, Valor de Y: %s, Valor de Z: %t", x, y, z)
	fmt.Println(s)
}
