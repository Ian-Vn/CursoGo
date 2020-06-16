package main

import "fmt"

// Creación de un tipo de datos
type newType int

// Variable global
var x newType

func main() {

	// Imprimiendo el valor de X
	fmt.Println("Valor de x: ", x)
	// Imprimiendo el tipo de la valor de X
	fmt.Printf("Tipo de dato de X: %T \n", x)
	// Reasignando el valor de la variable X
	x = 42
	// Imprimiendo el nuevo valor
	fmt.Println("Valor de x: ", x)

}
