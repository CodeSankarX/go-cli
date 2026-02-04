package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func addTodo(todo string) {
	bytearray, _ := os.ReadFile("todo.json")
	var myTodo []Todo
	var last Todo
	if len(bytearray) > 0 {
		err := json.Unmarshal(bytearray, &myTodo)
		if err != nil {
			fmt.Println("Error while parsing json")
		}
		if len(myTodo) > 1 {
			last = myTodo[len(myTodo)-1]
		}
	}

	insert := Todo{
		Id:        last.Id + 1,
		Task:      todo,
		Completed: false,
	}
	myTodo = append(myTodo, insert)
	writeback(myTodo)

}

func deleteTodo(todo string) {

	var myTodo []Todo
	myTodo = readFile()

	id, _ := strconv.Atoi(todo)
	for i := 0; i < len(myTodo); i++ {
		if id == myTodo[i].Id {
			myTodo = append(myTodo[:i], myTodo[i+1:]...)
		}
	}

	writeback(myTodo)

}

func listTodos() {
	var myTodo []Todo
	myTodo = readFile()

	for _, todo := range myTodo {
		fmt.Printf("ID: %d | Task: %s | Done: %t\n", todo.Id, todo.Task, todo.Completed)
	}

}

func update(todo string) {
	id, _ := strconv.Atoi(todo)
	myTodo := readFile()
	for i := 0; i < len(myTodo); i++ {
		if myTodo[i].Id == id {
			myTodo[i].Completed = true
			break
		}
	}
	writeback(myTodo)
}

func readFile() []Todo {
	bytearray, _ := os.ReadFile("todo.json")
	var myTodo []Todo
	err := json.Unmarshal(bytearray, &myTodo)
	if err != nil {
		fmt.Println("Error while parsing json")
	}
	return myTodo
}

func writeback(myTodo []Todo) {
	arr, err := json.MarshalIndent(myTodo, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// 2. Write to file
	// 0644 is the standard Linux permission (Owner: read/write, Others: read)
	err = os.WriteFile("todo.json", arr, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
