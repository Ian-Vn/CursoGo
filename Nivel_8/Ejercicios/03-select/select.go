package main

import "fmt"

func main() {
	// Se inicia el hilo principal
	salir := make(chan int)
	c := gen(salir)
	recibir(c, salir)
	fmt.Println("A punto de finalizar")

}

// Función que acepta como parametro un canal de solo lectura y devuelve un canal de solo lectura
func gen(salir chan<- int) <-chan int {

	// canal bidireccional
	c := make(chan int)

	// Se itera para escribir datos sobre el canal
	// se escribe una funcion anonima para iniciar otro hilo, esto devuelve el control al hilo principal
	go func() {
		for i := 0; i <= 10; i++ {
			// enviando datos al canal
			c <- i
		}
		// cuando acabe el ciclo enviamos un valor al canal salir
		salir <- 1
	}()

	// regresando el canal de solo lectura
	return c
}

// Funcion para recibir dos canales de solo lectura
func recibir(c, salir <-chan int) {

	// Iteramos sobre cada canal, esta funcion como es parte del hilo principal,
	// esperara a que se transmitan valores por algun canal, cuando eso pase
	// se ejecutará la linea de codigo correspondiente
	for {
		// Con select esperamos la señal de cada canal
		select {
		case v := <-c:
			// Si hacemos esto cuando llegue a 100 no compilará ya que esperar un valor del canal C y como la goroutine
			// de gen esta esperando a que otra tome el valor del canal salir tambien se bloqueará por cual se tendra un deadloc
			// fmt.Println("Valor recibido del canal c",<-c)
			fmt.Println("Valor del canal", v)
		case <-salir:
			fmt.Println("Saliendo del main...")
			return
		}
	}
}
