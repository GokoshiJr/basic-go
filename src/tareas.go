package main

import (
	"fmt"
)

func main() {
	t1 := &task {
		nombre: "Completar mi curso de Go",
		descripcion: "Haciendolo en Platzi",
	}

	t2 := &task {
		nombre: "Completar mi curso de Python",
		descripcion: "Haciendolo en Platzi",
	}

	t3 := &task {
		nombre: "Completar mi curso de Node Js",
		descripcion: "Haciendolo en Platzi",
	}

	listaTareas := &taskList {
		tasks: []*task{
			t1, t2,
		},
	}

	listaTareas.agregarALista(t3)

	for i := 0; i < len(listaTareas.tasks); i++ {
		fmt.Printf("Tarea %d %+v\n", i, listaTareas.tasks[i].nombre)
	}

	// equivalente al for each
	for i, tarea := range listaTareas.tasks {
		fmt.Printf("Tarea %d %+v\n", i, tarea.nombre)
	}

	// uso del break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // rompe el ciclo
		}
		fmt.Println(i)
	}

	// uso del continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // solamente rompe una iteracion
		}
		fmt.Println(i)
	}
}

type task struct {
	nombre string
	descripcion string
	completado bool
}

/* Receivers de tipo task */
func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

type taskList struct {
	// Slices, son como un array mas dinamico y mas poderoso
	tasks []*task // slice de tipo taks y que puede modificarlos
}

func (t *taskList) agregarALista(tarea *task) {
	// append, toma como parametro el slice y el elemento a agregar
	t.tasks = append(t.tasks, tarea)
}

func (t *taskList) eliminarDeLista(index int) {
	/* 
	Los tres puntos al final de t.task[index + 1:]… (operador ellipsis) es porque el segundo parámetro del append no es un slice y la función append recibe un item, con este operador lo que hacemos es decirle a go que tome ese slice y lo “desempaquete” para que sean muchos parámetros de 1 solo item y no un slice. t.tasks = append(t.tasks, task1, task2, task3)
	*/
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}