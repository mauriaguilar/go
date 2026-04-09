package main

import (
	"testing"
)

// 1 La funcion de logica de negocio (en un proyecto real, esto iria en testing.go)
func Sumar(a, b int) int {
	return a + b
}

// 2 El Test Unitario usando el patron Table-Driven Test
func TestSumar(t *testing.T) {
	// Armamos nuestra tabla definiendo los datos de entrada y el resultado esperado
	casos := []struct {
		nombreCaso	string
		a, b		int
		esperado	int
	}{
		{"Nros Positivos", 2, 3, 5},
		{"Nros Negativos", -1, -2, -3},
		{"Con Cero", 0, 0, 0},
		{"Mixtos", 10, -5, 5},
	}

	// Iteramos sobre cada caso de la tabla
	for _, caso := range casos {
		// t.Run ejecuta cada escenario como un sub-test independiente
		t.Run(caso.nombreCaso, func(t *testing.T) {
			resultado := Sumar(caso.a, caso.b)

			// Si el resultado no es el esperado, reportamos el error
			if resultado != caso.esperado {
				// t.Errorf marca el test como fallido pero permite que los demas casos se sigan ejecutando
				t.Errorf("Sumar(%d, %d) = %d; se esperaba %d", caso.a, caso.b, resultado, caso.esperado)
			}
		})
	}
}

// 3 El Benchmark para medir rendimiento y consumo de CPU
func BenchmarkSumar(b *testing.B) {
	// El valor b.N es inyectado dinamicamente por Go
	// El motor ejecutará este bucle miles o millones de veces hasta obtener una medicion estadistica confiable
	for i := 0; i < b.N; i++ {
		Sumar(2, 3)
	}
}

// go test -v -bench=. o-testing_test.go
// La bandera -v (verbose) te va a mostrar el detalle de cada escenario de tu tabla.

// La bandera -bench=. le indica a Go que, además de los tests de correctitud,
// también ejecute las pruebas de rendimiento (que por defecto vienen desactivadas
// para ahorrar tiempo).


// RESULTS

/*
	=== RUN   TestSumar
	=== RUN   TestSumar/Nros_Positivos
	=== RUN   TestSumar/Nros_Negativos
	=== RUN   TestSumar/Con_Cero
	=== RUN   TestSumar/Mixtos
	--- PASS: TestSumar (0.00s)
		--- PASS: TestSumar/Nros_Positivos (0.00s)
		--- PASS: TestSumar/Nros_Negativos (0.00s)
		--- PASS: TestSumar/Con_Cero (0.00s)
		--- PASS: TestSumar/Mixtos (0.00s)
	goos: linux
	goarch: amd64
	cpu: AMD Ryzen 5 5600H with Radeon Graphics         
	BenchmarkSumar
	BenchmarkSumar-12       1000000000               0.2558 ns/op
	PASS
	ok      command-line-arguments  0.285s
*/