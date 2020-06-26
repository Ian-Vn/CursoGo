package main

import "fmt"

// Ejercicio 2
// Crear una estructura
type persona struct {
	nombre string
	edad   int
}

func main() {

	// Apuntadores
	// ------------- Ejercicio 1 -----------------------------------
	// Crear una variables
	nombre := "Ian"
	// Se imprimime la direccion de memoria donde esta la variable
	fmt.Println("Informacion: ", nombre, " Localidad de mememoria: ", &nombre)

	// Ejercicio 3
	// Creando una variable de tipo persona
	al := persona{nombre: "Juan", edad: 24}
	// Imprimimos la variable a1
	fmt.Println("Antes de pasar a la funcion", al)
	// Llamando a la funcion
	cambiame(&al)
	fmt.Println("Despues de llamar a la funcion", al)
}

// Funcion que recibe como parametro un apuntador al tipo persona
func cambiame(p *persona) {
	// Como p es un apuntador a persona, se puede obtener los campos del struct por
	// desreferenciacion
	(*p).nombre = "Cambiado por puntero"
	(*p).edad = 12345

}
