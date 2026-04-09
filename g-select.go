package main

import (
	"fmt"
	"time"
)

func consultarBBDD(canal chan string) {
	// Simulamos una consulta lenta que demora 1 / 3 segundos
	time.Sleep(1 * time.Second)
	canal <- "Resultado de la consulta: 100 usuarios encontrados."
}

func main() {
	canalRespuesta := make(chan string)

	// Lanzamos la consulta en segundo plano
	go consultarBBDD(canalRespuesta)

	fmt.Println("Main: iniciando consulta con un tiempo maximo de espera de 2 segundos...")

	// El bloque select esperara al primer canal que reciba un dato
	select {
		case resultado := <-canalRespuesta:
			// Este caso se ejecuta si la base de datos responde a tiempo
			fmt.Printf("Main: Operacion exitosa. %s\n", resultado)
		
		case <-time.After(2 * time.Second):
			// Este caso se ejecuta si pasan 2 segundos y el canalRespuesta sigue vacio
			fmt.Println("Main: Error - Tiempo de espera afotado (timeout). Cancelando la operacion para liberar recursos.")
			// Valor agregado: Este concepto es un patrón de diseño fundamental en Go.
			// Se lo conoce como usar un canal como señal (signaling channel).
			// Es la forma más limpia y idiomática de notificar eventos entre goroutines cuando el dato transportado es irrelevante
			// y lo único que tiene valor es el momento en que ocurre la comunicación.

	}
}