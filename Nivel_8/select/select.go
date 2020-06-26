package main

import "fmt"

func main() {

	// Creamos tres canales
	p := make(chan int)
	q := make(chan int)
	r := make(chan int)

	// Hacemos la goroutine
	go enviar(p, q, r)

	// Ejecutamos la funcnio de recibir como parte del hilo principal para que unicamente estos dos hilos
	// se comuniquen
	recibir(p, q, r)
}

// Goroutine para poder escribir en los canales, es de solo escritura
func enviar(p, q, r chan<- int) {

	// Hacemos una recorrido para poder insertar los numeros pares en el canal p
	// los impares en el canal q, y para salir el canal r
	for i := 0; i <= 100; i++ {
		// Numeros pares
		if i%2 == 0 {
			// Escribimos al canal p
			p <- i
		} else {
			// Escribimos los numeros impares al canal q
			q <- i
		}
	}

	// Cuando se acabe el ciclo emitimos un valor 0 al canal r el cual denotara que
	// terminará la lectura del canal sin necesidad de cerrar el canal
	fmt.Println("Terminando goroutine...")
	r <- 0
}

// Funcion que solo recibe los datos del canal
func recibir(p, q, r <-chan int) {
	// Iniciamos un ciclo infinito
	for {
		// Usamos la sentencia select para indicar que cuando se recibe un determinado valor
		// en un canal se ejecuta una instruccio
		select {
		// Lo que indicamos es que cuando recibimos un valor del canal p lo asignamos a la variable v

		case v := <-p:
			fmt.Println("Desde el canal p", v)
		// Lo que indicamos es que cuando recibimos un valor del canal q lo asignamos a la variable v
		case v := <-q:
			fmt.Println("Desde el canal p", v)
		// Lo que indicamos es que cuando recibimos un valor del canal r entonces salimos de for infinito
		// y se acaba la funcion recibir por lo cual
		case <-r:
			fmt.Println("Se terminó la funcion recibir...")
			return
		}
	}
}
