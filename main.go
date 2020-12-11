package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/NewbJS/Go-Todo-CLI/ascii"
	"github.com/NewbJS/Go-Todo-CLI/todo-handler/options"
)

func main() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Println(ascii.ASCIITitle)
	fmt.Println("Welcome to the todo list CLI!")
	options.Options()
}
