package gotaskscheduler

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

//Function type definition
type Fn func()

//Timer type definition
type timer struct {
	name     string
	once     bool
	seconds  uint32
	function Fn
}

//List type definition
type TList struct {
	id      int
	name    string
	seconds uint32
	once    bool
}

//Timers (internal list)
var timers = map[int]*timer{}

//true while scheduler is started
var started bool

//true while scheduler is stopping
var doStop bool

//last task id
var lastTimer int

//max tasks
var tasksLimit int = 50

//Set max tasks, default is 50. Warning: too many tasks can dangerously increase CPU usage
func SetTasksLimit(limit int) {
	tasksLimit = limit
}

//Add new Task to execute every specified seconds (once=true for only one execution).
//Name can be an empty string.
//seconds will be 1 to 100years (in seconds).
//function is a func task
//once determine if it's an one time task
func AddTask(name string, seconds uint32, function Fn, once bool) (id int, err error) {
	if len(timers) >= tasksLimit {
		err = errors.New("too many tasks")
		return 0, err
	}

	//3155760000 seconds is 100 years, should be enough
	if seconds >= uint32(3155760001) {
		err = errors.New("too much time")
		return 0, err
	}

	//this should never happen
	if lastTimer >= 2147483645 {
		err = errors.New("too many tasks, danger of overflow")
	}

	//increase id
	lastTimer++
	id = lastTimer

	//if no name, set default "Task" + ID
	if name == "" {
		name = fmt.Sprintf("Task %d", id)
	}

	timers[id] = &timer{name, once, seconds, function}

	err = nil
	return id, err
}

//Manually task execution
func ExecTask(id int) (err error) {
	if timers[id] != nil {
		timers[id].function()
		return nil
	}
	err = errors.New("task does not exist")
	return err
}

//Count total tasks.
func CountTasks() int {
	return len(timers)
}

//List Tasks (ID, NAME, SECONDS)
func ListTasks() (list map[int]*TList) {
	list = make(map[int]*TList)

	for key, value := range timers {
		list[key] = &TList{key, value.name, value.seconds, value.once}
	}

	return list
}

//Delete a Task in the timers list (if exists)
func DeleteTask(id int) {
	delete(timers, id)
}

//Delete all Tasks in the timers list (if exists)
func DeleteAllTasks() {
	timers = map[int]*timer{}
}

//Stop Scheduler (and optionally Delete all Tasks) (if prev started)
func StopScheduler(DelTasks bool) {
	if !started {
		return
	}

	doStop = true

	if DelTasks {
		DeleteAllTasks()
		lastTimer = 0
	}

	//Wait until routine timers stop
	for !started && !doStop {
		time.Sleep(10 * time.Millisecond)
	}
}

//Start Scheduler.
func StartScheduler() {
	//Wait until previus routine timers stop
	for doStop {
		time.Sleep(10 * time.Millisecond)
	}

	//Prevent double excecution
	if started {
		return
	}

	started = true

	//Run Scheduler loop in go routine
	go func() {
		var tick uint32
		tick = uint32(time.Now().Unix())

		//While (!doStop)
		for !doStop {
			//wait 1 second (sleep routine)
			for uint32(time.Now().Unix()) < tick+1 {
				time.Sleep(10 * time.Millisecond)
			}

			for key, value := range timers {
				//calculate seconds Mod of tick = 0
				if tick%(value.seconds) == 0 {
					//Run timer function as go routine (async)
					go func(function func()) {
						function()
						return
					}(value.function)
					if value.once {
						DeleteTask(key)
					}
				}
			}
			go func() {
				runtime.GC()
			}()
		}
		started = false
		doStop = false
		return
	}()
}
