package main

import (
	"fmt"
	"sync"
)

func main() {
	var contador int
	var wg sync.WaitGroup

	// Lanzamos 1000 goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Aca esta el peligro: todas las goroutines acceden a la misma
			// variable sin proteccion
			contador++
		}()
	}

	wg.Wait()

	// Si todo fuera perfecto, el contador deberia imprimir exactamente 1000
	fmt.Printf("Main: El valor final del contador es %d.\n", contador)
}


// go run -race i-mutex.go 
// ==================
// WARNING: DATA RACE
// Read at 0x00c000112028 by goroutine 16:
//   main.main.func1()
//       /home/mau/dev/go/i-mutex.go:19 +0x7b

// Previous write at 0x00c000112028 by goroutine 11:
//   main.main.func1()
//       /home/mau/dev/go/i-mutex.go:19 +0x8d

// Goroutine 16 (running) created at:
//   main.main()
//       /home/mau/dev/go/i-mutex.go:15 +0x78

// Goroutine 11 (finished) created at:
//   main.main()
//       /home/mau/dev/go/i-mutex.go:15 +0x78
// ==================
// ==================
// WARNING: DATA RACE
// Write at 0x00c000112028 by goroutine 9:
//   main.main.func1()
//       /home/mau/dev/go/i-mutex.go:19 +0x8d

// Previous write at 0x00c000112028 by goroutine 13:
//   main.main.func1()
//       /home/mau/dev/go/i-mutex.go:19 +0x8d

// Goroutine 9 (running) created at:
//   main.main()
//       /home/mau/dev/go/i-mutex.go:15 +0x78

// Goroutine 13 (running) created at:
//   main.main()
//       /home/mau/dev/go/i-mutex.go:15 +0x78
// ==================
// Main: El valor final del contador es 751.
// Found 2 data race(s)
// exit status 66