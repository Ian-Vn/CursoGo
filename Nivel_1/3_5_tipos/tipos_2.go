package main

import "fmt"

// Creación de un tipo de datos
type newType int

// Variable global con el tipo de dato creado anteriormente
var x newType

// Otra variable global
var y int

func main() {

	// Imprimiendo el valor de X
	fmt.Println("Valor de x: ", x)
	// Imprimiendo el tipo de la valor de X
	fmt.Printf("Tipo de dato de X: %T \n", x)
	// Reasignando el valor de la variable X
	x = 42
	// Imprimiendo el nuevo valor
	fmt.Println("Valor de x: ", x)
	// Conversion de tipo al valor original
	// La siguiente sentencia no es valida porque no son del mismo tipo aunque internamente lo sean
	// y = x
	y = int(x)
	fmt.Println("Valor de Y: ", y)

}
