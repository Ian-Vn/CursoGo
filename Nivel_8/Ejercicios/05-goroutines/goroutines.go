package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	// Como se espera que el canal emita 100 valores se inicia el ciclo for hasta ese numero
	// si fuera mas de ese numero entonces se bloquearia la gorutina principal porque estaria
	// esperando un valor del canal y como no hay otras  gorutinas ejecutandose entonces
	// cae en un deadlock
	for i := 1; i <= x*y; i++ {
		fmt.Println("Valor del canal: ", <-c)
	}

	fmt.Println("GoRUTINAS", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	// Cada gorutina imprimirá el valor correspondiente entonces esperara a que otra goroutine
	// lea el valor del canal por el cual el hilo principal empezará a tomar dichos valores
	// consecutivamente, por cada emision de valores los hilos permaneceran esperando y el hilo principal se encargara
	// de tomar dichos valores
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("GoRUTINAS", runtime.NumGoroutine())
	}
	return c
}
