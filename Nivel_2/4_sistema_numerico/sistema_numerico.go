package main

import "fmt"

func main() {
	// Declaracion de un numero
	const numero = 150

	// Impresion en forma de sistema numerico
	fmt.Printf("Binario: %#b \t Decimal: %d \t Hexadecimal: %#x \n", numero, numero, numero)
	fmt.Printf("Binario: %b \t Decimal: %d \t Hexadecimal: %#x \n", numero, numero, numero)
}
