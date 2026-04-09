package main

import "fmt"

// 1 Definimos el comportamiento esperado mediante una interaz
type ProcesadorPago interface {
	Procesar(monto float64) string
}

// 2 Definimos nuestro primer struct (estado)
type TarjetaCredito struct {
	NumeroTarjeta string
	Titular	string
}

// Implementamos el metodo Procesar para TarjetaCredito
func (t TarjetaCredito) Procesar(monto float64) string {
	return fmt.Sprintf("Cobrando $%.2f a la tarjeta terminada en %s (Titular: %s)", monto, t.NumeroTarjeta[len(t.NumeroTarjeta)-4:], t.Titular)
}

// 3 Definimos nuestro segundo struct
type TransferenciaBancaria struct {
	CBU string
}

// Implementamos el metodo Procesar para TransferenciaBancaria
func (t TransferenciaBancaria) Procesar(monto float64) string {
	return fmt.Sprintf("Transfiriendo $%.2f desde el CBU %s", monto, t.CBU)
}

// 4 La magia del polimorfismo: esta funcion acepta CUALQUIER tipo que cumpla la interfaz
func EjecutarCobro(metodo ProcesadorPago, monto float64) {
	fmt.Println("Iniciando transaccion...")
	// Llamamos al metodo sin importar que struct lo este ejecutando por debajo
	resultado := metodo.Procesar(monto)
	fmt.Println(resultado)
	fmt.Println("Finalizando transaccion.\n")
}

func main() {
	// Instanciamos nuestros structs
	miTarjeta :=  TarjetaCredito{NumeroTarjeta: "123123123", Titular: "Juan Perez"}
	miBanco := TransferenciaBancaria{CBU: "010101010101"}

	// Pasamos distintos tipos de structus a la misma funcion
	EjecutarCobro(miTarjeta, 1500.55)
	EjecutarCobro(miBanco, 8500.00)
}