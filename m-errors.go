package main

import (
	"errors"
	"fmt"
)

// 1 Definimos un Error Centinela
// Es una variable global que representa un estado de falla especifico y tipificado
// la variable no almacena el string sino un puntero que identifica a este error
var ErrUsuarioNoEncontrado = errors.New("el usuario no existe en la BBDD")

// 2 Simulamos la capa de acceso a datos
func buscarUsuarioEnBBDD(id int) error {
	// Simulamos que la busqueda del ID 10 falla
	if id == 10 {
		return ErrUsuarioNoEncontrado
	}
	return nil
}

// 3 Simulamos la capa de logica de negocio
func actualizarPerfil(id int) error {
	err := buscarUsuarioEnBBDD(id)
	if err != nil {
		// Valor agregado: Envolver el error original usando %w
		// Agregamos contexto de negocio sin perder 
		return fmt.Errorf("fallo la actualizacion del perfil para el ID %d: %w", id, err)
	}
	return nil
}

func main() {
	// Intentamos actualizar un perfil que sabemos que va a fallar
	err := actualizarPerfil(10)

	if err != nil {
		// Imprimimos el error envuelto. Vas a notar como contatena los textos.
		fmt.Printf("Log del sistema: %v \n\n", err)

		// 4 Inpeccionamos la cadena de errores
		// errors.Is desenvuelve el error recursivamente buscando una coincidencia exacta
		// Compara direcciones de memoria de los punte errores, no compara los strings, asi que no hay problema alli.
		if errors.Is(err, ErrUsuarioNoEncontrado) {
			fmt.Println("Decision de negocio: El usuario no existe, vamos a redirigirlo a la pantalla de registro.")
		} else {
			fmt.Println("Decision de negocio: Ocurrio un error tecnico desconocido, contactar a soporte.")
		}
	} else {
		fmt.Println("Perfil actualizado con exito.")
	}
}