package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Creando una variable del tipo waitgroup para la sincronizacion de datos
var wg sync.WaitGroup

func main() {

	// Usamos el metodo Add para decirle al main que se estan ejecutando n goroutines
	// Cada vez que se termina la goroutine se decrementa el numero de goroutines y finaliza el programa
	wg.Add(2)

	// Ejecutamos las goroutines
	go hello()
	go bay()

	fmt.Printf("Número de CPUs en el medio: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas en el medio: %v\n", runtime.NumGoroutine())

	// El hilo principal se detiene a que finalicen las goroutines
	wg.Wait()

	// Despues de que han finalizado se ejecuta lo siguiente
	fmt.Println("A punto de finalizar main.")
	fmt.Printf("Número de CPUs al final: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas al final: %v\n", runtime.NumGoroutine())

}

// Primera goroutine
func hello() {
	defer wg.Done()
	fmt.Println("Hola soy la funcion hello")
}

// Segunda goroutine
func bay() {
	defer wg.Done()
	fmt.Println("Hola soy la funcion Bay")
}
