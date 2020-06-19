package main

import "fmt"

func main() {
	// Ejericio 1. Creando un array
	x := [5]int{1, 2, 3, 4, 5}
	// Iterando sobre el array
	for i, v := range x {
		fmt.Printf("Indice: %d Valor: %d \n", i, v)
	}
	// Imprimiendo el tipo de dato del array
	fmt.Printf("%T \n", x)

	// Ejercicio 2. Crear un slice, un slice es como array dinamico
	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Iterando sobre el slice
	for i, v := range y {
		fmt.Printf("Indice: %d Valor: %d \n", i, v)
	}
	// Imprimiendo el tipo de dato del array
	fmt.Printf("%T \n", y)

	// Haciendo slicing, es decir, particionando el slice
	fmt.Println(y[:5])
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println(y[1:6])

	/*---------- Ejercicio 4 -----------*/
	y = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// Agregando un valor al slice y
	y = append(y, 52)
	// Imprimiendo el slice
	fmt.Println("Slice: ", y)
	// Agregando nuevos valores al slice
	y = append(y, 53, 54, 55)

	// Imprimir el slice
	fmt.Println("Slice: ", y)

	// Agregando un slice al otro slice
	z := []int{56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println("Slice: ", y)

	/*---------- Ejercicio 5 -----------*/
	// Usando append y slicing
	y = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	w := append(y[:3], y[6:]...)
	fmt.Println("Slice w: ", w)

	/*---------- Ejercicio 6 -----------*/
	// Slice multidimensional
	matriz := [][]string{
		[]string{"Batman", "Jefe", "Vestido de negro"},
		[]string{"Robin", "Ayudanete", "Vestido de colores"},
	}

	// Imprimiendo la mtriz
	fmt.Println("Matriz", matriz)

	// Haciendo range sobre todos los datos
	for i, v := range matriz {
		fmt.Println(i, v)
	}

	// Haciendo range sobre cada elemento
	for i, v := range matriz {
		for j, w := range v {
			fmt.Printf("Elemento [%d][%d]: %s \n", i, j, w)
		}
	}

	// Creando mapas. Similar a un objeto de JS
	user := map[string][]string{
		"eduard":       []string{"computadoras", "montaña", "playa"},
		"carlos_ramon": []string{"leer", "comprar", "musica"},
		"juan_bimba":   []string{"helado", "pintar", "bailar"},
	}

	// recorriendo el map
	for key, value := range user {
		fmt.Printf("Clave: %s \n", key)
		// Recorriendo el slice como valor
		for i, v := range value {
			fmt.Printf("\t indice: %d Valor: %s \n", i, v)
		}
	}

	/*---------- Ejercicio 7 -----------*/
	// Agregando un nuevo objeto al mapa anterior
	user["Ian"] = []string{"Uno", "Dos", "Tres"}
	// Imprimiendo el mapa
	fmt.Println("Mapa: ", user)
	// Recorriendo de nuevo el mapa
	for key, value := range user {
		fmt.Printf("Clave: %s \n", key)
		// Recorriendo el slice como valor
		for i, v := range value {
			fmt.Printf("\t indice: %d Valor: %s \n", i, v)
		}
	}

	/*---------- Ejercicio 8 -----------*/
	// Eliminar un objeto del mapa
	delete(user, "Ian")
	// Recorriendo de nuevo el mapa
	fmt.Println()
	for key, value := range user {
		fmt.Printf("Clave: %s \n", key)
		// Recorriendo el slice como valor
		for i, v := range value {
			fmt.Printf("\t indice: %d Valor: %s \n", i, v)
		}
	}
}
