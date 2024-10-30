# Rutinas de Entrenamiento por MALO

## Descripción
Este proyecto consiste en el desarrollo de una aplicación para la gestión de rutinas de entrenamiento. La aplicación permite elegir entre seguir una rutina predefinida o crear una nueva de manera dinámica.

## Estructura del Proyecto
El proyecto está organizado en los siguientes archivos:

- `ejercicio.go`: Contiene la definición y funcionalidades relacionadas con los ejercicios.
- `rutina.go`: Contiene la definición y funcionalidades relacionadas con las rutinas.
- `manejocsv.go`: Funciones para la gestión de datos (lectura y escritura de archivos CSV).
- `menu.go`: Proporciona una interfaz de usuario para que interactúe con las diferentes funcionalidades del programa.
- `CLI.go`: Contiene el desarrollo del menu interactivo para el usuario.


## Ejercicios
####  Crear un Ejercicio

Para crear un nuevo ejercicio, se debe llamar a la funcion `NewEjercicio()` con los siguientes parámetros:
- **Nombre**: Nombre del ejercicio.
- **Descripción**: Descripción detallada del ejercicio.
- **Tiempo**: Duración estimada del ejercicio.
- **Calorías**: Cantidad de calorías quemadas por el ejercicio.
- **Tipo**: Tipo de ejercicio (por ejemplo, cardio, fuerza, flexibilidad).
- **Grupo Muscular**: Grupo muscular al que se dirige el ejercicio.
- **Pts**: Puntos asignados al ejercicio para cada uno de sus tipos.
- **Dificultad**: Nivel de dificultad del ejercicio.

La funcion retorna un struct de tipo `Ejercicio`


#### Funciones disponibles para el uso de Ejercicios
- `mostrarEjercicio()`: Muestra los detalles del ejercicio en la consola, pertenece al struct Ejercicio y receptor es un puntero a un Ejercicio.
- `ModificarDatos()`: Permite modificar los campos del ejercicio. Recibe como parámetros el campo que se desea modificar y el nuevo valor para ese campo. Retorna un error si el campo especificado es inválido o si ocurre algún problema durante la modificación.
- `numberLevelToString()`: Recibe como parámetro el nivel de dificultad del ejercicio en formato string. Retorna una cadena de texto que representa el nivel de dificultad en palabras ("Bajo", "Medio", "Alto" o "Error en la dificultad").
- `NewLinkedListEjercicios`: Crea una nueva linked list de Ejercicios. Recibe como parametro una lista de Ejercicios `listaEj []Ejercicio` y retorna un puntero a una linked list de Ejercicio `*list.LinkedList[Ejercicio]` .
- **Funciones auxiliares**
- `IDInt()`, `TiempoInt()`, `CaloriasInt()`, `PuntosInt()`:  Estos métodos reciben un campo específico del ejercicio en formato string y retornan su valor convertido a int cuando sea necesario.
- `ArrayTipos()`: Retorna un array con los tipos del ejercicio
- `EsDeTipo()`: Retorna un booleano que representa si el ejercicio es del tipo especificado por parámetro
- `PtsDeTipo()`: Retorna los puntos del tipo especificado por parametro

## Rutinas
####  Crear una Rutina

Para crear una nueva rutina, se debe llamar a la funcion `NewRutina()` con los siguientes parámetros:

- `Nombre`: Nombre de la rutina.
- `Ejercicios`: String que almacena los IDs de los ejercicios de la rutina separados por comas.
- `EjerciciosDisponibles`: Lista enlazada de los ejercicios disponibles para crear la rutina.

La funcion retorna un struct de tipo `Rutina`

#### Funciones disponibles para el uso de Rutinas
- `NewLinkedListEjerciciosById()`: Esta función recibe una lista enlazada de ejercicios disponibles y una cadena de IDs de ejercicios. Retorna una lista enlazada de ejercicios correspondiente a los IDs proporcionados. Su funcion es validar los ids ingresados y luego utilizarlos en la rutina como Ejercicios
- ` MostrarRutina()`: La función mostrarRutina pertenece al struct Rutina y se utiliza para mostrar los detalles de la rutina en la consola.
- `ModificarDatos()`: Permite modificar los campos de la rutina. Recibe como parámetros el campo que se desea modificar y el nuevo valor para ese campo. Retorna un error si el campo o valor especificado es inválido o si ocurre algún problema durante la modificación.
- `AddEjer()`: Ésta función agrega un nuevo ejercicio a la lista de ejercicios de una rutina, recibe como parámetro un objeto de tipo Ejercicio. Asi mismo actualiza la cadena de IDs de ejercicios de la rutina con el ID del nuevo ejercicio agregado. 
- `RemoveEjer()`: Ésta función elimina un ejercicio específico de la lista de ejercicios de una rutina y actualiza la cadena de IDs de ejercicios de la rutina sin incluir el ID del ejercicio eliminado. Recibe como parámetro un objeto de tipo Ejercicio.
- `NewLinkListRutinas()`:Crea y retorna una Linked List de rutinas a partir de un array de structs Rutina  y otro LinkedList de ejercicios disponibles para agregar a las mismas
- `Calorias()`: Calcula y retorna la cantidad total de calorías quemadas en una rutina sumando las calorías de cada ejercicio presente en la Rutina.
- `Duracion()`: Calcula la duración total de una rutina sumando los tiempos de cada ejercicio presente en la Rutina. 
- `Dificultad()`: Retorna la dificultad mas frencuente entre los ejercicios de la Rutina
- `TipoDeEjercicios()`: Retorna el tipo mas frencuente entre los ejercicios de la Rutina
- `EjerciciosDeRutina()`: Retorna un array de Ejercicios que componen a la Rutina
- `NewRutinaPorTipoyDificultad()`: (Automágica parte 1) Retorna una Rutina creada de forma automágica  a partir de un tipo y dificultad maximizando la cantidad de ejercicios en un tiempo determinado 
- `NewRutinaPorCalorias()`: (Automágica parte 2) Retorna una Rutina creada de forma automágica  a partir de un unas calorías definidas a quemar y con el criterio de retornar la menor cantidad de ejercicios
- `NewRutinaPorTipoPtsYDuracion()`: (Automágica parte 3) Retorna una Rutina creada de forma automágica  a partir de un tipo que se busca priorizar y una duracion total a cubrir. Si no hay suficientes ejercicios del tipo especificado, el tiempo definido se completa con otros ejercicios de tipos no priorizados.

## Manejo de CSV
El paquete `code` proporciona funciones para el manejo de archivos CSV relacionados con ejercicios y rutinas. Estas funciones se encargan de leer los datos desde los archivos CSV y guardar los cambios realizados en ellos.
Se definen constantes que identifican los nombres de los archivos CSV utilizados para almacenar los datos de ejercicios y rutinas, respectivamente.
`ARCHIVO_EJERCICIOS`: Almacena el nombre del archivo CSV que contiene los datos de los ejercicios.
`ARCHIVO_RUTINAS`: Almacena el nombre del archivo CSV que contiene los datos de las rutinas.

#### Funciones Disponibles

- `GetEjercicios() *list.LinkedList[Ejercicio]`: Esta función lee los datos del archivo CSV de ejercicios (`ARCHIVO_EJERCICIOS`) y devuelve una lista enlazada de Ejercicios.

- `UpdateEjerciciosCSV(lle *list.LinkedList[Ejercicio]) error`: Actualiza el archivo CSV de ejercicios con los datos proporcionados por parámetro de un puntero a una lista enlazada de Ejercicio `lle`. Retorna un error si ocurre algún problema durante la actualización.

- `GetRutinas(ejercicios *list.LinkedList[Ejercicio]) *list.LinkedList[Rutina]`: Lee los datos del archivo CSV de rutinas (`rutinas.csv`) y devuelve una lista enlazada de Rutinas. Recibe como parámetro una lista enlazada de Ejercicio `ejercicios` y retorna un puntero a una lista enlazada de Rutina.

- `UpdateRutinasCSV(llr *list.LinkedList[Rutina]) error`: Actualiza el archivo CSV de rutinas con los datos proporcionados en la lista enlazada de rutinas. Recibe como parámetro un puntero a una lista enlazada de Rutina `llr` y retorna un error si ocurre algún problema.

## Menu
####  ABMCL de Ejercicios
- **Alta** `AgregarEjercicio(lle, newEj Ejercicio) error`: Agrega un nuevo ejercicio a la lista enlazada de ejercicios. Recibe como parámetros la lista enlazada de ejercicios `lle` y el nuevo ejercicio a agregar `newEj`.En el alta del ejercicio, a este se le asigna un Id para ser ubicado dentro de la `lle` y se guarda en la base de datos CSV.
- **Baja** `BorrarEjercicio(lle, id string, llr) error` : Elimina un ejercicio de la lista enlazada de ejercicios. Recibe como parámetros la lista enlazada de ejercicios `lle` y el ID del ejercicio a eliminar `id`. Retorna un error si el ejercicio no se encuentra en la lista o si el ejercicio se encuentra en la Linked List de Rutinas definida por parametro en `llr`
- **Modificación** `ModificarEjercicio(lle, id, fieldToChange, valueToChange string) error` : Permite modificar los atributos de un ejercicio existente. Recibe como parámetros la lista enlazada de ejercicios `lle`, el ID del ejercicio a modificar `id`, el campo que se desea modificar `fieldToChange` y el nuevo valor para ese campo `valueToChange`. Retorna un error si el ejercicio no se encuentra en la lista.
- **Consulta** `MostrarEjercicio(lle, id string) Ejercicio, error`: Muestra los detalles de un ejercicio específico por consola, ademas de retornar el ejercicio buscado o un error si el ejercicio no se encontró. Recibe como parámetros la lista enlazada de ejercicios `lle` al que pertenece y el ID del ejercicio a mostrar `id`.
- **Listado** `ListarEjercicios(lle) []Ejercicio`: Lista todos los ejercicios disponibles por consola ademas de retornar un array de Ejercicio. Recibe como parámetro la lista enlazada de ejercicios `lle` de la que va a mostrar los Ejercicios y el array resulta vacio si este `lle` es vacío.

### ABMCL de Rutinas
- **Alta** `AgregarRutina(llr, newRut Rutina) error`: Agrega una nueva rutina a la lista enlazada de rutinas. Recibe como parámetros la lista enlazada de rutinas  `llr ` y la nueva rutina a agregar  `newRut`. En el alta de la rutina, a esta se le asigna un Id para ser ubicado dentro de la `llr` y se guarda en la base de datos CSV.
- **Baja** `BorrarRutina(llr, id string) error` : Elimina una rutina de la lista enlazada de rutinas. Recibe como parámetros la lista enlazada de rutinas `llr` y el ID de la rutina a eliminar `id`. Retorna un error si la rutina no se encuentra en la lista definida por el parametro `llr`
- **Modificación** :`ModificarRutina(llr, lle, id, fieldToChange, valueToChange string) error` : Permite modificar campos de una rutina. Recibe como parámetros la lista enlazada de rutinas `llr`,la lista enlazada de ejercicios relacionados `lle`, el ID de la rutina a modificar `id`, el campo a modificar `fieldToChange` y el nuevo valor del campo. Retorna un error si la rutina no se encuentra en la lista o si ocurre algun problema con la modificación.
`AddRemoveEjercicioToRutina(llr, lle, id, idEjer, accion string) error`: Agrega o elimina un ejercicio de una rutina existente. Recibe como parámetros la lista enlazada de rutinas `llr`, la lista enlazada de ejercicios `lle`, el ID de la rutina `id`, el ID del ejercicio `idEjer` y la acción a realizar ("A" para agregar, "R" para remover). Retorna un error si la rutina o el ejercicio no se encuentran en la lista o si ocurre algun problema con la modificación.
- **Consulta** `MostrarRutina(llr, idRut string) Rutina,error` : Muestra los detalles de una rutina específica por consola y la retorna. Recibe como parámetros la lista enlazada de rutinas `llr` y el ID de la rutina a mostrar `idRut`. Retorna un struct de tipo rutina o un error si la rutina no se encuentra en la lista enlazada.
- **Listado** `ListarRutinas(llr) []Rutina`: Lista todas las rutinas disponibles por consola. Recibe como parámetro la lista enlazada de rutinas `llr` de la que se identifican las rutinas disponibles, las añade a un array y las retorna. El array resulta vacio si la `llr` es vacío.
