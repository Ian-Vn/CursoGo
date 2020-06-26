package main

import "fmt"

func main() {

	// declarando un canal
	c := make(chan int, 1)

	// declarando un funcion anonima
	go func() {
		c <- 42
	}()

	fmt.Println("Valor de canal", <-c)

}
