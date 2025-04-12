package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tarea struct {
	Descripcion string
	Completada  bool
}

func (t Tarea) MostrarTarea() {
	var estado string
	estado = "Pendiente"
	if t.Completada {
		estado = "Completada"
	}
	fmt.Printf("%s - [%s]\n", t.Descripcion, estado)
}

func (t Tarea) MostrarDescripcionTarea() {
	fmt.Printf("%s\n", t.Descripcion)
}

func main() {
	var tareas []Tarea

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("¡Bienvenido al gestor de tareas!")
	fmt.Println("Opciones:")
	fmt.Println("1. Agregar tarea")
	fmt.Println("2. Listar tareas")
	fmt.Println("3. Marcar tarea como completada")
	fmt.Println("4. Salir")

	for {
		fmt.Print("Selecciona una opción:")

		scanner.Scan()
		opcion := scanner.Text()

		switch opcion {
		case "1":
			fmt.Print("\nEscribe la descripción de la tarea:")
			scanner.Scan()
			descripcion := scanner.Text()
			tareas = append(tareas, Tarea{Descripcion: descripcion, Completada: false})
			fmt.Println("Tarea agregada con éxito")

		case "2":
			fmt.Println("\nLista de Tareas:")
			for i, tarea := range tareas {
				fmt.Printf("%d.", i+1)
				tarea.MostrarTarea()
			}
		case "3":
			fmt.Println("Lista de Tareas:")

			if len(tareas) == 0 {
				fmt.Println("No hay tareas registradas")
			} else {
				for i, tarea := range tareas {
					fmt.Printf("%d. - ", i+1)
					tarea.MostrarDescripcionTarea()
				}
				fmt.Println("Elige la tarea a completar:")
				scanner.Scan()
				numeroStr := scanner.Text()
				numero, err := strconv.Atoi(numeroStr)

				if err != nil {
					fmt.Println("Número inválido. Recuerda que debes indicar es un número")
				} else {
					if numero > 0 && numero <= len(tareas) {
						tareas[numero-1].Completada = true
						fmt.Println("Tarea completada con éxito.")
					} else {
						fmt.Println("Opción no válida")
					}
				}
			}
		case "4":
			fmt.Println("Gracias por usar nuestra aplicación")
			return
		}
	}

}
