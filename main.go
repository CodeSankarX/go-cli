package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Todo application written in GO 🚀")
	fmt.Print("todo add delete, ls, \n")
	fmt.Println("")

	_, err := os.Stat("todo.json")
	if err != nil {
		file, err := os.Create("todo.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
	}

	for {
		fmt.Print("todo>")
		var action string
		var change string
		_, err := fmt.Scan(&action)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		switch action {
		case "add":
			fmt.Scan(&change)
			addTodo(change)
			fmt.Print("Task added !")
		case "rm":
			fmt.Scan(&change)
			deleteTodo(change)
			fmt.Print("Task removed !")
		case "ls":
			listTodos()
		case "done":
			fmt.Scan(&change)
			update(change)
			fmt.Print("Task marked as done !")
		default:
			fmt.Println("Invalid action. Please use 1 for add, 2 for delete, 3 for list.")
		}
		fmt.Println()
	}

}
