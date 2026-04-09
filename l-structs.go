package main

import "fmt"

// 1 Definimos nuestro struct base con estado y comportamiento
type Persona struct {
	Nombre string
	Edad int
}

func (p Persona) Presentarse() {
	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", p.Nombre, p.Edad)
}

// 2 Definimos el struct avanzado usando Composicion
type Empleado struct {
	Persona // Incrustacion anonima: no le ponemos nombre a la variable, solo el tipo
	Cargo string
	Sueldo float64
}

// Opcional: Empleado puede tener sus propios metodos
func (e Empleado) Trabajar() {
	fmt.Printf("%s esta trabajando como %s.\n", e.Nombre, e.Cargo)
}

func main() {
	// instanciamos el Empleado
	emp := Empleado{
		Persona: Persona{
			Nombre: "Carlos",
			Edad: 35,
		},
		Cargo: "Desarrollador Backend",
		Sueldo: 1500000.00,
	}

	// Promocion de campos: accedemos a Nombre directamente desde emp
	// sin necesidad de hacer em.Persona.Nombre (aunque tambien es valido)
	fmt.Printf("Ficha: %s - %s\n", emp.Nombre, emp.Cargo)

	// Promocion de metodos: emp usa el metodo de Persona direcamente
	emp.Presentarse()

	// Y obbiamente usa sus propios metodos
	emp.Trabajar()
}