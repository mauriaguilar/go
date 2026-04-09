package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 1 Definimos la estructura de datos
// Las etiquetas (tags) 'json:"..."' le indican a Go como nombrar los campos al convertirlos
type Transaccion struct {
	ID      string    `json:"id"`
	Monto   float64   `json:"monto"`
	Detalle string    `json:"detalle"`
	Fecha   time.Time `json:"fecha"`	
}

// 2 Simulamos una base de datos en memoria y su prpoteccion concurrente
var (
	transacciones	[]Transaccion
	mu				sync.Mutex // Protege el acceso al slice anterior
)

// HANDLERS (Controladores)

// listarTransacciones maneja GET /transactions
func listarTransacciones(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convertimos el slice a JSON y lo enviamos directamente al cliente (w)
	json.NewEncoder(w).Encode(transacciones)
}

// crearTransaccion manjea POST /transactions
func crearTransaccion(w http.ResponseWriter, r *http.Request) {
	var nuevaTransaccion Transaccion

	// Leemos el cuerpo de la petición (que es un flujo de bytes continuo)  y lo decodificamos en nuestro struct
	// NewDecoder: Se conecta al flujo entrante (r.Body) que contiene texto en formato JSON.
	// Decode: Lee ese flujo progresivamente y mapea los datos a las propiedades del struct 'nuevaTransaccion'.
	err := json.NewDecoder(r.Body).Decode(&nuevaTransaccion)
	if err != nil {
		http.Error(w, "Cuerpo de peticion invalido", http.StatusBadRequest)
		return
	}

	// Asignamos una fecha de creacion automatica
	nuevaTransaccion .Fecha = time.Now()

	// Bloqueamos para escritura segura
	mu.Lock()
	transacciones = append(transacciones, nuevaTransaccion)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaTransaccion)
}

// obtenerTransaccion maneja GET /transactions/{id}
func obtenerTransaccion(w http.ResponseWriter, r *http.Request) {
	// Extraemos el ID directamente de la URL gracias al nuevo enrutador de Go
	idBuscado := r.PathValue("id")

	mu.Lock()
	defer mu.Unlock()

	// Buscamos la transaccion linealmente
	for _, t := range transacciones {
		if t.ID == idBuscado {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	// Si el bucle termina y no retorno, el ID no existe
	http.Error(w, "Transaccion no encontrada", http.StatusNotFound)
}

func main() {
	// 3 Inicializamos el enrutador (multiplexor)
	mux := http.NewServeMux()

	// 4. Registramos las rutas
	// Observa la sintaxis moderna: "METODO /ruta"
	mux.HandleFunc("GET /transactions", listarTransacciones)
	mux.HandleFunc("POST /transactions", crearTransaccion)
	mux.HandleFunc("GET /transactions/{id}", obtenerTransaccion)

	// 5 Levantamos el servidor
	puerto := ":8080"
	fmt.Printf("Servidor API iniciado correctamente. Escuchando en el puerto %s...\n", puerto)

	// ListenANdServe es una operacion bloqueante. Si falla, el programa termina.
	err := http.ListenAndServe(puerto, mux)
	if err != nil {
		fmt.Printf("Error critico en el servidor: %v\n", err)
	}
}


// curl -X POST http://localhost:8080/transactions \
//   -H "Content-Type: application/json" \
//   -d '{"id": "tx-001", "monto": 15500.50, "detalle": "Supermercado"}'

// curl -X GET http://localhost:8080/transactions

// curl -X GET http://localhost:8080/transactions/tx-001