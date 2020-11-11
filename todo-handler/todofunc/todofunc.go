package todofunc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/NewbJS/Go-Todo-CLI/todo-handler/todo"
)

// removeIndex(s []Todo, index int8) []Todo should only be called inside an option.
func removeIndex(s []todo.Todo, index int8) []todo.Todo {
	return append(s[:index], s[index+1:]...)
}

// DeleteTodo will handle deleting todos.
func DeleteTodo(ts *[]todo.Todo, defered func()) {
	defer defered()

	// Make sure the user has some todos.
	if len(*ts) > 0 {
		ListTodos(*ts, false, defered)
		var choice string
		fmt.Println("Which todo number do you want to delete?")
		fmt.Scanln(&choice)
		index, err := strconv.ParseInt(choice, 10, 8)
		if err != nil {
			fmt.Println("Enter a number.")
			return
		}
		fmt.Printf("Deleted '%v'.\n", (*ts)[index-1].Desc)
		*ts = removeIndex(*ts, int8(index-1))

	} else {
		fmt.Println("You have no todos. Add some.")
		return
	}
}

// EditTodo will handle editing todos.
func EditTodo(ts *[]todo.Todo, defered func()) {
	defer defered()

	// Make sure the user has some todos.
	if len(*ts) > 0 {
		ListTodos(*ts, false, defered)
		fmt.Println("Which todo number would you like to edit?")
		choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")
		index, err := strconv.ParseInt(choice, 10, 64)
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
		fmt.Println("What new text do you want?")
		newText, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		newText = strings.TrimSuffix(newText, "\n")
		fmt.Printf("Edited '%v' to be '%v'.\n", (*ts)[index-1].Desc, newText)
		(*ts)[index-1].Desc = newText
		(*ts)[index-1].Edited = true
		(*ts)[index-1].TimeEdited = time.Now().Format(time.UnixDate)
	} else {
		fmt.Println("You have no todos. Add some.")
		return
	}

}

// ListTodos lists all todos.
func ListTodos(slice []todo.Todo, isDefer bool /* isDefer should be false when this function is being called inside an option. */, defered func()) {

	/*
		listTodos(slice []Todo, isDefer bool) is called inside of two functions, which are already defering options().
		This if statement allows us to make sure options() isn't called when it isn't wanted through the isDefer paramter.
	*/
	if isDefer {
		defer defered()
	}

	// Make sure the user has some todos.
	if len(slice) > 0 {
		fmt.Println("Your todos are:")
		fmt.Print("\n")
		for i := 0; i < len(slice); i++ {
			fmt.Print("TODO NUMBER ", i+1, ":\n")
			fmt.Println("Description:", slice[i].Desc)
			fmt.Println("Time published:", slice[i].TimePub)
			if slice[i].Edited {
				fmt.Println("Time edited:", slice[i].TimeEdited)
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("You have no todos. Add some.")
		return
	}
}

// AddTodo will handle adding todos.
func AddTodo(ts *[]todo.Todo, defered func()) {
	defer fmt.Print("\n")
	defer defered()
	fmt.Print("Enter a new todo description: ")
	descScanner := bufio.NewReader(os.Stdin)
	newDesc, err := descScanner.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	newDesc = strings.TrimSuffix(newDesc, "\n")
	newTodo := todo.Todo{Desc: newDesc, TimePub: time.Now().Format(time.UnixDate), TimeEdited: "", Edited: false}
	*ts = append(*ts, newTodo)
	fmt.Printf("Successfully added '%v' to your todos.\n", newTodo.Desc)
}
