package options

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/NewbJS/Go-Todo-CLI/todo-handler/todo"
	"github.com/NewbJS/Go-Todo-CLI/todo-handler/todofunc"
)

var todoSlice []todo.Todo = make([]todo.Todo, 0)

// Options should be defered on every function, to make sure the program stays running.
func Options() {
	optList := [6]string{"List todos", "Add todo", "Edit todo", "Copy todo", "Delete todo", "Quit"}
	fmt.Println("What would you like to do next? Your options are:")
	for i := 0; i < len(optList); i++ {
		if i+1 < len(optList) {
			fmt.Print(i+1, ": ", optList[i], ", ")
		} else {
			fmt.Print("and ", i+1, ": ", optList[i])
		}
	}
	fmt.Print("\n")
	var choice string
	fmt.Scanln(&choice)
	fmt.Print("\n")
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	switch choice {
	case "1":
		todofunc.ListTodos(todoSlice, true, Options)
	case "2":
		todofunc.AddTodo(&todoSlice, Options)
	case "3":
		todofunc.EditTodo(&todoSlice, Options)
	case "4":
		todofunc.CopyTodo(&todoSlice, Options)
	case "5":
		todofunc.DeleteTodo(&todoSlice, Options)
	case "6":
		fmt.Println("Quiting...")
		return
	default:
		fmt.Printf("'%v' is not a valid option.\n", choice)
		todofunc.ListTodos(todoSlice, true, Options)
	}
}
