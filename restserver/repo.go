package main

import "fmt"

var currentId int
var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Write Presentation"})
	RepoCreateTodo(Todo{Name: "Host Meetup"})
}

func RepoFindToDo(id int) Todo {
	for _, td := range todos {
		if td.Id == id {
			return td
		}
	}
	return Todo{}
}

func RepoCreateTodo(todo Todo) Todo {
	currentId += 1
	todo.Id = currentId
	todos = append(todos, todo)
	return todo
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
