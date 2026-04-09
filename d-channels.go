package main

import (
	"fmt"
	"time"
)

// Unbuffered chanels

func generarMensaje(canal chan string) {
	fmt.Println("Goroutine: Procesando informacion pesada...")
	time.Sleep(2 * time.Second) // Simulamos una tarea de 2 segundos

	// Enviamos un dato HACIA el canal usando el operador <-
	canal <- "Mision cumplida desde goroutine!"
}

func main() {
	// 1 Creamos un canal de tipo string
	mensajeChan := make(chan string)

	// 2 Lanzamos la goroutine pasandole el canal como argumento
	go generarMensaje(mensajeChan)

	fmt.Println("Main: hice mi parte, ahora me quedo esperando el mensaje...")

	// 3 Recibimos el datos DESDE el canal usando el operador <-
	// La ejecucion de main se bloqueara automaticamente en esta linea
	// hasta que el dato ingrese al canal
	respuesta := <-mensajeChan

	fmt.Printf("Main: Recibi el siguiente mensaje: %s\n", respuesta)
}