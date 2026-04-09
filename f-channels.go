package main

import (
	"fmt"
	"time"
)

// Buffered chanels

// El verdadero poder de un canal con búfer (buffered channel) se ve cuando
// el productor es mucho más rápido que el consumidor.
// El búfer actúa como un amortiguador, permitiendo que el productor siga trabajando
// y dejando mensajes "en la cola" hasta que el buzón se llene,
// momento en el cual recién se bloqueará.

func producirNumeros(canal chan int) {
	// Vamos a enviar 5 numeros al canal
	for i := 1; i <= 5; i++ {
		fmt.Printf("Productor: Enviando el numero %d...\n", i)
		canal <- i
		fmt.Printf("Productor: Ingresé el numero %d al bufer.\n", i)
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

	// 3 Pausamos el main 1 segundo para darle ventaja al productor y que llene el buffer
	time.Sleep(1 * time.Second)
	fmt.Println("-----El consumidor despierta----------")

	for numero := range canalNumeros {
		fmt.Printf("Consumidor (Main): Procesando el numero %d...\n", numero)
		// Simulamos un consumidor lento
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Main: Todos los datos fueron procesados exitosamente.")
}
