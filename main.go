package main

import (
	"TP-2024-malo/code"
	"fmt"
)

func main() {
	lle := code.GetEjercicios()
	// Creo que estos checks no estan implementados para estas funciones, osea se asume que
	// lo que esta guardado en el csv es válido.
	if lle == nil {
		fmt.Println("Hubo un error en la inicialización de los ejercicios.")
	}
	llr := code.GetRutinas(lle)
	if llr == nil {
		fmt.Println("Hubo un error en la inicialización de las rutinas.")
	}

	// Dentro de esta función se va a desarrollar el main loop con la interacción
	// medinate CLI con el programa.
	state := true
	for state {
		state = code.MainMenu(lle, llr)
	}

}
