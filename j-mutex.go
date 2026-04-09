package main

import (
	"fmt"
	"sync"
)

func main() {
	var contador int
	var wg sync.WaitGroup

	// Declaramos el Mutex. Al igual que el WaitGroup, su valor cero ya es utilizable
	var mu sync.Mutex

	// Lanzamos 1000 goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Aca esta el peligro: todas las goroutines acceden a la misma
			// variable sin proteccion
			mu.Lock()
			contador++
			mu.Unlock()
		}()
	}

	wg.Wait()

	// Si todo fuera perfecto, el contador deberia imprimir exactamente 1000
	fmt.Printf("Main: El valor final del contador es %d.\n", contador)
}


// go run -race j-mutex.go 
// Main: El valor final del contador es 1000.