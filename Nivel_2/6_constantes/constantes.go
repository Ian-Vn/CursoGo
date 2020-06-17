package main

import "fmt"

// Creando constantes sin tipo
const (
	a = 1
	b = 2.5
	c = "Hello"
)

// Creando constantes con tipo
const (
	d int     = 3
	e float64 = 20.5
	f string  = "World"
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
