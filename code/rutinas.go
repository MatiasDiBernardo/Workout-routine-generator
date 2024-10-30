package code

import (
	"errors"
	"fmt"
	"strings"

	"github.com/untref-ayp2/data-structures/list"
)

// Crear la rutina como struct (sería un nodo de la link list de rutinas)
type Rutina struct {
	Id             string                      `csv:"id"`
	Nombre         string                      `csv:"nombre"`
	Ejercicios     string                      `csv:"ejercicios"`
	ListEjercicios *list.LinkedList[Ejercicio] // Link list de los ejercicios de esa rutina
}

// Retorna una linkedList de Ejercicio a partir de otro ll de Ejercicio "padre" y un string con los ids de los mismos
func NewLinkedListEjerciciosById(ejerciciosDisponibles *list.LinkedList[Ejercicio], idsEjercicios string) (*list.LinkedList[Ejercicio], error) {

	arrayEjercicios := strings.Split(idsEjercicios, ",")

	listEjercicios := list.NewLinkedList[Ejercicio]()
	for _, id := range arrayEjercicios {
		ejEncontrado := findEjercicioById(id, ejerciciosDisponibles)
		// Devuelve error si el ejercicio no fue encontrado
		if ejEncontrado == nil {
			return nil, errors.New(fmt.Sprint("El ejercicion con id", id, "no existe."))
		}
		listEjercicios.Append(ejEncontrado.Data())
	}

	return listEjercicios, nil
}

// Inicializa una rutina
func NewRutina(nombre, ejercicios string, ejerciciosDisponibles *list.LinkedList[Ejercicio]) (*Rutina, error) {
	// Habría que validar si le paso un ejercicio que no esta en la lista de ejercicios
	listEjercicios, err := NewLinkedListEjerciciosById(ejerciciosDisponibles, ejercicios)
	if listEjercicios == nil {
		return nil, err
	}
	return &Rutina{
		Nombre:         nombre,
		Ejercicios:     ejercicios,
		ListEjercicios: listEjercicios,
	}, nil
}

// Imprime las caracteristicas de la Rutina por consola
func (rut *Rutina) MostrarRutina() {
	fmt.Println("Id: ", rut.Id)
	fmt.Println("Nombre: ", rut.Nombre)
	fmt.Println("Calorias: ", rut.Calorias())
	fmt.Println("Duracion: ", rut.Duracion())
	fmt.Println("Dificultad: ", rut.Dificultad())
	fmt.Println("Tipo: ", rut.TipoDeEjercicios())
	fmt.Println("Ejercicios: ")
	fmt.Println("-------------------------")
	for node := rut.ListEjercicios.Head(); node != nil; node = node.Next() {
		ejer := node.Data()
		fmt.Println("Id ", ejer.Id, ": ", ejer.Nombre, " - ", ejer.Descripcion, "[cal.:", ejer.Calorias, "][t.:", ejer.Tiempo, "]")
	}
	fmt.Println("-------------------------")
}

// Modifica los datos de la Rutina informando el campo a modificar y su nuevo valor
func (rut *Rutina) ModificarDatos(field string, value string, lle *list.LinkedList[Ejercicio]) error {
	switch field {
	case "Nombre":
		rut.Nombre = value
	case "Ejercicios":
		// En ejercicios actualizo el string y la lle asociada
		lleActualizada, err := NewLinkedListEjerciciosById(lle, value)
		if lleActualizada == nil {
			return err
		}
		rut.Ejercicios = value
		rut.ListEjercicios = lleActualizada
	default:
		return errors.New("campo invalido")
		// cuando no coincide ningun dato deberia dar un error?
	}
	return nil
}

// Anade un ejercicio a la Rutina
func (rut *Rutina) AddEjer(nuevoEjer Ejercicio) {
	rut.ListEjercicios.Append(nuevoEjer)
	rut.Ejercicios = fmt.Sprint(rut.Ejercicios, ",", nuevoEjer.Id)
}

// Remueve un ejercicio de la Rutina
func (rut *Rutina) RemoveEjer(ejerABorrar Ejercicio) {
	idEliminado := ejerABorrar.Id
	//Saca al ejercicio de la LinkedList del struct
	rut.ListEjercicios.Remove(ejerABorrar)

	//Saca el id del ejercicio en el campo de ids del struct
	idsEjercicios := strings.Split(rut.Ejercicios, ",")
	rut.Ejercicios = ""

	for _, id := range idsEjercicios {
		if id != idEliminado {
			rut.Ejercicios = fmt.Sprint(rut.Ejercicios, id, ",")
		}
	}

	if len(rut.Ejercicios) > 0 && string(rut.Ejercicios[len(rut.Ejercicios)-1]) == "," {
		rut.Ejercicios = rut.Ejercicios[:len(rut.Ejercicios)-1]
	}
}

// Crea una LinkedList de Rutina a partir de un array de Rutina y un LinkedList de ejercicios disponibles
func NewLinkedListRutinas(listaRut []Rutina, ejercicios *list.LinkedList[Ejercicio]) *list.LinkedList[Rutina] {
	llr := list.NewLinkedList[Rutina]()
	for _, rut := range listaRut {
		// Crea la LinkedList de ejercicios a partir de los informados en el campo Ejercicios (ids) de cada rutina
		rut.ListEjercicios, _ = NewLinkedListEjerciciosById(ejercicios, rut.Ejercicios)
		llr.Append(rut)
	}
	return llr
}

// Calcula calorias quemadas en la Rutina a partir de los ejercicios
func (r *Rutina) Calorias() int {
	sumaCalorias := 0
	for nodo := r.ListEjercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		calorias := ejercicio.CaloriasInt()
		sumaCalorias += calorias
	}
	return sumaCalorias
}

// Calcula duracion total de la Rutina a partir de los ejercicios
func (r *Rutina) Duracion() int {
	sumaTiempo := 0
	for nodo := r.ListEjercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		Tiempo := ejercicio.TiempoInt()
		sumaTiempo += Tiempo
	}
	return sumaTiempo
}

// Retorna la dificultad mas frecuente entre los ejercicios de la Rutina
func (r *Rutina) Dificultad() string {
	sumaDificultades := 0
	numEjercicios := 0
	for nodo := r.ListEjercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		dificultad := ejercicio.Dificultad
		switch dificultad {
		case "Bajo":
			sumaDificultades += 1
		case "Medio":
			sumaDificultades += 2
		case "Alto":
			sumaDificultades += 3
		default:
			sumaDificultades += 2
		}
		numEjercicios++
	}
	// Calcular el promedio de dificultad
	promedioDificultad := float64(sumaDificultades) / float64(numEjercicios)

	//return numberLevelToString(strconv.Itoa(int(promedioDificultad)))

	// Lo modifico para seguir la consigna:
	// "El tipo de ejercicios y la dificultad de la rutina serán los más frecuentes entre los ejercicios que la componen."
	if promedioDificultad < 1.5 {
		return "Bajo"
	}

	if promedioDificultad < 2.5 {
		return "Medio"
	}

	return "Alto"
}

// Retorna el tipo mas frecuente entre los ejercicios de la Rutina
func (r *Rutina) TipoDeEjercicios() string {

	// Lo modifico para seguir la consigna:
	// "El tipo de ejercicios y la dificultad de la rutina serán los más frecuentes entre los ejercicios que la componen."

	// Recorro la ller y stackeo el tipo de ejercicios
	tiposMap := make(map[string]int)
	for nodo := r.ListEjercicios.Head(); nodo != nil; nodo = nodo.Next() {
		for _, tipo := range strings.Split(nodo.Data().Tipo, ",") {

			// Agregar el tipo de ejercicio al mapa si aún no está presente
			if tiposMap[tipo] == 0 {
				tiposMap[tipo] = 1
			} else {
				cant := tiposMap[tipo]
				tiposMap[tipo] = cant + 1
			}
		}
	}

	max := 0
	tipoFinal := ""

	for tipo, cant := range tiposMap {
		if cant > max {
			max = cant
			tipoFinal = tipo
		}
	}

	return tipoFinal

	// Convertir las claves del mapa en un array de strings
	//tipos := make([]string, 0, len(tiposMap))
	//for tipo := range tiposMap {
	//	tipos = append(tipos, tipo)
	//}
	//return tipos

}

func (rut *Rutina) EjerciciosDeRutina() []Ejercicio {
	return ListarEjercicios(rut.ListEjercicios)
}

// generacion automagica 1
func NewRutinaPorTipoyDificultad(nombre string, duracionTotal int, tipo, dificultad string, lle *list.LinkedList[Ejercicio]) (*Rutina, error) {
	/*
		ejerciciosFiltrados := list.NewLinkedList[Ejercicio]()

		for e := lle.Head(); e != nil; e = e.Next() {
			ejercicio := e.Data()
			if ejercicio.Tipo == tipo && ejercicio.Dificultad == dificultad {
				ejerciciosFiltrados.Append(ejercicio)
			}
		}
		if ejerciciosFiltrados.Size() == 0 {
			return nil, errors.New("no hay ejercicios disponibles que cumplan con los criterios")
		}
		//ordenar la lista de ejercicios filtrados de menor a mayor tiempo (use un metodo de burbujeo no muy optimo)
		ejerciciosOrdenados := make([]Ejercicio, 0, ejerciciosFiltrados.Size())
		for e := ejerciciosFiltrados.Head(); e != nil; e = e.Next() {
			ejerciciosOrdenados = append(ejerciciosOrdenados, e.Data())
		}
	*/

	//en lugar de hacer un linkeslist con los filtrados y despues un array con ese linked list se puede hacer el array directamente

	// Se filtran los ejercicios que cumplan con los criterios
	ejerciciosFiltrados := []Ejercicio{}

	for e := lle.Head(); e != nil; e = e.Next() {
		ejercicio := e.Data()
		if ejercicio.Dificultad == dificultad {
			for _, tipoEj := range ejercicio.ArrayTipos() {
				if tipoEj == tipo && ejercicio.Dificultad == dificultad {
					ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
				}
			}
		}
	}
	if len(ejerciciosFiltrados) == 0 {
		return nil, errors.New("no hay ejercicios disponibles que cumplan con los criterios")
	}

	// Se ordenan los ejercicios filtrados en base a tiempo (de min a max)
	ejerciciosOrdenados := ejerciciosFiltrados

	n := len(ejerciciosOrdenados)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ejerciciosOrdenados[j].TiempoInt() > ejerciciosOrdenados[j+1].TiempoInt() {
				ejerciciosOrdenados[j], ejerciciosOrdenados[j+1] = ejerciciosOrdenados[j+1], ejerciciosOrdenados[j]
			}
		}
	}

	//Se añaden todos los ejercicios que se puedan dentro del tiempo establecido a un linkedList
	ejerciciosSeleccionados := list.NewLinkedList[Ejercicio]()
	idsEjercicios := []string{} //make([]string, 0, ejerciciosSeleccionados.Size())
	tiempoTotal := 0

	for _, ejercicio := range ejerciciosOrdenados {
		if tiempoTotal+ejercicio.TiempoInt() <= duracionTotal {
			ejerciciosSeleccionados.Append(ejercicio)
			tiempoTotal += ejercicio.TiempoInt()
			//if ejerciciosSeleccionados.Size() > 0 { // no hace falta - siempre se cumple
			idsEjercicios = append(idsEjercicios, ejercicio.Id)
			//}
		}
	}

	//se retorna la rutina con los ids seleccionados (no se tiene en cuenta el error porque se sabe que los ejercicios existen)
	idsEjerciciosStr := strings.Join(idsEjercicios, ",")
	rutina, _ := NewRutina(nombre, idsEjerciciosStr, ejerciciosSeleccionados)
	return rutina, nil
}

// generacion automagica 2
func NewRutinaPorCalorias(nombre string, caloriasTotales int, lle *list.LinkedList[Ejercicio]) (*Rutina, error) {
	// Se pasan TODOS los ejercicios a un array
	ejerciciosOrdenados := make([]Ejercicio, 0, lle.Size())
	for e := lle.Head(); e != nil; e = e.Next() {
		ejerciciosOrdenados = append(ejerciciosOrdenados, e.Data())
	}

	// Se ordena ese array por caloria/tiempo (de mayor a menor)
	n := len(ejerciciosOrdenados)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			caloriasPorTiempo := float64(ejerciciosOrdenados[j].CaloriasInt()) / float64(ejerciciosOrdenados[j].TiempoInt())
			elSiguiente := float64(ejerciciosOrdenados[j+1].CaloriasInt()) / float64(ejerciciosOrdenados[j+1].TiempoInt())
			if caloriasPorTiempo < elSiguiente {
				ejerciciosOrdenados[j], ejerciciosOrdenados[j+1] = ejerciciosOrdenados[j+1], ejerciciosOrdenados[j]
			}
		}
	}

	//Se añaden los ejercicios hasta llegar a las calorias requeridas - o aproximarse a ellas
	ejerciciosSeleccionados := list.NewLinkedList[Ejercicio]()
	idsEjercicios := []string{}
	caloriasAcumuladas := 0
	//tiempoTotal := 0 // no hace falta

	for _, ejercicio := range ejerciciosOrdenados {
		if caloriasAcumuladas+ejercicio.CaloriasInt() <= caloriasTotales {
			ejerciciosSeleccionados.Append(ejercicio)
			caloriasAcumuladas += ejercicio.CaloriasInt()
			//tiempoTotal += ejercicio.TiempoInt()
			idsEjercicios = append(idsEjercicios, ejercicio.Id)
		}
	}

	// Hay que mirar de nuevo este chequeo porque sino entra en casos donde no corresopnde
	// if caloriasAcumuladas < caloriasTotales { // preguntar si esto deberia dar error o darlo como valido
	// 	return nil, errors.New("no hay suficientes ejercicios para alcanzar las calorías deseadas")
	// }

	//se retorna la rutina con los ids seleccionados (no se tiene en cuenta el error porque se sabe que los ejercicios existen)
	idsEjerciciosStr := strings.Join(idsEjercicios, ",")
	rutina, _ := NewRutina(nombre, idsEjerciciosStr, ejerciciosSeleccionados)
	return rutina, nil
}

// generacion automagica 3
func NewRutinaPorTipoPtsYDuracion(nombre string, tipoMaximizado string, duracionTotal int, lle *list.LinkedList[Ejercicio]) (*Rutina, error) {
	// Se separan los ejercicios que tengan el tipoMaximizado y los que no lo tengan
	ejerciciosPrioridad := []Ejercicio{}
	ejerciciosNoPrioritarios := []Ejercicio{}

	for e := lle.Head(); e != nil; e = e.Next() {
		ejer := e.Data()
		esPrioritario := ejer.EsDeTipo(tipoMaximizado)
		if esPrioritario {
			ejerciciosPrioridad = append(ejerciciosPrioridad, ejer)
		} else {
			ejerciciosNoPrioritarios = append(ejerciciosNoPrioritarios, ejer)
		}
	}

	// Se ordena el array prioritario con pts/tiempo (maximos puntos - minima duracion)
	n := len(ejerciciosPrioridad)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			ptsPorTiempo := float64(ejerciciosPrioridad[j].PtsDeTipo(tipoMaximizado)) / float64(ejerciciosPrioridad[j].TiempoInt())
			elSiguiente := float64(ejerciciosPrioridad[j+1].PtsDeTipo(tipoMaximizado)) / float64(ejerciciosPrioridad[j+1].TiempoInt())
			if ptsPorTiempo < elSiguiente {
				ejerciciosPrioridad[j], ejerciciosPrioridad[j+1] = ejerciciosPrioridad[j+1], ejerciciosPrioridad[j]
			}
		}
	}

	//Se añaden los ejercicios hasta llegar a la duracion total
	ejerciciosSeleccionados := list.NewLinkedList[Ejercicio]()
	idsEjercicios := []string{}
	tiempoTotal := 0

	for _, ejercicio := range ejerciciosPrioridad {
		if tiempoTotal+ejercicio.TiempoInt() <= duracionTotal {
			ejerciciosSeleccionados.Append(ejercicio)
			tiempoTotal += ejercicio.TiempoInt()
			idsEjercicios = append(idsEjercicios, ejercicio.Id)
		}
	}

	//si queda tiempo disponibele en la duración establecida se añaden ejercicios no prioritarios
	if tiempoTotal < duracionTotal {
		for _, ejercicio := range ejerciciosNoPrioritarios {
			if tiempoTotal+ejercicio.TiempoInt() <= duracionTotal {
				ejerciciosSeleccionados.Append(ejercicio)
				tiempoTotal += ejercicio.TiempoInt()
				idsEjercicios = append(idsEjercicios, ejercicio.Id)
			}
		}
	}

	//se retorna la rutina con los ids seleccionados (no se tiene en cuenta el error porque se sabe que los ejercicios existen)
	idsEjerciciosStr := strings.Join(idsEjercicios, ",")
	rutina, _ := NewRutina(nombre, idsEjerciciosStr, ejerciciosSeleccionados)
	return rutina, nil
}
