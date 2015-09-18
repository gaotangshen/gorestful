package model

import (
	"fmt"
	"strconv"
)

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoGet() Todos {
	return todos
}
func RepoFindTodo(id string) Todo {
	todoid, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return Todo{}
	}
	for _, t := range todos {
<<<<<<< HEAD
		if int(t.Id) == int(todoid) {
=======
		if int64(t.Id) == todoid {
>>>>>>> 27fd1295805bf04e58c439a17410db32abe191ba
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
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
