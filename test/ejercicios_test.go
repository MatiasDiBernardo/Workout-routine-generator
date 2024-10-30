package test

import (
	"TP-2024-malo/code"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscarEjercicioExistente(t *testing.T) {
	lle := code.GetEjercicios()

	ejercicios := code.ListarEjercicios(lle)

	ejercicioExistente := ejercicios[len(ejercicios)-1]
	ejercicioExistenteId := ejercicioExistente.Id

	ejercicio, _ := code.MostrarEjercicio(lle, ejercicioExistenteId)

	assert.Equal(t, ejercicio.Id, ejercicioExistenteId)
}

func TestBuscarEjercicioInexistente(t *testing.T) {
	lle := code.GetEjercicios()

	_, err := code.MostrarEjercicio(lle, "124")

	assert.Error(t, err)
}

func TestAgregarYBorrarEjercicio(t *testing.T) {
	/*
		para no usar la lle ni el size ni nada se me ocurre una func tipo
		"iniciar sesion"  o algo asi y capaz un struct sesion que tenga los ejercicios
		y rutinas en dos variables, asi nos ahorramos llamar al lle o llr por parametro desde el test
		capaz sirva para el CLI tambien [OA]
	*/

	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	nuevoEjercicio := code.NewEjercicio(
		"Nuevo Ej",
		"Desc",
		"10",
		"600",
		"Brazos",
		"Brazos",
		"5",
		"1",
	)

	//Agregar
	code.AgregarEjercicio(lle, nuevoEjercicio)
	ejercicios := code.ListarEjercicios(lle)

	ultimoEjercicio := ejercicios[len(ejercicios)-1]
	ultimoEjercicioID := ultimoEjercicio.Id

	assert.Equal(t, ultimoEjercicio.Nombre, "Nuevo Ej")
	//se pueden sumar asserts de mas variables del ejercicio para que sea mas completo

	//Borrar
	code.BorrarEjercicio(lle, ultimoEjercicioID, llr)
	_, err := code.MostrarEjercicio(lle, ultimoEjercicioID)
	assert.Error(t, err)
}

func TestBorrarEjercicioDeUnaRutina(t *testing.T) {

	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)

	//Agregar
	rutinas := code.ListarRutinas(llr)

	//Borrar ejercicio de la ultima Rutina
	ejerciciosDeLaUltimaRutina := rutinas[len(rutinas)-1].Ejercicios
	idEjercicio := strings.Split(ejerciciosDeLaUltimaRutina, ",")[0]

	fmt.Println("Borrar ejercicio -> ", idEjercicio)

	err := code.BorrarEjercicio(lle, idEjercicio, llr)

	assert.Error(t, err)
}

// con los cambios no haria falta este Test, se puede recorrer el array que trae ListarEjercicios e imprimir alguna variable de los mismos [OA]
func TestImprimirEjercicios(t *testing.T) {
	lle := code.GetEjercicios()
	code.ListarEjercicios(lle)
}

//Estos falta adecuarlos a la nueva lógica [OA]

func TestModificarNombreEjercicio(t *testing.T) {
	lle := code.GetEjercicios()

	nombreViejo := lle.Tail().Data().Nombre
	idEjercicioModificado := lle.Tail().Data().Id
	code.ModificarEjercicio(lle, idEjercicioModificado, "Nombre", "Nuevo nombre")
	nombreNuevo := lle.Tail().Data().Nombre

	assert.NotEqual(t, nombreViejo, nombreNuevo)
	code.ModificarEjercicio(lle, idEjercicioModificado, "Nombre", nombreViejo)
}

func TestModificarActualizaCSV(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	nuevoEjercicio := code.NewEjercicio(
		"Nuevo Ej",
		"Desc",
		"10",
		"600",
		"Brazos",
		"Brazos",
		"5",
		"1",
	)
	code.AgregarEjercicio(lle, nuevoEjercicio)

	// Vuelve a cargar los datos del csv
	lle2 := code.GetEjercicios()

	assert.Equal(t, lle.Size(), lle2.Size())

	code.BorrarEjercicio(lle, lle.Tail().Data().Id, llr)
	assert.NotEqual(t, lle.Size(), lle2.Size())
}
