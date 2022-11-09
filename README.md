# Go Task Scheduler

- [Summary (English)](#Summary)
- [Resumen (Spanish)](#Resumen)
- [Code Example (Código de Ejemplo)](#Example)

---

# Summary

Go Task Scheduler is a basic (periodic and one-shot) task scheduler.

Go Task Scheduler does not control the state of the executed tasks (functions), nor if they have finished. It is up to the developer to control these aspects.

If you add a one-shot task to Go Task Scheduler and specify a past time, the task will not be executed (and will remain in the queue). Remember that the internal time counter starts from 0 when running StartScheduler, that will be your starting time/point

StartScheduler Starts the task scheduler. They have no effect if it is already started. It does not return information. It is not necessary to call StartScheduler in a specific order, it can be before or after the calls to AddTask.

StopScheduler Stops the task scheduler. They have no effect if it is already started. You can optionally delete all existing tasks. It does not return information.

SetTasksLimit Changes the limit of maximum tasks that can be added to the task scheduler, by default the value is 50. Warning: Too many tasks can consume a lot of CPU and Memory, depending on the function called. It does not return information.

AddTask Adds a new task to the scheduler, specifying a name (human readable, optional), if it will be executed once or periodically, the period of time (in seconds, counting from StartScheduler was executed) and the function to invoke . It can return different errors.

ExecTask Executes the function associated with the task whose ID is passed as a parameter. If the task does not exist, the execution is ignored and an error is returned. If it was successful, it returns null value.

CountTasks Returns the count of existing tasks in the scheduler.

ListTasks Returns the list of existing tasks in the scheduler (ID, Name, Time Interval (in seconds, counted from when StartScheduler was executed) and if it is a one-time task (if false, the task is periodic)

DeleteTask ​​Deletes a task (according to the supplied ID) from the scheduler. If it doesn't exist, it has no effect. It does not return information.

DeleteAllTasks Deletes all tasks from the scheduler. If none exists, it has no effect. It does not return information.

---

# Resumen

Go Task Scheduler es un programador de tareas basico (periodicas y de una sola ejecución).

Go Task Scheduler no controla el estado de las tareas (funciones) ejecutadas, ni si han finalizado. Queda a criterio del desarrollador controlar estos aspectos.

Si añade una tarea de una sola ejecución a Go Task Scheduler y especifica un tiempo pasado, la tarea no será ejecutada (y permanecerá en cola). Recuerde que el contador de tiempo interno comienza desde 0 al ejecutar StartScheduler, ese será su tiempo/punto de partida

StartScheduler Inicia el programador de tareas. No surten efecto si ya esta iniciado. No devuelve informacion. No es necesario llamar a StartScheduler en un orden especifico, puede estar antes o despues de las llamadas a AddTask.

StopScheduler Detiene el programador de tareas. No surten efecto si ya esta iniciado. Opcionalmente puede borrar todas las tareas existentes. No devuelve informacion.

SetTasksLimit Cambia el limite de tareas maximas que se pueden añadir al programador de tareas, por defecto el valor es 50. Advertencia: Demasiadas tareas pueden consumir mucho CPU y Memoria, dependiendo de la funcion invocada. No devuelve informacion.

AddTask Añade una nueva tarea al programador, especificando un nombre (legible para humanos, opcional), si se ejecutarà una sola vez o será periodico, el periodo de tiempo (en segundos, a contar desde que se ejecutó StartScheduler) y la funcion a invocar. Puede devolver diferentes errores.

ExecTask Ejecuta la funcion asociada a la tarea cuyo ID es pasado como parametro. Si la tarea no existe se ignora la ejecucion y retorna error. Si tuvo exito retorna valor nulo.

CountTasks Devuelve la cuenta de las tareas existentes en el programador.

ListTasks Devuelve la lista de tareas existentes en el programador (ID, Nombre, Intervalo de tiempo (en segundos, a contar desde que se ejecutó StartScheduler) y si es una tarea de una sola vez (si es falso, la tarea es periodica)

DeleteTask Elimina una tarea (segun el ID proporcionado) del programador. Si no existe no surte efecto. No devuelve informacion.

DeleteAllTasks Elimina todas las tareas del programador. Si no existe ninguna no surte efecto. No devuelve informacion.

# Example

```go
package main

import "github.com/SERBice/gotaskscheduler"

func myFunction() {
	fmt.Println("One print each 1 minute!")
}

func main() {
	AddTask("", 1, func() { fmt.Println("One print each second!") }, false)
	AddTask("custom task name", 12, func() { fmt.Println("One print each 12 seconds!") }, false)
	AddTask("minute task", 60, myFunction, false)
	AddTask("One Time Task", 15, func() { fmt.Println("One-Time execution at 15 seconds") }, true)
	StartScheduler()

	//Wait for CTRL+C
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	fmt.Println("Bye!")
}
```
