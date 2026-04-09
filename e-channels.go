package main

import (
	"fmt"
	"time"
)

// Buffered chanels

func producirNumeros(canal chan int) {
	// Vamos a enviar 5 numeros al canal
	for i := 1; i <= 5; i++ {
		fmt.Printf("Productor: Enviando el numero %d...\n", i)
		canal <- i
		time.Sleep(1300 * time.Millisecond) // Simulamos un breve tiempo de proceso
	}

	// Una vez que terminamos de vniar todo, cerramos el canal
	// Regla de oro: El cierre SIEMPRE debe hacerlo la goroutine que envia
	// nunca la que recibe
	close(canal)
	fmt.Println("Productor: Tarea finalizada y canal cerrado.")
}

func main() {
	// 1 Creamos un canal de tipo int con capacidad 3
	canalNumeros := make(chan int, 3)

	// 2 Lanzamos la goroutine pasandole el canal como argumento
	go producirNumeros(canalNumeros)

	// 3 Utilizamos range para consumir los datos del canal a medida que ingresan.
	// Este bucle intentara de forma segura y finalizara solo cuando canalNumeros sea cerrado.
	for numero := range canalNumeros {
		fmt.Printf("Consumidor (Main): Recibi el numero %d\n", numero)
	}

	fmt.Println("Main: Todos los datos fueron recibidos exitosamente.")
}
