package code

import (
	"fmt"
	"strconv"
	"strings"

	"errors"

	"github.com/untref-ayp2/data-structures/list"
)

type Ejercicio struct {
	Id            string `csv:"id"`
	Nombre        string `csv:"nombre"`
	Descripcion   string `csv:"descripcion"`
	Tiempo        string `csv:"tiempo"`
	Calorias      string `csv:"calorias"`
	Tipo          string `csv:"tipo"`
	GrupoMuscular string `csv:"grupoMuscular"`
	Pts           string `csv:"pts"`
	Dificultad    string `csv:"dificultad"`
}

func NewEjercicio(nombre, desc, tiempo, cal, tipo, grupoMusc, pts, dificultad string) *Ejercicio {
	return &Ejercicio{
		Nombre:        nombre,
		Descripcion:   desc,
		Tiempo:        tiempo,
		Calorias:      cal,
		Tipo:          tipo,
		GrupoMuscular: grupoMusc,
		Pts:           pts,
		Dificultad:    dificultad}
}

func (ejer *Ejercicio) mostrarEjercicio() {
	fmt.Println("Id: ", ejer.Id)
	fmt.Println("Nombre: ", ejer.Nombre)
	fmt.Println("Descripción: ", ejer.Descripcion)
	fmt.Println("Tiempo: ", ejer.Tiempo)
	fmt.Println("Calorías: ", ejer.Calorias)
	fmt.Println("Grupo Muscular: ", ejer.GrupoMuscular)
	fmt.Println("Tipo: ", ejer.Tipo)
	fmt.Println("Puntos del ejercicio: ", ejer.Pts)
	fmt.Println("Dificultad: ", numberLevelToString(ejer.Dificultad))
	fmt.Println("-----------------------------")
}

func (ejer *Ejercicio) ModificarDatos(field string, value string) error {
	switch field {
	case "Nombre":
		ejer.Nombre = value
	case "Tiempo":
		ejer.Tiempo = value
	case "Calorias":
		ejer.Calorias = value
	case "Tipo":
		ejer.Tipo = value
	case "GrupoMuscular":
		ejer.GrupoMuscular = value
	case "Puntos":
		ejer.Pts = value
	case "Dificultad":
		ejer.Dificultad = value
	default:
		return errors.New("campo invalido")
		// cuando no coincide ningun dato deberia dar un error?
	}
	return nil
}

func numberLevelToString(levelDif string) string {
	if levelDif == "1" {
		return "Bajo"
	}

	if levelDif == "2" {
		return "Medio"
	}

	if levelDif == "3" {
		return "Alto"
	}
	return "Error en la dificultad"
}

func NewLinkedListEjercicios(listaEj []Ejercicio) *list.LinkedList[Ejercicio] {
	lle := list.NewLinkedList[Ejercicio]()
	for _, val := range listaEj {
		lle.Append(val)
	}
	return lle
}

func (ejer *Ejercicio) IDInt() int {
	intID, _ := strconv.Atoi(ejer.Id) //traduce entero a string
	return intID
}

func (ejer *Ejercicio) TiempoInt() int {
	intTiempo, _ := strconv.Atoi(ejer.Tiempo)
	return intTiempo
}

func (ejer *Ejercicio) CaloriasInt() int {
	intCal, _ := strconv.Atoi(ejer.Calorias)
	return intCal
}

func (ejer *Ejercicio) PuntosInt() int {
	intPts, _ := strconv.Atoi(ejer.Pts)
	return intPts
}

func (ejer *Ejercicio) ArrayTipos() []string {
	return strings.Split(ejer.Tipo, ",")
}

func (ejer *Ejercicio) EsDeTipo(tipo string) bool {
	for _, tipoEjer := range ejer.ArrayTipos() {
		if tipoEjer == tipo {
			return true
		}
	}
	return false
}

func (ejer *Ejercicio) PtsDeTipo(tipo string) int {
	arrayPts := strings.Split(ejer.Pts, ",")
	for i, tipoEjer := range ejer.ArrayTipos() {
		if tipoEjer == tipo {
			ptsInt, _ := strconv.Atoi(arrayPts[i])
			return ptsInt
		}
	}
	return 0
}
