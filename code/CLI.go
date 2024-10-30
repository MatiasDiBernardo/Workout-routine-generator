package code

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/untref-ayp2/data-structures/list"
)

func MainMenu(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina]) bool {
	state := true

	// Muestra las opciones del menú principal
	fmt.Println("Menú Principal")
	fmt.Println("Ingrese el número de la funcionalidad que desee acceder:")
	fmt.Println("1) Listar")
	fmt.Println("2) Buscar")
	fmt.Println("3) Agregar")
	fmt.Println("4) Borrar")
	fmt.Println("5) Modificar")
	fmt.Println("6) Generar Rutinas")
	fmt.Println("7) Salir")

	// Lee el input del usuario
	opcionElegida := getInput(1, 7)
	if opcionElegida != -1 {
		switch opcionElegida {
		case 1:
			chooseOption(lle, llr, opcionElegida-1)
		case 2:
			chooseOption(lle, llr, opcionElegida-1)
		case 3:
			chooseOption(lle, llr, opcionElegida-1)
		case 4:
			chooseOption(lle, llr, opcionElegida-1)
		case 5:
			chooseOption(lle, llr, opcionElegida-1)
		case 6:
			chooseRoutineGeneration(lle, llr)
		case 7:
			state = false
		}
	}

	return state
}
func getInput(min int, max int) int {
	reader := bufio.NewReader(os.Stdin) //*
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el input, intente de nuevo.", err)
		return -1
	}
	opcionElegida, err := validateInput(input, min, max)
	if err != nil {
		return -1
	}
	return opcionElegida
}

func getInputString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el input, intente de nuevo.", err)
		return ""
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		fmt.Println("Ingrese el campo solicitado.", err)
		return ""
	}
	return input
}

func validateInput(input string, minimum int, maximum int) (int, error) {
	// Elimina espacios
	input = strings.TrimSpace(input)

	// Convierte string a int (para verificar rango)
	intValue, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("El input tiene que ser un número.")
		return 0, errors.New("el input tiene que ser un número")
	}

	// Valida el rango de opciones
	if intValue >= minimum && intValue <= maximum {
		return intValue, nil
	} else {
		fmt.Println("Indicar un número que este entre las opciones.")
		return 0, errors.New("indicar un número que este entre las opciones")
	}
}

func chooseOption(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina], accion int) {
	acciones := [5]string{"listar", "buscar", "agregar", "borrar", "modificar"}
	// Muestra opciones a elegir
	fmt.Println("Que quiere", acciones[accion], ":")
	fmt.Println("1) Ejercicios")
	fmt.Println("2) Rutinas")
	fmt.Println("")
	fmt.Println("0) Vuelve al menú principal")

	opcionElegida := getInput(0, 2)
	if opcionElegida != -1 {
		switch acciones[accion] {
		case "listar":
			listarElementoSeleccionado(lle, llr, opcionElegida)
		case "buscar":
			buscarElementoSeleccionado(lle, llr, opcionElegida)
		case "agregar":
			agregarElementoSeleccionado(lle, llr, opcionElegida)
		case "borrar":
			borrarElementoSeleccionado(lle, llr, opcionElegida)
		case "modificar":
			modificarElementoSeleccionado(lle, llr, opcionElegida)
		}
	}
}

func listarElementoSeleccionado(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina], opcion int) {
	switch opcion {
	case 1:
		ListarEjercicios(lle)
	case 2:
		ListarRutinas(llr)
	}
}

func buscarElementoSeleccionado(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina], opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Ingrese el ID del ejercicio que quiere ver.")
		inputID := getInput(1, lle.Size()+1)
		MostrarEjercicio(lle, strconv.Itoa(inputID))
	case 2:
		fmt.Println("Ingrese el ID de la rutina que quiere ver.")
		inputID := getInput(1, llr.Size()+1)
		MostrarRutina(llr, strconv.Itoa(inputID))
	}
}

func agregarElementoSeleccionado(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina], opcion int) {
	switch opcion {
	case 1:
		agregarEjercicioCLI(lle)
	case 2:
		agregarRutinaCLI(llr, lle)
	}
}

func agregarEjercicioCLI(lle *list.LinkedList[Ejercicio]) {
	fmt.Println("Complete los campos del nuevo ejercicio a agregar.")
	fmt.Println("Nombre del ejercicio: ")
	nombreEj := getInputString()
	fmt.Println("Descripción del ejercicio: ")
	descripcionEj := getInputString()
	fmt.Println("Tiempo en minutos del ejercicio: ")
	tiempoEj := getInputString()
	fmt.Println("Calories del ejercicio: ")
	caloriasEj := getInputString()
	fmt.Println("Específique el tipo de ejercicio: ")
	tipoEj := getInputString()
	fmt.Println("Específique el grupo múscilar del ejercicio: ")
	grupoMuscularEj := getInputString()
	fmt.Println("Específique los puntos que vale este ejercicio: ")
	puntosEJ := getInputString()
	fmt.Println("Específique la dificultad del ejercicio (entre 1 y 3)")
	dificultadEj := getInputString()

	ejercicoNuevo := NewEjercicio(nombreEj, descripcionEj, tiempoEj, caloriasEj, tipoEj, grupoMuscularEj, puntosEJ, dificultadEj)
	AgregarEjercicio(lle, ejercicoNuevo)

}

func agregarRutinaCLI(llr *list.LinkedList[Rutina], lle *list.LinkedList[Ejercicio]) {
	fmt.Println("Complete los campos de la nueva rutina a agregar.")
	fmt.Println("Nombre de la rutina: ")
	nombreRut := getInputString()
	fmt.Println("IDs de los ejercicios asociados. Ej: 1,3,5")
	ejerciciosRut := getInputString()

	rutinaNueva, err := NewRutina(nombreRut, ejerciciosRut, lle)
	if err != nil {
		fmt.Printf("No se pudo crear correctamente la rutina. Ingrese los parámetros específicados.")
	}
	AgregarRutina(llr, rutinaNueva)
}

func borrarElementoSeleccionado(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina], opcion int) {
	switch opcion {
	case 1:
		// TODO: Agregar una confirmación para estar seguro de borrar el elemento deseado
		fmt.Println("Ingrese el ID del ejercicio a borrar.")
		inputID := getInput(1, lle.Size()+1)
		BorrarEjercicio(lle, strconv.Itoa(inputID), llr)
	case 2:
		fmt.Println("Ingrese el ID de la rutina a borrar.")
		inputID := getInput(1, llr.Size()+1)
		BorrarRutina(llr, strconv.Itoa(inputID))
	}
}

func modificarElementoSeleccionado(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina], opcion int) {
	switch opcion {
	case 1:
		modificarEjercicioCLI(lle)

	case 2:
		modificarRutinaCLI(lle, llr)
		fmt.Println("Acá tendría que modifiar una rutina")
	}
}

func modificarEjercicioCLI(lle *list.LinkedList[Ejercicio]) {
	fmt.Println("Ingrese el ID del ejercicio a modificar.")
	inputID := getInput(1, lle.Size()+1)

	fmt.Println("Ingrese el campo que quiere modificar")
	fmt.Println("1) Nombre")
	fmt.Println("2) Tiempo")
	fmt.Println("3) Calorías")
	fmt.Println("4) Tipo")
	fmt.Println("5) Grupo Muscular")
	fmt.Println("6) Puntos")
	fmt.Println("7) Dificultad")
	inputCampoModificar := getInput(1, 7)
	campoAModificar := eleccionModificarACampo(inputCampoModificar)

	campoModificado := getInputString()

	ModificarEjercicio(lle, strconv.Itoa(inputID), campoAModificar, campoModificado)

}

func eleccionModificarACampo(numSeleccion int) string {
	switch numSeleccion {
	case 1:
		fmt.Println("Ingrese el nuevo nombre")
		return "Nombre"
	case 2:
		fmt.Println("Ingrese el tiempo del ejercicio")
		return "Tiempo"
	case 3:
		fmt.Println("Ingrese las calorías del ejercicio")
		return "Calorias"
	case 4:
		fmt.Println("Ingrese el tipo del ejercicio")
		return "Tipo"
	case 5:
		fmt.Println("Ingrese el grupo múscular del ejercicio")
		return "GrupoMuscular"
	case 6:
		fmt.Println("Ingrese los puntos del ejercicio")
		return "Puntos"
	case 7:
		fmt.Println("Ingrese la dificultad del ejercicio del 1 al 3. Donde 1 Bajo, 2 Medio y 3 Alto")
		return "Dificultad"
	}
	return ""
}

func modificarRutinaCLI(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina]) {
	fmt.Println("Ingrese el ID de la rutina a modificar.")
	inputID := getInput(1, llr.Size()+1)

	fmt.Println("Ingrese el campo que quiere modificar")
	fmt.Println("1) Nombre")
	fmt.Println("2) Ejercicios")
	inputCampoModificar := getInput(1, 2)
	var campoAModificar string
	if inputCampoModificar == 1 {
		fmt.Println("Ingrese el nuevo nombre")
		campoAModificar = "Nombre"
	}
	if inputCampoModificar == 2 {
		fmt.Println("Ingrese los nuevos ejercicios separados por coma. Ej: (1,2,5,7)")
		campoAModificar = "Ejercicios"
	}

	campoModificado := getInputString()
	ModificarRutina(llr, lle, strconv.Itoa(inputID), campoAModificar, campoModificado)

}

func chooseRoutineGeneration(lle *list.LinkedList[Ejercicio], llr *list.LinkedList[Rutina]) {
	fmt.Println("Elija que tipo de rutina quiere generar")
	fmt.Println("1) Máxima cantidad de ejercicios en un tiempo determinado filatrando por tipo y dificultad")
	fmt.Println("2) Mínima duración de los ejercicios para un valor de calorías quemadas objetivo")
	fmt.Println("3) Máximos puntos en una categoría para una duración fija")
	fmt.Println("")
	fmt.Println("0) Vuelve al menú principal")

	opcionElegida := getInput(0, 3)
	switch opcionElegida {
	case 1:
		fmt.Println("Ingerse el nombre de la rutina a generar.")
		nombreRutinaGenerada := getInputString()

		fmt.Println("Ingerse la duración total de la rutina a generar en minutos (mayor a 2 minutos)")
		duracionTotalObjetivo := getInput(2, 1000)

		fmt.Println("Ingerse el tipo de ejercicios a optimizar.")
		tipoEjercicio := getInputString()

		fmt.Println("Ingerse la dificultad de ejercicios a optimizar. (Valor de 1 a 3)")
		dificultadInt := getInput(1, 3)
		dificultadEjercicio := strconv.Itoa(dificultadInt)

		rutinaGenerada, err := NewRutinaPorTipoyDificultad(nombreRutinaGenerada, duracionTotalObjetivo, tipoEjercicio, dificultadEjercicio, lle)
		if err != nil {
			fmt.Println("Error al generar la rutina. Por favor ingrese nuevamente los campos.")
		}
		AgregarRutina(llr, rutinaGenerada)

	case 2:
		fmt.Println("Ingerse el nombre de la rutina a generar.")
		nombreRutinaGenerada := getInputString()

		fmt.Println("Ingrese el valor objetivo de calorías a quemar")
		caloriasObjetivo := getInput(200, 100000)

		rutinaGenerada, err := NewRutinaPorCalorias(nombreRutinaGenerada, caloriasObjetivo, lle)
		if err != nil {
			fmt.Println("Error al generar la rutina. Por favor ingrese nuevamente los campos.")
		}
		AgregarRutina(llr, rutinaGenerada)
	case 3:
		fmt.Println("Ingerse el nombre de la rutina a generar.")
		nombreRutinaGenerada := getInputString()

		fmt.Println("Ingrese el tipo que quiera maximizar")
		tipoObjetivo := getInputString()

		fmt.Println("Ingerse la duración total de la rutina a generar en minutos (mayor a 2 minutos)")
		duracionTotalObjetivo := getInput(2, 1000)

		rutinaGenerada, err := NewRutinaPorTipoPtsYDuracion(nombreRutinaGenerada, tipoObjetivo, duracionTotalObjetivo, lle)
		if err != nil {
			fmt.Println("Error al generar la rutina. Por favor ingrese nuevamente los campos.")
		}
		AgregarRutina(llr, rutinaGenerada)
	}
}
