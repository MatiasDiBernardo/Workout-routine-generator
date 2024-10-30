package code

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/untref-ayp2/data-structures/list"
)

// Abría que agregar a getEjerciciso y getFuncions la posibilidad de leer los datos
// del csv general y los del test, así tenemos un environment para probar separado del main

// Para test:
// const ARCHIVO_EJERCICIOS = "ejercicios_test.csv"
// const ARCHIVO_RUTINAS = "rutinas_test.csv"

// Para main:
const ARCHIVO_EJERCICIOS = "ejercicios.csv"
const ARCHIVO_RUTINAS = "rutinas.csv"

//type Client struct { // Our example struct, you can use "-" to ignore a field
//	Id            string `csv:"client_id"`
//	Name          string `csv:"client_name"`
//	Age           string `csv:"client_age"`
//	NotUsedString string `csv:"-"`
//}

func GetEjercicios() *list.LinkedList[Ejercicio] { //[]Ejercicio {

	file, err := os.Open(ARCHIVO_EJERCICIOS) //Abre el archivo utilizando la función os.Open().
	if err != nil {
		panic(err) //lanza una excepción con panic(err)
	}
	defer file.Close()

	var ejercicios []Ejercicio

	if err := gocsv.UnmarshalFile(file, &ejercicios); err != nil { //se utiliza la función gocsv.UnmarshalFile() para leer los datos del archivo CSV y convertirlos en una lista de estructuras de Ejercicio
		panic(err)
	}

	return NewLinkedListEjercicios(ejercicios)
}

func UpdateEjerciciosCSV(lle *list.LinkedList[Ejercicio]) error { //actualiza el archivo csv
	arrayEjercicios := []Ejercicio{} //slice de tipo Ejercicios

	if !lle.IsEmpty() { //se verifica si esta vacia la lista

		for nodo := lle.Head(); nodo != nil; nodo = nodo.Next() { //se itera sobre los nodos del primer al ultimo
			arrayEjercicios = append(arrayEjercicios, nodo.Data()) //agrega dato de cada del nodo actual al slice
		}

		file, err := os.Create(ARCHIVO_EJERCICIOS) //crea el archivo
		if err != nil {
			return err
		}
		defer file.Close()

		if err := gocsv.MarshalFile(&arrayEjercicios, file); err != nil { //se escribn los datos del array sobre el archivo
			return err
		}

		return nil
	}

	return nil
}

func GetRutinas(ejercicios *list.LinkedList[Ejercicio]) *list.LinkedList[Rutina] { //lee los datos de las rutinas desde un archivo CSV y devuelve una lista enlazada de rutinas

	file, err := os.Open(ARCHIVO_RUTINAS)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rutinas []Rutina                                        //slice que almacenar los datos de las rutinas que se leerán del archivo CSV.
	if err := gocsv.UnmarshalFile(file, &rutinas); err != nil { // lee los datos del archivo CSV y los vuelca en la variable rutinas.
		panic(err)
	}

	return NewLinkedListRutinas(rutinas, ejercicios) //Crea lista de rutina a partir de los datos leidos y la lista enlazada de ejercicios
}

func UpdateRutinasCSV(llr *list.LinkedList[Rutina]) error { //actualiza un archivo CSV que contiene datos de rutinas
	arrayRutinas := []Rutina{}

	if !llr.IsEmpty() {

		for nodo := llr.Head(); nodo != nil; nodo = nodo.Next() {
			arrayRutinas = append(arrayRutinas, nodo.Data())
		}

		file, err := os.Create(ARCHIVO_RUTINAS)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := gocsv.MarshalFile(&arrayRutinas, file); err != nil {
			return err
		}

		return nil
	}

	return nil
}
