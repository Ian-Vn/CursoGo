package main

import "fmt"

func main() {
	// Asignando un int a una variable
	numero := 10

	// Imprimiendo el numero en decimal, binario y hexadecimal
	fmt.Printf("Decimal: %d \t Binario: %b \t Hexadecimal: %x \n", numero, numero, numero)

	// Haciendo el desplazamiento de bits a la izquierda
	numero = numero << 1

	// Imprimiendo el numero en decimal, binario y hexadecimal
	fmt.Printf("Decimal: %d \t Binario: %b \t Hexadecimal: %x \n", numero, numero, numero)
}
