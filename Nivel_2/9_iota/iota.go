package main

import "fmt"

// Iota permite incrementar una variable en uno
const (
	// Para que no use el incremento en 0
	_ = iota
	// Para hacer un incremento en 1, las siguientes variables se incrementan en 1, si ponemos
	// iota a cada variable se reinicia en 1 y se incrementan, por defecto iota se empieza en 0
	a = iota + 2020
	// Si hacemos lo siguiente la variable b toma el valor de 2 porque se reinicio
	// b = iota
	// Si no declaramos iota en las siguientes variables se toma como referencia la variable anterior y se
	// incrementa
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
