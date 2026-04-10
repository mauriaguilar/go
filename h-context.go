package main

import (
	"context"
	"fmt"
	"time"
)

// Context es el primer param, este es el estandar en Go.
func procesarDatos(ctx context.Context, canalResultado chan string) {
	fmt.Println("Goroutine: Iniciando procesamiento pesado (estimado 4 segundos)....")

	// Canal para recibir la señal del trabajo completado
	canalTrabajo := make(chan string)

	// Lanzar el trabajo real en una goroutine
	go func() {
		time.Sleep(4 * time.Second) // Simulamos el trabajo pesado
		canalTrabajo <- "ok, terminé el trabajo"
	}()

	// Utilizamos un select para competir: termina el trabajo o se cancela?
	select {
	case resultado := <-canalTrabajo: // Esperar la señal del trabajo completado
		canalResultado <- "Goroutine: Trabajo finalizado exitosamente. " + resultado
	case <-ctx.Done(): // Este canal emite una señal si el contexto expira o se cancela
		mensajeError := fmt.Sprintf("Goroutine: Trabajo abortado. Motivo: %v.", ctx.Err())
		canalResultado <- mensajeError
	}
}

func main() {
	// 1 Creamos un contexto base vacio llamado Background
	ctxBase := context.Background()

	// 2 Derivamos un nuevo contexto con un tiempo maximo de vida de 2 segundos
	// Esta funcion devuelve el nuevo contexto y una funcion cancel para liberar recursos
	ctx, cancel := context.WithTimeout(ctxBase, 2*time.Second)

	// Regla estricta: Siempre debemos llamar a cancel() al salir de la funcion
	// para evitar fugas de memoria (goroutine leaks), incluso si la tarea terminó a tiempo
	defer cancel()

	canalRespuesta := make(chan string)

	// 3 Lanzamos la goroutine pasandole el contexto
	go procesarDatos(ctx, canalRespuesta)

	fmt.Println("Main: Esperando el resultado de la goroutine...")

	// 4 Nos bloqueamos esperando la respuesta por el canal
	resultado := <-canalRespuesta
	fmt.Printf("Main: Recibi el siguiente reporte:\n-> %s\n", resultado)
}
