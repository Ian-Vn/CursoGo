// Range deja de leer desde un canal cuando esta cerrado
package main

import "fmt"

func main() {

	// Creamos un canal
	c := make(chan int)

	// Ejecutamos la goroutine
	go escribir(c)

	// Hacemos un ciclo para leer datos del canal
	// NOTA: EL CICLO SEGUIRA LEYENDO O ESPERANDO VALORES DEL CANAL, CUANDO ESTO PASA
	// EL HILO QUE EJECUTA EL RANGE ESPERARA A QUE EL CANAL EMITA UN NUEVO VALOR,
	// SI EL CANAL HA DEJADO DE EMITIR VALORES Y NO FUE CERRADO HARÁ QUE EL HILO ESPERE
	// INDEFINIDAMENTE Y POR LO TANTO CAERA EN UN DEADLOCK, POR EL CONTRARIO SI EL CANAL SE CIERRA
	// ENTONCES EL HILO DE LA GOROUTINE SE DARÁ CUENTA QUE EL CANAL ESTA CERRADO Y POR LO TANTO
	// NO ESPERARA A RECIBIR MAS VALORES Y LA EJECUCION CONTINUA
	for value := range c {
		// Se termina el ciclo cuando el canal se cierra
		fmt.Println("Valor del canal", value)
	}

	// Si ponemos esto cuando el canal esta cerrado, el valor se inicializa al valor cero, porque hasta este momento
	// el canal esta cerrado, en pocas palabras, CUANDO SE ESTA ESPERANDO UN VALOR Y EL CANAL NO HA SIDO CERRADO
	// LA GOROUTINE SE BLOQUEA Y PASA EL CONTROL A OTRAS GOROUTINAS, SI NO HAY OTRAS, CAE EN DEADLOCK, SI EL CANAL ESTA
	// CERRADO ENTONCES LOS VALORES ESPERADOS DEL CANAL SE INICIALIZAN AL VALOR CERO
	c <- 42
	fmt.Println("Valor despues del range: ", <-c)

}

// Goroutine para poder escribir sobre el canal
func escribir(c chan int) {
	// Ciclo para insertar datos en el canal
	for i := 0; i <= 10; i++ {
		// Escribiendo al canal
		c <- i
	}
	// CERRAMOS EL CANAL PARA QUE SE NOTIFIQUE A LAS GOROUTINES QUE ESPERAN UN VALOR DE DICHO CANAL, CON ESTO
	// SE EVITA EL DEADLOCK
	fmt.Println("Canal cerrado...")
	close(c)
	fmt.Println("Goroutine terminada...")
}
