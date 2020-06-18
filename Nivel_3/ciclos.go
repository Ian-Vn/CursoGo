package main

import "fmt"

func main() {

	// Ciclo for: Imprimir numeros del 1 al 10,000
	for i := 1; i <= 10000; i++ {
		fmt.Printf("i: %d \n", i)
	}

	// Ciclo doble e impresion de los caracteres ascii
	for i := 65; i <= 90; i++ {
		fmt.Printf("Numero: %d \n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t %#U \n", i)
		}
	}

	// Ciclo como condicion
	anios := 1
	for anios <= 24 {
		fmt.Printf("Año: %d \n", anios)
		anios++
	}

	anios = 1
	// Ciclo como condicion infinita
	for {
		fmt.Printf("Años: %d \n", anios)
		if anios >= 24 {
			break
		}
		anios++
	}

	// Trabajando con modulo
	for i := 10; i <= 100; i++ {
		// Equivalente a 10/4
		fmt.Printf("Numero: %d Resto: %d \n", i, i%4)
	}

	// Uso de if-else if-else, es importante la identacion
	numero := 10
	if numero == 5 {
		fmt.Println("Es un numero igual a 5")
	} else if numero < 5 {
		fmt.Println("Es un numero menor a 5")
	} else {
		fmt.Println("Es un numero mayor a 5")
	}

	// Uso de switch sin una expresion, al hacer esto se evalua cada concidion
	// y la primera en ser true es ejecutada
	fmt.Println("------- Uso de switch/case ----------")
	switch {
	case 1 > 2:
		fmt.Println("Esto es imposible")
	case 1 < 2:
		fmt.Println("Esto es cierto")
	}

	// Uso de switch con una expresion
	deportefav := "Luchas"
	fmt.Println("----- Switch/case con expresion ----- ")
	switch deportefav {
	case "Futbol":
		fmt.Println("Es futbol")
	case "Luchas":
		fmt.Println("Es luchas")
	case "Basquetball":
		fmt.Println("Es basquetball")
	default:
		fmt.Println("No es ninguna")
	}
}
