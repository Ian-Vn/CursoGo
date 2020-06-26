package main

import "fmt"

// Creando la estructura persona
type persona struct {
	Nombre   string
	Apellido string
	Sabores  []string
}

// Ejercicio 3
// Creando una estructura de vehiculo
type vehiculo struct {
	puertas int
	color   string
}

// Creando dos estructuras que heredan de vehiculo
type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

// Fin del ejercicio 3

func main() {
	// Creando un tipo de dato de persona
	persona1 := persona{
		Nombre:   "Ian",
		Apellido: "Vn",
		Sabores:  []string{"Vainilla", "Chocolate", "Fresa"},
	}

	persona2 := persona{
		Nombre:   "Lucia",
		Apellido: "Mendez",
		Sabores:  []string{"Vainilla", "Chocolate", "Fresa", "Limon", "Mango"},
	}

	// Imprimiendo los valores de los sabores
	for _, value := range persona1.Sabores {
		fmt.Printf("%s \t", value)
	}
	fmt.Println()
	// --------- Ejercicio 2 ----------- */
	// Creando un mapa con las claves como apellido de cada instancia y el valor sera la persona
	mapa := map[string]persona{
		persona1.Apellido: persona1,
		persona2.Apellido: persona2,
	}

	// Recorriendo el slice de cada objeto
	for _, value := range mapa {
		fmt.Println("Nombre: ", value.Nombre)
		fmt.Println("Apellido: ", value.Apellido)
		for _, i := range value.Sabores {
			fmt.Printf("\t Sabores: %s", i)
		}
		fmt.Println()
	}

	// ---------------- Ejercio 3 --------------------------
	// Creando variables del tipo camion y sedan
	camion1 := camion{
		vehiculo:     vehiculo{puertas: 4, color: "Azul"},
		cuatroRuedas: true,
	}

	sedan1 := sedan{
		vehiculo: vehiculo{puertas: 2, color: "Rojo"},
		lujoso:   false,
	}

	fmt.Println("Camion 1: ", camion1)
	fmt.Println("Sedan 1:", sedan1)
	// Imprimiendo caracteristicas
	fmt.Println("Total de puertas del camion 1: ", camion1.puertas)
	fmt.Println("Total de puertas del sedan 1: ", sedan1.puertas)

	// ----------- Ejercicio 5 -----------------
	// Creando un struct anonimo
	anonimo := struct {
		nombre   string
		apellido string
		edad     int
		logeado  bool
	}{
		nombre:   "Ian",
		apellido: "Vn",
		edad:     24,
		logeado:  false,
	}
	// Imprimiendo el struct anonimo
	fmt.Println("Anonimo", anonimo)
	fmt.Println("Edad del anonimo", anonimo.edad)

}
