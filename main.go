package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func leerEntrada() string {
	// scanner para leer lo que ingrese el usuario
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func (t *task) marcarCompletado() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (tl *taskList) agregarALista(newTask *task) {
	tl.tasks = append(tl.tasks, newTask)
}

func (tl *taskList) eliminarDeList(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func imprimirTitulo(titulo string) {
	fmt.Printf("\n------------------------------ %s -----------------------------\n", strings.ToUpper(titulo))
}

func imprimirTarea(tarea *task, index int) {
	fmt.Printf("\nNumero de tarea: %d", index+1)
	fmt.Printf("\nNombre: %s", tarea.nombre)
	fmt.Printf("\nDescripcion: %s", tarea.descripcion)
	fmt.Printf("\nCompletada: %t\n", tarea.completado)
}

func (tl *taskList) verLista() {
	if len(tl.tasks) == 0 {
		fmt.Println("La lista esta vacia")
	} else {
		imprimirTitulo("lista de tareas")
		for index, tarea := range tl.tasks {
			imprimirTarea(tarea, index)
		}
	}
}

func (tl *taskList) verListaFiltrada(completada bool) {
	if completada {
		imprimirTitulo("lista de tareas completadas")
	} else {
		imprimirTitulo("lista de tareas no completadas")
	}

	for index, tarea := range tl.tasks {
		if tarea.completado == completada {
			imprimirTarea(tarea, index)
		}
	}
}

func main() {

	var opcion string
	lista := &taskList{tasks: []*task{}}

	for opcion != "0" {
		fmt.Println("\nElija la opcion que desea (Solo el numero ej: 1)")
		fmt.Println("1. Ver lista")
		fmt.Println("2. Crear nueva tarea")
		fmt.Println("3. Eliminar tarea")
		fmt.Println("4. Modificar Tarea")
		fmt.Println("0. Salir")
		opcion = leerEntrada()

		switch opcion {
		case "0":
			fmt.Println("Gracias por usar este programa.")

		case "1":
			// Ver lista de tareas
			fmt.Println("\nElija la opcion que desea (Solo el numero ej: 1)")
			fmt.Println("1. Ver toda la lista")
			fmt.Println("2. Ver tareas completas")
			fmt.Println("3. Ver tareas no completas")
			opcion2 := leerEntrada()

			switch opcion2 {
			case "1":
				lista.verLista()
			case "2":
				lista.verListaFiltrada(true)
			case "3":
				lista.verListaFiltrada(false)
			default:
				fmt.Println("Opcion invalida")

			}

		case "2":
			// Crear tarea
			fmt.Println("\nIngrese el nombre de la tarea:")
			nombreI := leerEntrada()
			fmt.Println("Ingrese la descripcion de la tarea:")
			descripcionI := leerEntrada()

			t := &task{
				nombre:      nombreI,
				descripcion: descripcionI,
			}
			lista.agregarALista(t)

		case "3":
			// Eliminar tarea
			fmt.Println("\nIngrese el numero de la tarea:")
			indexTarea, err := strconv.Atoi(leerEntrada())

			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			lista.eliminarDeList(indexTarea - 1)

		case "4":
			// Actualizar tarea
			fmt.Println("\nElija la opcion que desea (Solo el numero ej: 1)")
			fmt.Println("1. Modificar nombre")
			fmt.Println("2. Modificar descripcion")
			fmt.Println("3. Marcar como completada")

			opcion2 := leerEntrada()
			fmt.Println("\nIngrese el numero de la tarea:")
			indexTarea, err := strconv.Atoi(leerEntrada())

			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			switch opcion2 {
			case "1":
				fmt.Println("\nIngrese el nuevo nombre de la tarea:")
				nombre := leerEntrada()
				lista.tasks[indexTarea-1].actualizarNombre(nombre)
			case "2":
				fmt.Println("\nIngrese la nueva descripcion de la tarea:")
				descripcion := leerEntrada()
				lista.tasks[indexTarea-1].actualizarDescripcion(descripcion)
			case "3":
				lista.tasks[indexTarea-1].marcarCompletado()
			default:
				fmt.Println("Opcion invalida")

			}

		default:
			fmt.Println("Opcion no valida elejir: 1, 2, 3, 4, 0")
		}
	}

}
