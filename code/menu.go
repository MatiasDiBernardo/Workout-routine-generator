package code

// Este file contiene la funcionalidad de agregar, borrar, modificar, consultar y listar los datos
// La implementación que haría yo sería tener cargados en el programa dos link list, una de ejercicios
// y otra de rutinas

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/untref-ayp2/data-structures/list"
)

// ______     _ ______ _____   _____ _____ _____ _____ ____   _____
//|  ____|   | |  ____|  __ \ / ____|_   _/ ____|_   _/ __ \ / ____|
//| |__      | | |__  | |__) | |      | || |      | || |  | | (___
//|  __| _   | |  __| |  _  /| |      | || |      | || |  | |\___ \
//| |___| |__| | |____| | \ \| |____ _| || |____ _| || |__| |____) |
//|______\____/|______|_|  \_\\_____|_____\_____|_____\____/|_____/

func findEjercicioById(id string, lle *list.LinkedList[Ejercicio]) *list.LinkedNode[Ejercicio] {

	if !lle.IsEmpty() {
		for node := lle.Head(); node != nil; node = node.Next() {
			if node.Data().Id == id {
				return node
			}
		}
	}
	return nil
}

// ANIAdir
func AgregarEjercicio(lle *list.LinkedList[Ejercicio], newEj *Ejercicio) error {
	//Busca ultimo id
	ultimoIDint, _ := strconv.Atoi(lle.Tail().Data().Id)
	//Le asigna el ultimo id + 1 al nuevo ejercicio
	newEj.Id = fmt.Sprint(ultimoIDint + 1)
	//ANIade al nuevo ejercicio a la linkedlist
	lle.Append(*newEj)
	//Actualiza Csv
	err := UpdateEjerciciosCSV(lle)

	return err
	/*
		if err != nil {
			fmt.Println("> No se pudo aniadir el ejercicio - error en el archivo de guardado")
		}
		fmt.Println("> Se aniadio el ejercicio [id: ", newEj.Id, "]")
	*/
}

// Borrar
func BorrarEjercicio(lle *list.LinkedList[Ejercicio], id string, llr *list.LinkedList[Rutina]) error {
	// TODO: Borrar también en las rutinas que aparece este ejercicio
	// meto una validación, mucho quilombo sino -> "si el ejer está en una rutina, no se puede borrar"
	// Creo que con eso alcanza [MD]. Otra cosa que cuando borramos estaría bueno recorrer y restar
	// menos 1 a los IDs de todos los siguientes así tenemos la lista siempre con todos los IDs completados.
	// si hacemos eso rompemos todas las rutinas porque habria que cambiar los ids tambien ahi D: [OA]

	nodeTarget := findEjercicioById(id, lle)

	if nodeTarget != nil {
		if !existeEnRutinas(nodeTarget.Data().Id, llr) {
			lle.Remove(nodeTarget.Data())
			//actualiza el csv
			err := UpdateEjerciciosCSV(lle)
			return err
			/*
				if err != nil {
					return err
				}
				fmt.Println("> Se borro el ejercicio [id: ", id, "]")
				return nil
			*/
		} else {
			return errors.New("no se puede borrar el ejercicio porque se utiliza en rutinas")
		}
	}
	return errors.New("Ejercicio no encontrado")
}

// Terminar // capaz agregar id de rutina
func existeEnRutinas(id string, llr *list.LinkedList[Rutina]) bool {
	for node := llr.Head(); node != nil; node = node.Next() {
		for _, ejer := range strings.Split(node.Data().Ejercicios, ",") {
			if ejer == id {
				return true
			}
		}
	}
	return false
}

// Modificar
func ModificarEjercicio(lle *list.LinkedList[Ejercicio], id, fieldToChange, valueToChange string) error {
	// busca si existe el ejercicio por Id
	nodeTarget := findEjercicioById(id, lle)

	if nodeTarget != nil {
		//Trae el ejercicio encontrado
		ejer := nodeTarget.Data()
		//Modifica el campo especificado
		err := ejer.ModificarDatos(fieldToChange, valueToChange)
		if err != nil {
			return err
		}
		//Modifica el nodo de la linked list
		nodeTarget.SetData(ejer)
		//chequear despues si hace falta o se puede hacer nodeTarget.Data().ModificarDatos(fieldToChange, valueToChange) directamente [OA]
		//actualiza el csv
		err = UpdateEjerciciosCSV(lle)
		return err
		/*
			if err != nil {
				return err
			}
			fmt.Println("> Se modifico el ejercicio [id: ", ejer.Id, ", ", fieldToChange, ": ", valueToChange, "]")
			return nil
		*/
	}
	return errors.New("Ejercicio no encontrado")
}

// Consulta
func MostrarEjercicio(lle *list.LinkedList[Ejercicio], idEjer string) (Ejercicio, error) {
	nodeTarget := findEjercicioById(idEjer, lle)
	if nodeTarget != nil {
		ejer := nodeTarget.Data()
		ejer.mostrarEjercicio()
		return ejer, nil
	} else {
		return Ejercicio{}, errors.New("Ejercicio no encontrado")
	}

	/*
		if nodeTarget != nil {
			ejer := nodeTarget.Data()
			ejer.mostrarEjercicio()
		} else {
			fmt.Println("No se encontro el ejercicio solicitado [id: ", id, "]")
		}*/

}

// Listado
func ListarEjercicios(lle *list.LinkedList[Ejercicio]) []Ejercicio {
	ejercicios := []Ejercicio{}

	if lle.Size() > 0 {
		for node := lle.Head(); node != nil; node = node.Next() {
			ejer := node.Data()
			ejercicios = append(ejercicios, ejer)
			ejer.mostrarEjercicio()
		}
	}
	return ejercicios
}

//_____  _    _ _______ _____ _   _           _____
//|  __ \| |  | |__   __|_   _| \ | |   /\    / ____|
//| |__) | |  | |  | |    | | |  \| |  /  \  | (___
//|  _  /| |  | |  | |    | | | . ` | / /\ \  \___ \
//| | \ \| |__| |  | |   _| |_| |\  |/ ____ \ ____) |
//|_|  \_\\____/   |_|  |_____|_| \_/_/    \_\_____/

func findRutinaById(id string, llr *list.LinkedList[Rutina]) *list.LinkedNode[Rutina] {

	if !llr.IsEmpty() {
		for node := llr.Head(); node != nil; node = node.Next() {
			if node.Data().Id == id {
				return node
			}
		}
	}
	return nil
}

// Añadir
func AgregarRutina(llr *list.LinkedList[Rutina], newRut *Rutina) error {

	//Busca ultimo id
	ultimoIDint, _ := strconv.Atoi(llr.Tail().Data().Id)
	//Le asigna el ultimo id + 1 a la nueva rutina
	newRut.Id = fmt.Sprint(ultimoIDint + 1)
	//Añade la nueva rutina a la linkedlist
	llr.Append(*newRut)
	//Actualiza Csv
	err := UpdateRutinasCSV(llr)

	return err
	/*
		if err != nil {
			fmt.Println("> No se pudo aniadir la rutina - error en el archivo de guardado")
		}
		fmt.Println("> Se aniadio la rutina [id: ", newRut.Id, "]")
	*/
}

// Borrar
func BorrarRutina(llr *list.LinkedList[Rutina], id string) error {
	nodeTarget := findRutinaById(id, llr)

	if nodeTarget != nil {
		llr.Remove(nodeTarget.Data())
		//actualiza el csv
		err := UpdateRutinasCSV(llr)
		return err

		/*
			if err != nil {
				return err
			}
			fmt.Println("> Se borro la rutina [id: ", id, "]")
			return nil
		*/
	}
	return errors.New("Rutina no encontrada")
}

// Modificar nombre con keyword Nombre. Modificar ejercicios con keyword Ejercicios y un string con formato "1, 2, 3" que reemplaza a los ejs existentes
func ModificarRutina(llr *list.LinkedList[Rutina], lle *list.LinkedList[Ejercicio], id, fieldToChange, valueToChange string) error {
	// busca si existe la rutina por Id
	nodeTarget := findRutinaById(id, llr)

	if nodeTarget != nil {
		//Trae la rutina encontrada
		rut := nodeTarget.Data()
		//Modifica el campo específicado
		err := rut.ModificarDatos(fieldToChange, valueToChange, lle)
		if err != nil {
			return err
		}

		//Modifica el nodo de la linked list
		nodeTarget.SetData(rut)
		//actualiza el csv
		err = UpdateRutinasCSV(llr)
		return err
		/*
			if err != nil {
				return err
			}
			fmt.Println("> Se modifico la rutina [id: ", rut.Id, ", Nombre: ", valueToChange, "]")
			return nil
		*/
	}

	return errors.New("Ejercicio no encontrado")
}

// Modificar - Add o remove Ejercicios. Para agregar accion == "A". Para remover accion == "R"
func AddRemoveEjercicioToRutina(llr *list.LinkedList[Rutina], lle *list.LinkedList[Ejercicio], id, idEjer, accion string) error {
	nodeTarget := findRutinaById(id, llr)
	nodeTargetEjer := findEjercicioById(idEjer, lle)

	if nodeTarget != nil {
		if nodeTargetEjer != nil {
			rut := nodeTarget.Data()
			ejer := nodeTargetEjer.Data()
			if accion == "A" {
				rut.AddEjer(ejer)
				nodeTarget.SetData(rut)

				err := UpdateRutinasCSV(llr)
				if err != nil {
					return err
				}
				//fmt.Println("> Se anadio el ejercicio [id: ", ejer.Id, "] a la rutina [id: ", rut.Id, "]")
				return nil

			} else if accion == "R" {
				rut.RemoveEjer(ejer)
				nodeTarget.SetData(rut)

				err := UpdateRutinasCSV(llr)
				if err != nil {
					return err
				}
				//fmt.Println("> Se elimino el ejercicio [id: ", ejer.Id, "] a la rutina [id: ", rut.Id, "]")
				return nil
			}
		}
		//fmt.Println("> No se encontro el ejercicio [id: ", idEjer, "]")
		return errors.New(fmt.Sprint("No se encontro el ejercicio [id: ", idEjer, "]"))
	}
	//fmt.Println("> No se encontro la rutina [id: ", id, "]")
	return errors.New(fmt.Sprint("> No se encontro la rutina [id: ", id, "]"))
}

// Consulta
func MostrarRutina(llr *list.LinkedList[Rutina], idRut string) (Rutina, error) {
	nodeTarget := findRutinaById(idRut, llr)
	if nodeTarget != nil {
		rut := nodeTarget.Data()
		rut.MostrarRutina()
		return rut, nil
	} else {
		return Rutina{}, errors.New("Rutina no encontrada")
	}

	/*
		if nodeTarget != nil {
			rut := nodeTarget.Data()
			rut.mostrarRutina()
		} else {
			fmt.Println("No se encontro la rutina solicitado [id: ", idRut, "]")
		}
	*/
}

// Listado
func ListarRutinas(llr *list.LinkedList[Rutina]) []Rutina {
	rutinas := []Rutina{}

	if llr.Size() > 0 {
		for node := llr.Head(); node != nil; node = node.Next() {
			rut := node.Data()
			rutinas = append(rutinas, rut)
			rut.MostrarRutina()
		}
	}
	return rutinas
}
