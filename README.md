# Golang

## Install

sudo snap install go --classic
go version
go 1.25.7 from Canonical✓ installed

Nota de valor: A partir de la versión 1.11, Go utiliza Go Modules para gestionar las dependencias. Esto es una gran ventaja porque ya no estás obligado a configurar la histórica variable de entorno GOPATH ni a guardar tus proyectos en un directorio específico. Podés crear tu carpeta de trabajo en cualquier lugar de tu equipo y simplemente inicializar el módulo con go mod init <nombre-del-modulo>.

## Project Layout

mi-microservicio/
├── go.mod            <-- Dependencias del proyecto
├── go.sum            <-- Checksums de dependencias
├── cmd/
│   └── server/
│       └── main.go   <-- Entry point, solo inicializa y arranca
├── internal/         <-- Código privado del proyecto
│   ├── auth/         <-- Lógica de autenticación
│   ├── storage/      <-- Capa de base de datos
│   └── service/      <-- Lógica de negocio (use cases)
├── pkg/              <-- Código reutilizable/público (opcional)
└── tests/            <-- Tests de integración (opcional)

## Basic commands

- `go run main.go` - Ejecuta el programa sin compilar
- `go build` - Compila el proyecto en un ejecutable
- `go mod init nombre` - Inicializa un nuevo módulo Go
- `go mod tidy` - Limpia dependencias no usadas y descarga las faltantes
- `go get paquete` - Descarga e instala un paquete
- `go test` - Ejecuta los tests del proyecto
- `go fmt` - Formatea el código según estándares de Go
- `go vet` - Analiza el código en busca de errores comunes
- `go install` - Compila e instala el binario en $GOPATH/bin
- `go clean` - Elimina archivos compilados y cache

## Basic concepts

- Goroutines: Funciones que se ejecutan de forma concurrente con `go func()`, livianas y eficientes
- Context: Controla timeouts, cancelaciones y deadlines en operaciones asíncronas
- Channels: Comunican datos entre goroutines de forma segura con `ch <- valor` y `valor := <-ch`
- Interfaces: Contratos implícitos que definen comportamiento, cualquier tipo que implemente los métodos cumple la interfaz
- Defer: Ejecuta una función al finalizar la función actual, útil para cleanup (cerrar archivos, unlock mutex)
- Pointers: Pasa referencias con `*` y `&`, permite modificar valores y evitar copias grandes
- Error handling: No hay excepciones, se retorna `error` explícitamente y se valida con `if err != nil`
- Structs: Define tipos personalizados con campos, la base de la OOP en Go
- Slices: Arrays dinámicos con `[]tipo`, más usados que arrays fijos `[n]tipo`
- Maps: Diccionarios clave-valor con `map[string]int`, declarar con `make()` o literal
- Methods: Funciones asociadas a un tipo con `func (r Receiver) Method()`, no hay clases
- Select: Espera múltiples operaciones de channels simultáneamente, como switch para channels
- WaitGroup: Sincroniza goroutines esperando que todas terminen con `Add()`, `Done()`, `Wait()`
- Mutex: Protege acceso concurrente a datos compartidos con `Lock()` y `Unlock()`
- Exported/Unexported: Nombres con mayúscula son públicos, minúscula son privados al package
- Embedding: Composición de structs para reutilizar código, no hay herencia clásica
- Type assertions: Convierte interface a tipo concreto con `valor.(Type)` o `valor, ok := inter.(Type)`
- Range: Itera sobre slices, maps, channels con `for i, v := range collection`
- Init function: Se ejecuta automáticamente antes de `main()`, útil para inicialización
- Blank identifier: Usa `_` para ignorar valores de retorno que no necesitas
- Zero values: Tipos tienen valores por defecto (0, false, nil, "") sin inicializar

## Key Conceptos

🧵 Goroutines: Creación, ciclo de vida y prevención de fugas (goroutine leaks).

📡 Channels: Diferencias entre unbuffered y buffered, canales direccionales y patrones de cierre seguro.

🔀 Select Statement: Multiplexación de canales, operaciones no bloqueantes y manejo de timeouts.

🔒 Sincronización Clásica: Uso de sync.WaitGroup para coordinar múltiples tareas y sync.Mutex para evitar condiciones de carrera (race conditions).

🛑 Context: Propagación de estados de cancelación y control de tiempos máximos de ejecución a través de múltiples goroutines.

🧩 Interfaces y Composición: Aplicación práctica del duck typing, polimorfismo y composición de structs (en lugar de herencia).

⚠️ Manejo Avanzado de Errores: Envoltura de errores (wrapping), errores personalizados y uso de errors.Is y errors.As.

🧪 Testing y Benchmarks: Creación de pruebas robustas usando el patrón de table-driven tests y medición de rendimiento.****

## Testing patterns

- Table-driven tests: Usa slice de structs con casos de prueba `tests := []struct{input, want}`
- Test files: Mismo nombre con `_test.go`, función `TestNombre(t *testing.T)`
- Subtests: Organiza con `t.Run("caso", func(t *testing.T) {...})`
- Mocking: Usa interfaces para inyectar dependencias fake en tests
- Benchmarks: `func BenchmarkNombre(b *testing.B)` con loop `for i := 0; i < b.N; i++`
- Coverage: `go test -cover` o `go test -coverprofile=cover.out`

## Common packages

- `fmt`: Print, formateo de strings (`Printf`, `Sprintf`, `Fprintf`)
- `net/http`: Servidor y cliente HTTP (`ListenAndServe`, `HandleFunc`, `Get`)
- `encoding/json`: Marshal/Unmarshal JSON (`json.Marshal`, `json.Unmarshal`)
- `time`: Fechas, timers, durations (`time.Now()`, `time.Sleep()`, `time.Parse()`)
- `context`: Control de cancelación y timeouts (`context.WithTimeout`, `WithCancel`)
- `log`: Logging básico (`log.Println`, `log.Fatal`, `log.Printf`)
- `os`: Sistema operativo (env vars, archivos, args) (`os.Getenv`, `os.Open`)
- `strings`: Operaciones con strings (`Split`, `Join`, `Contains`, `Replace`)

## Best practices

- Acepta interfaces, retorna structs concretos
- Keep interfaces small (1-2 métodos idealmente)
- Handle errors explícitamente, no uses `panic` salvo casos extremos
- Usa `gofmt` y `golangci-lint` para mantener código limpio
- Evita `init()` cuando sea posible, prefiere funciones explícitas
- Composition over inheritance: usa embedding en lugar de jerarquías
- Nombra paquetes con lowercase, sin guiones bajos ni camelCase

## HTTP basics

```go
// Servidor básico
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})
http.ListenAndServe(":8080", nil)

// Cliente básico
resp, err := http.Get("https://api.example.com/data")
defer resp.Body.Close()
json.NewDecoder(resp.Body).Decode(&result)
```

## Database patterns

```go
// Conexión DB
db, err := sql.Open("postgres", connStr)
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)

// Query simple
row := db.QueryRow("SELECT name FROM users WHERE id = $1", userID)
err = row.Scan(&name)

// Transacción
tx, _ := db.Begin()
tx.Exec("INSERT INTO...")
if err != nil {
    tx.Rollback()
} else {
    tx.Commit()
}
```

## Error patterns

```go
// Wrapping errors (Go 1.13+)
return fmt.Errorf("failed to connect: %w", err)

// Custom error
type MyError struct { Code int; Msg string }
func (e *MyError) Error() string { return e.Msg }

// Panic/Recover (solo casos extremos)
defer func() {
    if r := recover(); r != nil {
        log.Println("Recovered:", r)
    }
}()
```

## Environment vars

- `GOPATH`: Workspace de Go (deprecado con Go modules)
- `GOOS`: Sistema operativo target (`linux`, `darwin`, `windows`)
- `GOARCH`: Arquitectura target (`amd64`, `arm64`)
- `CGO_ENABLED`: Habilita/deshabilita CGO (0 o 1)
- Cross-compile: `GOOS=linux GOARCH=amd64 go build` ****

## Errores comunes

Error:
package command-line-arguments is not a main package

Solucion:
package main

# Concurrencia y paralelismo

El comportamiento interno de Go:
En las versiones modernas de Go, el entorno de ejecución lee la arquitectura de tu PC y configura automáticamente una variable interna llamada GOMAXPROCS para que coincida con la cantidad de núcleos lógicos de tu procesador. Por lo tanto, si tu máquina tiene múltiples núcleos, el planificador de Go tomará tus goroutines concurrentes y las distribuirá para que corran en distintos hilos del sistema operativo simultáneamente.

En conclusión: tu código es concurrente por diseño, pero se ejecuta en paralelo por la capacidad del hardware y la gestión del motor de Go.

# La Sentencia Select

¿Qué sucede si tu programa necesita prestar atención a múltiples canales en simultáneo? Si usás el operador <- de forma secuencial, tu programa se va a bloquear esperando el primer canal, ignorando completamente si los demás ya tienen información lista para ser procesada.

Acá es donde entra en escena la sentencia select.

Podés pensar en select como un conmutador o la clásica estructura switch, pero diseñada exclusivamente para operar con canales. Su valor agregado radica en que le permite a una goroutine esperar múltiples operaciones de comunicación de forma concurrente.

Características principales:

    Bloqueo inteligente: select pausa la ejecución hasta que al menos uno de sus casos (un envío o una recepción en un canal) pueda proceder.

    Selección aleatoria: Si múltiples canales están listos al mismo tiempo, select elige uno al azar para asegurar que no haya favoritismos (evitando la inanición o starvation de las demás rutinas).

    Manejo de Timeouts: Es la herramienta por excelencia en Go para evitar que un proceso se quede colgado para siempre esperando una respuesta de una base de datos o una API externa.

case <-time.After(2 * time.Second):
Valor agregado: Este concepto es un patrón de diseño fundamental en Go. Se lo conoce como usar un canal como señal (signaling channel). Es la forma más limpia y idiomática de notificar eventos entre goroutines cuando el dato transportado es irrelevante y lo único que tiene valor es el momento en que ocurre la comunicación.

# Sin clases ni herencia

Go propone un enfoque mucho más pragmático basado en dos conceptos:

Structs: Para agrupar el estado (los datos).

Interfaces: Para definir el comportamiento (los métodos).

El valor agregado: Duck Typing y la Norma General
En Go, la implementación de una interfaz es implícita. Esto se conoce coloquialmente como Duck Typing: "Si camina como un pato y hace cuac como un pato, entonces el compilador asume que es un pato". Si tu struct posee los métodos exactos que exige una interfaz, Go automáticamente considera que tu struct cumple con esa interfaz, sin que vos tengas que declararlo explícitamente en el código.

Como norma general de diseño en Go: "Aceptá interfaces, devolvé structs" (Accept interfaces, return structs). Si tus funciones piden interfaces como parámetros, tu código se vuelve inmensamente flexible y muy fácil de probar (ideal para inyectar mocks en los tests).

# Composicion

En la programación orientada a objetos tradicional, las jerarquías de herencia suelen volverse un dolor de cabeza cuando el sistema crece. Go soluciona esto mediante la Incrustación de Structs (Struct Embedding).

La regla es simple: si declarás un struct dentro de otro sin ponerle un nombre de variable, el struct "padre" hereda automáticamente todos los atributos y métodos del struct "hijo" incrustado. A esto se lo llama promoción de campos.

Como norma general de diseño: Favorecé siempre la composición por sobre la herencia. Te va a dar sistemas mucho más modulares y fáciles de mantener.

# Errores

Si creás un error nuevo que diga error leyendo configuración, perdés el error original (file not found), lo cual dificulta la depuración.

La solución: Envoltura de Errores (Error Wrapping)
A partir de la versión 1.13, Go introdujo la capacidad de "envolver" un error dentro de otro, como si fueran muñecas mamushka. De esta forma, podés agregar tu propio contexto (qué estabas intentando hacer) sin destruir el error original subyacente.

Para esto utilizamos la función fmt.Errorf con el verbo especial %w (de wrap).

Además, nos provee la función errors.Is(), que permite inspeccionar toda esa cadena de errores envueltos para preguntar: "¿Algún error en toda esta cadena es de este tipo específico?".

# Testing

Las reglas estructurales son muy simples:

El archivo: Todo archivo de pruebas debe terminar con el sufijo _test.go.

El Test Unitario: La función debe llamarse Test seguido del nombre de lo que vas a probar (con la primera letra en mayúscula), y debe recibir el parámetro t *testing.T.

El Benchmark: Para medir rendimiento, la función debe empezar con Benchmark y recibir el parámetro b *testing.B.

En la industria, el estándar absoluto para escribir pruebas en Go es el patrón Table-Driven Tests (pruebas guiadas por tablas). En lugar de escribir diez funciones distintas para probar diez escenarios, armamos una estructura de datos (una "tabla") con los valores de entrada y el resultado que esperamos, y luego iteramos sobre ella. Esto deja el código sumamente limpio y fácil de mantener.

# Headers

En la web, toda respuesta tiene dos partes: el "sobre" (los headers y el código de estado) y la "carta" (el cuerpo o body, que es tu JSON). La regla estricta e inquebrantable del protocolo HTTP es que primero se entrega el sobre y después se lee la carta.

WriteHeader sirve exclusivamente para definir el Código de Estado HTTP (200 OK, 201 Creado, 404 No Encontrado, 500 Error de Servidor).

La trampa en Go: Si usás Encode para mandar el JSON antes de usar WriteHeader, Go asume automáticamente que todo salió bien, envía un código 200 OK por defecto, y sella el sobre. Si después de eso intentás hacer un w.WriteHeader(404), la terminal te va a tirar una advertencia avisando que los encabezados ya fueron enviados y no podés cambiarlos.

Por eso el orden siempre es:
    * Armás los encabezados (w.Header().Set(...)).
    * Escribís el código de estado (w.WriteHeader(...)).
    * Enviás el cuerpo (json.NewEncoder(w).Encode(...)).

El stream es el medio de transporte y el JSON es el formato de los datos que viajan por ahí. Imaginate una manguera: el stream (r.Body) es la manguera por donde entra la información de a poco, y el JSON es el agua estructurada que viaja por adentro. No es que el decodificador convierte el flujo a JSON; el flujo ya trae texto en formato JSON.

Respecto a si asigna las propiedades a medida que llegan: sí, exactamente. El decodificador no espera a descargar los (por ejemplo) 10 Megabytes enteros del cuerpo de la petición HTTP en la memoria RAM. Va leyendo la "manguera" en pequeños bloques (chunks). Apenas lee la clave "monto": 1500, la inyecta en el struct en memoria, descarta ese texto y sigue leyendo el resto del flujo. Sin embargo, la línea de código Decode() recién te devuelve el control (pasa a la siguiente línea de tu programa) cuando termina de procesar el objeto JSON completo y cerró la llave }.

# Estructura de una API

Para una API REST como la nuestra, la estructura de carpetas quedaría exactamente así:

📂 cmd/: Es el punto de entrada de la aplicación.

    📂 api/: Acá adentro va un archivo main.go muy chiquito. Su única responsabilidad es inicializar el servidor, cargar variables de entorno y llamar a las rutas. Nada de lógica de negocio.

📂 internal/: Es la carpeta más importante en Go. Todo el código que pongas acá adentro es privado e inaccesible para otros proyectos que quieran importar tu repositorio. Aquí vive tu lógica.

    📂 models/: Acá guardamos las estructuras de datos (nuestro struct Transaccion).

    📂 handlers/: (o controllers). Acá van las funciones que reciben el http.ResponseWriter y el *http.Request. Se encargan de leer el JSON entrante y devolver la respuesta HTTP.

    📂 storage/: (o repository). Acá aislamos la base de datos. En nuestro caso, nuestro slice en memoria y el Mutex van a vivir acá, con funciones limpias como Guardar(t Transaccion) o ObtenerTodas().