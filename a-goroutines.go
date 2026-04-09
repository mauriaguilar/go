package main

import (
	"fmt"
	"time"
)

func imprimirNumeros() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func imprimirLetras() {
	for i := 0; i <= 4; i++ {
		letras := []string{"A", "B", "C", "D", "E"}
		fmt.Println(letras[i])
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go imprimirNumeros()
	imprimirLetras()

}