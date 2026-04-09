package main

import (
	"fmt"
	"sync" // --> para utilizar WaitGroups y sincronizar goroutines
	"time"
)

// Recibimos un puntero al WaitGroup para modificar el original, no una copia
func imprimirNumeros(wg *sync.WaitGroup) {
	// defer asegura que Done se ejecute siempre,
	// justo antes de salir de esta funcion
	defer wg.Done()
	
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(800 * time.Millisecond)
	}
}

func imprimirLetras(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 4; i++ {
		letras := []string{"A", "B", "C", "D", "E"}
		fmt.Println(letras[i])
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Declaramos el contador
	var wg sync.WaitGroup

	// Le indicamos al contador que vamos a esperar exactamente 2 goroutines
	wg.Add(2)

	// Lanzamos ambas funciones de forma concurrentes
	// Le pasamos la direccion de memoria del wg
	go imprimirNumeros(&wg)
	go imprimirLetras(&wg)

	// El programa se detiene en esta linea hasta que ambas
	// goroutines llamen a wg.Done()
	wg.Wait()

	fmt.Println("Todas las goroutines finalizaron. Cerrando main() de forma segura.")

}