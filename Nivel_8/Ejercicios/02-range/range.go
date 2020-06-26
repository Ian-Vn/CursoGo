package main

import "fmt"

func main() {
	// asignando un canal de una función a una variable
	canal := make(chan int)
	gen(canal)

	// Ejecutando la funcion que imprime los valores del canal
	recibir(canal)

	fmt.Println(" A punto de finalizar ")
}

// función que devuelve un canal del tipo int
func gen(c chan int) {

	// Escribimos en el canal, cuando hacemos esta gorutina
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
}

// Funcion que obtiene valores del canal
func recibir(c chan int) {
	// iteremoa sobre el canal
	for v := range c {
		fmt.Println("Valor en el canal", v)
	}
}
