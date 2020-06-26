package main

import "fmt"

// Ejercicio 4
type persona struct {
	nombre   string
	apellido string
	edad     int
}

// Creando un metodo de la 'clase' persona
func (p *persona) presentar() {
	fmt.Println("Hola soy", p.nombre, p.edad)
}

func main() {

	// Ejercicio 3: Uso de defer, cuando se usa defer la funcion que va delante de la palabra clave defer
	// sera ejecutada cuando la funcion contenedora finalice o retorne
	defer final()

	// Ejericio 1
	fmt.Println(foo())
	numero, saludo := bar()
	fmt.Println("Dijo", saludo, numero, "veces")

	// Ejercicio 2
	// Creando un slice
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Imprimiendo desde fooA: ", fooA(numeros...))

	// Llamando a la funcion barA que recibe un slice
	fmt.Println("Imprimiendo desde barA: ", barA(numeros))

	// Ejercicio 4
	// Creando un tipo de dato de persona
	persona1 := persona{nombre: "Ian", apellido: "Vn", edad: 25}
	// Llamando al metodo de la estructura persona
	persona1.presentar()

}

// ----------- Ejercicio 1 -----------
// Crear una funcion con un solo retorno int (foo) y otra con dos retornos que son de int y string (bar)

func foo() int {
	return 20
}

func bar() (int, string) {
	return 10, "Hello world"
}

// ------------- Ejercicio 2 -----------
// Crear una funcion que tenga parametros variables
func fooA(numero ...int) int {
	// Haciendo la suma de los numero pasados como parametros
	suma := 0
	for _, i := range numero {
		suma += i
	}
	return suma
}

// Funcion que toma como parametro un slice de enteros
func barA(numeros []int) int {
	// Haciendo la suma de los numero pasados como parametros
	suma := 0
	for _, i := range numeros {
		suma += i
	}
	return suma
}

// --------------- Ejercicio 3 -----------------------------
// funcion que sera llamada con la palabra clave defer
func final() {
	fmt.Println("Funcion ejecutada al final de main")
}
