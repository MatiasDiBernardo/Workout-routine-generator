package test

import (
	"TP-2024-malo/code"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgregarYBorrarRutina(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	nuevaRutina, _ := code.NewRutina("Rutina básica", "1,3", lle)

	//Agregar
	code.AgregarRutina(llr, nuevaRutina)
	rutinas := code.ListarRutinas(llr)

	ultimaRutina := rutinas[len(rutinas)-1]
	ultimaRutinaID := ultimaRutina.Id

	assert.Equal(t, ultimaRutina.Nombre, "Rutina básica")
	//se pueden sumar asserts de mas variables del ejercicio para que sea mas completo

	//Borrar
	code.BorrarRutina(llr, ultimaRutinaID)
	_, err := code.MostrarRutina(llr, ultimaRutinaID)
	assert.Error(t, err)
}

func TestAgregarRutinaConEjercicioInexistente(t *testing.T) {
	lle := code.GetEjercicios()
	_, err := code.NewRutina("Rutina básica", "1,454", lle)

	assert.Error(t, err)
}

func TestImprimirRutina(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	code.ListarRutinas(llr)
}

func TestModificarNombreRutina(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	rutinas := code.ListarRutinas(llr)

	rutinaEjemplo := rutinas[len(rutinas)-1]
	nombreViejo := rutinaEjemplo.Nombre

	code.ModificarRutina(llr, lle, rutinaEjemplo.Id, "Nombre", "Nuevo nombre rutina")

	rutinaModificada, _ := code.MostrarRutina(llr, rutinaEjemplo.Id)

	assert.NotEqual(t, nombreViejo, rutinaModificada.Nombre)
	assert.Equal(t, "Nuevo nombre rutina", rutinaModificada.Nombre)

	// Resetea el nombre para mantener consistencia
	code.ModificarRutina(llr, lle, rutinaEjemplo.Id, "Nombre", nombreViejo)
}

func TestModificarEjerciciosValidosRutina(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	rutinas := code.ListarRutinas(llr)

	rutinaEjemplo := rutinas[len(rutinas)-1]

	ejsViejos := rutinaEjemplo.Ejercicios

	ejsTarget := "2,3"
	err := code.ModificarRutina(llr, lle, rutinaEjemplo.Id, "Ejercicios", ejsTarget)

	rutinaModificada, _ := code.MostrarRutina(llr, rutinaEjemplo.Id)

	assert.Nil(t, err)
	assert.NotEqual(t, ejsViejos, rutinaModificada.Ejercicios)
	assert.Equal(t, ejsTarget, rutinaModificada.Ejercicios)

	// Resetea los ejercicios para mantener consistencia
	code.ModificarRutina(llr, lle, rutinaModificada.Id, "Ejercicios", "1,2")
}

func TestModificarEjerciciosInvalidosRutina(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	rutinas := code.ListarRutinas(llr)

	rutinaEjemplo := rutinas[len(rutinas)-1]

	ejsViejos := rutinaEjemplo.Ejercicios

	ejsTarget := "2,100000"
	err := code.ModificarRutina(llr, lle, rutinaEjemplo.Id, "Ejercicios", ejsTarget)

	rutinaNoModificada, _ := code.MostrarRutina(llr, rutinaEjemplo.Id)

	assert.Equal(t, ejsViejos, rutinaNoModificada.Ejercicios)

	assert.Error(t, err)

}

func TestAgregarEjercicioARutina(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)

	rutinas := code.ListarRutinas(llr)

	rutinaEjemplo := rutinas[len(rutinas)-1]

	ejsViejos := rutinaEjemplo.Ejercicios

	ejAgregado := "1"

	code.AddRemoveEjercicioToRutina(llr, lle, rutinaEjemplo.Id, ejAgregado, "A")

	rutinaNoModificada, _ := code.MostrarRutina(llr, rutinaEjemplo.Id)

	assert.NotEqual(t, ejsViejos, rutinaNoModificada.Ejercicios)
	assert.Contains(t, rutinaNoModificada.Ejercicios, ejAgregado)

	// Remueve el ejercicio agregado para mantener consistencia
	code.AddRemoveEjercicioToRutina(llr, lle, rutinaEjemplo.Id, ejAgregado, "R")

	rutinaNoModificada, _ = code.MostrarRutina(llr, rutinaEjemplo.Id)

	assert.Equal(t, ejsViejos, rutinaNoModificada.Ejercicios)
}

func TestModificarCambiaCSV(t *testing.T) {
	lle := code.GetEjercicios()
	llr := code.GetRutinas(lle)
	nuevaRutina, _ := code.NewRutina("Rutina básica", "1,3", lle)

	if nuevaRutina != nil {
		code.AgregarRutina(llr, nuevaRutina)
	}

	// Vuelve a cargar los datos del CSV
	llr2 := code.GetRutinas(lle)

	rutinas2 := code.ListarRutinas(llr2)
	ultimaAgregada := rutinas2[len(rutinas2)-1]

	assert.Equal(t, nuevaRutina.Nombre, ultimaAgregada.Nombre)

	// Remueve el ejercicio agregado para mantener consistencia
	code.BorrarRutina(llr, ultimaAgregada.Id)

	_, err := code.MostrarRutina(llr, ultimaAgregada.Id)

	assert.Error(t, err)
}

func TestNewRutinaPorTipoyDificultad(t *testing.T) {
	lle := code.GetEjercicios()
	rutina, err := code.NewRutinaPorTipoyDificultad("Rutina Filtrada", 30, "Cardio", "2", lle)
	assert.Nil(t, err)
	assert.NotNil(t, rutina)
	assert.Equal(t, "Rutina Filtrada", rutina.Nombre)

	tiempoTotal := 0
	for e := rutina.ListEjercicios.Head(); e != nil; e = e.Next() {
		ejercicio := e.Data()
		assert.Equal(t, "Cardio", ejercicio.Tipo)
		assert.Equal(t, "2", ejercicio.Dificultad)
		tiempoTotal += ejercicio.TiempoInt()
	}
	assert.LessOrEqual(t, tiempoTotal, 30)
}

func TestNewRutinaPorCalorias(t *testing.T) {
	lle := code.GetEjercicios()
	rutina, err := code.NewRutinaPorCalorias("Rutina Calorías", 5000, lle)
	assert.Nil(t, err)
	assert.NotNil(t, rutina)
	assert.Equal(t, "Rutina Calorías", rutina.Nombre)
	caloriasAcumuladas := 0
	tiempoTotal := 0
	for e := rutina.ListEjercicios.Head(); e != nil; e = e.Next() {
		ejercicio := e.Data()
		caloriasAcumuladas += ejercicio.CaloriasInt()
		tiempoTotal += ejercicio.TiempoInt()
	}
	assert.GreaterOrEqual(t, caloriasAcumuladas, 500)
}

func TestNewRutinaPorPtsYTipoSoloPrioritarios(t *testing.T) {
	lle := code.GetEjercicios()
	rutina, err := code.NewRutinaPorTipoPtsYDuracion("Rutina Tipo y Pts", "Velocidad", 20, lle)
	assert.Nil(t, err)
	assert.NotNil(t, rutina)
	rutina.MostrarRutina()
	rutina.EjerciciosDeRutina()
	assert.Equal(t, "Rutina Tipo y Pts", rutina.Nombre)
}

func TestNewRutinaPorPtsYTipoPrioritariosYNoPrioritarios(t *testing.T) {
	lle := code.GetEjercicios()
	rutina, err := code.NewRutinaPorTipoPtsYDuracion("Rutina Tipo y Pts 2", "Velocidad", 25, lle)
	assert.Nil(t, err)
	assert.NotNil(t, rutina)
	rutina.MostrarRutina()
	rutina.EjerciciosDeRutina()
	assert.Equal(t, "Rutina Tipo y Pts 2", rutina.Nombre)
}
