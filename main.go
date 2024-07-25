package main

import (
	"fmt"
	"time"
)

type todo struct {
	ID    int
	Title string
	Done  bool
	Date  time.Time
}

type todoList []todo

func main() {
	fmt.Println("Hello, Welcome to my todo!")
	todoListItems := todoList{}
	var command string
	var id int
	for {
		fmt.Println("Enter a command:")
		fmt.Println("1. Add a todo")
		fmt.Println("2. Delete a todo")
		fmt.Println("3. List all todos")
		fmt.Println("4. Update a todo")
		fmt.Println("5. Exit")
		fmt.Scanln(&command)
		switch command {
		case "1":
			fmt.Println("Please enter todo")
			todoListItems = addTodo(todo{ID: 1, Title: "Drink water", Done: false, Date: time.Now()}, todoListItems)
			fmt.Println("Todo added")
		case "2":
			fmt.Println("Please enter todo id")
			fmt.Scanln(&id)
			todoListItems = deleteTodo(1, todoListItems)
			fmt.Println("Todo deleted")
		case "3":
			printTodoList(todoListItems)
		case "4":
			fmt.Println("Please enter todo id")
			fmt.Scanln(&id)
			updateTodo(id, todoListItems)
			fmt.Println("Todo updated")
		default:
			fmt.Println("Invalid command")
		}
	}
}

func addTodo(todoitem todo, todoListItems todoList) todoList {
	todoListItems = append(todoListItems, todoitem)
	return todoListItems
}

func deleteTodo(id int, todoListItems todoList) todoList {
	for i, v := range todoListItems {
		if v.ID == id {
			todoListItems = append(todoListItems[:i], todoListItems[i+1:]...)
		}
	}
	return todoListItems
}

func printTodoList(todoListItems todoList) {
	fmt.Println(todoListItems)
}

func updateTodo(id int, todolistItems todoList) {
	if id > len(todolistItems) {
		fmt.Println("Todo not found")
	}
	todolistItems[id].Done = true
}
