package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tl, err := loadTodoListFromDefaultFile()
	if err != nil {
		tl = TodoList{}
	}
	if os.Args[1] == "add" {
		text := strings.Join(os.Args[2:], " ")
		addTodo(&tl, text)
	}
	if os.Args[1] == "list" {
		tl.listPending()
	}
	if os.Args[1] == "done" {
		id, _ := strconv.Atoi(os.Args[2])
		tl.completeTodo(id)
	}
	home := os.Getenv("HOME")
	f, err := os.Create(home + "/.gotodo.txt")
	if err != nil {
		fmt.Println(err)
	}
	json_tl, _ := json.Marshal(tl)
	f.Write(json_tl)
}

func loadTodoListFromDefaultFile() (TodoList, error) {
	home := os.Getenv("HOME")
	contents, err := os.ReadFile(home + "/.gotodo.txt")
	if err != nil {
		os.Create(home + "/.gotodo.txt")
	}
	tl := TodoList{}
	err = json.Unmarshal(contents, &tl)
	if err != nil {
		return TodoList{}, nil
	}
	return tl, nil
}

func addTodo(tl *TodoList, text string) {
	id := len(tl.Todos) + 1
	todo := Todo{Id: id, Text: text, Done: false}
	tl.add(todo)
}

type Todo struct {
	Id   int
	Text string
	Done bool
}

type TodoList struct {
	Todos []Todo `json:"Todos"`
}

func (tl *TodoList) add(todo Todo) {
	if tl.has(todo.Id) {
		return
	}
	tl.Todos = append(tl.Todos, todo)
}

func (tl *TodoList) has(id int) bool {
	for i := 0; i < len(tl.Todos); i++ {
		if tl.Todos[i].Id == id {
			return true
		}
	}
	return false
}

func (tl *TodoList) completeTodo(id int) {
	for i := 0; i < len(tl.Todos); i++ {
		if tl.Todos[i].Id == id {
			tl.Todos[i].Done = true
		}
	}
}

func (tl *TodoList) listPending() {
	for i := 0; i < len(tl.Todos); i++ {
		if !tl.Todos[i].Done {
			fmt.Printf("[%d] %s\n", tl.Todos[i].Id, tl.Todos[i].Text)
		}
	}
}
