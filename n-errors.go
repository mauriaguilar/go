package main

import (
	"errors"
	"fmt"
)

func main() {
	// Creamos dos errores distintos pero con EXACTAMENTE el mismo texto
	err1 := errors.New("error de conexion")
	err2 := errors.New("error de conexion")

	// Comparamos si son el mismo error
	if errors.Is(err1, err2) {
		fmt.Println("Go cree que son el mismo error (comparó textos).")
	} else {
		fmt.Println("Go sabe que son errores distintos (comparó punteros de memoria).")
	}
	
	// Para ver los punteros reales en memoria, usamos el verbo %p
	fmt.Printf("Dirección en memoria de err1: %p\n", err1)
	fmt.Printf("Dirección en memoria de err2: %p\n", err2)
}
