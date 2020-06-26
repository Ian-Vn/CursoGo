package main

import (
	"fmt"
	"math"
)

// Ejercicio 6
// Crear una estructura de tipo cuadrado
type cuadrado struct {
	lado float64
}

// Crear una estructura llamada circulo
type circulo struct {
	radio float64
}

// Creando metodos para ambos tipos de estructuras
func (c cuadrado) area() float64 {
	return math.Pow(c.lado, 2)
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

// Creando una interfaz que sea implementada por ambas estructuras
type forma interface {
	area() float64
}

func main() {
	// Creando una variable del tipo cuadrado
	figura1 := cuadrado{lado: 20.0}
	figura2 := circulo{radio: 5.0}
	fmt.Println(figura1, figura2)
	info(figura1)
	info(figura2)
}

// Creando una funcion que apartir de la interface calcule el area
func info(figura forma) {
	// Al ser de tipo forma y como esta interfaz es implementada por cuadrado y circulo
	// tiene acceso al metodo area
	fmt.Println("Area: ", figura.area())
}
