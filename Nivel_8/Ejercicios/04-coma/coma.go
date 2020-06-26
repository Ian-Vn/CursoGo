package main

import "fmt"

func main() {

	// creamos un canal
	c := make(chan int)

	// Iniciamos una gorutina
	go func() {
		// Escribimos un valor en la gorutina
		c <- 32

	}()
	// Comprobamos si un canal esta abierto o cerrado, devolvera true en ok si esta abierto
	v, ok := <-c

	//Imprimimos el resultado
	fmt.Println(v, ok)

	// cerramos el canal, si no cerramos el canal se dara un deadlock
	close(c)

	// volvemos a verificar pero como el canal se cerro entonces devolvera false y el valor de v es el valor cero
	v, ok = <-c

	//Imprimimos el resultado
	fmt.Println(v, ok)

}
