package main

import (
	"fmt"

	"github.com/NewbJS/Go-Todo-CLI/ascii"
	"github.com/NewbJS/Go-Todo-CLI/todo-handler/options"
)

func main() {

	fmt.Println(ascii.ASCIITitle)
	fmt.Println("Welcome to the todo list CLI!")
	options.Options()
}
