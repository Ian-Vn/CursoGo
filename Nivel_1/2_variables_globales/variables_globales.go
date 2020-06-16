package main

import "fmt"

// Creacion de variables globales, el operador de asignacion corta no funciona
var x int
var y string
var z bool

func main() {
	// Impresión de los valores, se inicializan a 0, "" y a false
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
