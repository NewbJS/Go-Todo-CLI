package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Todo is the main todo structure.
type Todo struct {
	desc    string
	timePub string

	// Only show timeEdited if edited == true.
	timeEdited string
	edited     bool
}

var todoSlice []Todo = make([]Todo, 0)

// Defer options() on every function, to make sure the program stays running.
func options() {
	optList := [5]string{"List todos", "Add todo", "Edit todo", "Delete todo", "Quit"}
	fmt.Println("What would you like to do next? Your options are:")
	for i := 0; i < len(optList); i++ {
		fmt.Print(i+1, ": ", optList[i], " ")
	}
	fmt.Print("\n")
	var choice string
	fmt.Scanln(&choice)
	fmt.Print("\n")
	switch choice {
	case "1":
		listTodos(todoSlice, true)
	case "2":
		addTodo(&todoSlice)
	case "3":
		editTodo(&todoSlice)
	case "4":
		deleteTodo(&todoSlice)
	case "5":
		fmt.Println("Quiting...")
		return
	default:
		fmt.Printf("'%v' is not a valid option.\n", choice)
		options()
	}

}

// removeIndex(s []Todo, index int8) []Todo should only be called inside an option.
func removeIndex(s []Todo, index int8) []Todo {
	return append(s[:index], s[index+1:]...)
}

func deleteTodo(ts *[]Todo) {
	defer options()

	// Make sure the user has some todos.
	if len(*ts) > 0 {
		listTodos(*ts, false)
		var choice string
		fmt.Println("Which todo number do you want to delete?")
		fmt.Scanln(&choice)
		index, err := strconv.ParseInt(choice, 10, 8)
		if err != nil {
			fmt.Println("Enter a number.")
			return
		}
		*ts = removeIndex(*ts, int8(index-1))
	} else {
		fmt.Println("You have no todos. Add some.")
		return
	}
}

func editTodo(ts *[]Todo) {
	defer options()

	// Make sure the user has some todos.
	if len(*ts) > 0 {
		listTodos(*ts, false)
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
		fmt.Printf("Edited '%v' to be '%v'.\n", (*ts)[index-1].desc, newText)
		(*ts)[index-1].desc = newText
		(*ts)[index-1].edited = true
		(*ts)[index-1].timeEdited = time.Now().Format(time.UnixDate)
	} else {
		fmt.Println("You have no todos. Add some.")
		return
	}

}

func listTodos(slice []Todo, isDefer bool /* isDefer should be false when this function is being called inside an option. */) {

	/*
		listTodos(slice []Todo, isDefer bool) is called inside of two functions, which are already defering options().
		This if statement allows us to make sure options() isn't called when it isn't wanted through the isDefer paramter.
	*/
	if isDefer {
		defer options()
	}

	// Make sure the user has some todos.
	if len(slice) > 0 {
		fmt.Println("Your todos are:")
		fmt.Print("\n")
		for i := 0; i < len(slice); i++ {
			fmt.Print("TODO NUMBER ", i+1, ":\n")
			fmt.Println("Description:", slice[i].desc)
			fmt.Println("Time published:", slice[i].timePub)
			if slice[i].edited {
				fmt.Println("Time edited:", slice[i].timeEdited)
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("You have no todos. Add some.")
		return
	}
}

func addTodo(ts *[]Todo) {
	defer fmt.Print("\n")
	defer options()
	fmt.Print("Enter a new todo description: ")
	descScanner := bufio.NewReader(os.Stdin)
	newDesc, err := descScanner.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	newDesc = strings.TrimSuffix(newDesc, "\n")
	newTodo := Todo{newDesc, time.Now().Format(time.UnixDate), "", false}
	*ts = append(*ts, newTodo)
	fmt.Printf("Successfully added '%v' to your todos.\n", newTodo.desc)
}

func main() {
	asciiArt :=
		`
====================================================
_____  _____ ______  _____   _____  _      _____ 
|_   _||  _  ||  _  \|  _  | /  __ \| |    |_   _|
  | |  | | | || | | || | | | | /  \/| |      | |  
  | |  | | | || | | || | | | | |    | |      | |  
  | |  \ \_/ /| |/ / \ \_/ / | \__/\| |____ _| |_ 
  \_/   \___/ |___/   \___/   \____/\_____/ \___/ 
                                                                                                                                                                                                                                        	  
====================================================
`
	fmt.Println(asciiArt)
	fmt.Println("Welcome to the todo list CLI!")
	options()
}
