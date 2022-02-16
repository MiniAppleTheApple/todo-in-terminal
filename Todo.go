package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Todo struct {
	list  []*Task
	input string
}

func NewTodo() Todo {
	return Todo{
		list:  []*Task{},
		input: "",
	}
}
func (todo *Todo) NewTask(name string) *Task {
	task := NewTask(name)
	todo.list = append(todo.list, task)
	return task
}
func (todo *Todo) InsertTask(index int, name string) {
	task := NewTask(name)
	if len(todo.list) == index { // nil or empty slice or after last element
		todo.list = append(todo.list, task)
	}
	todo.list = append(todo.list[:index+1], todo.list[index:]...) // index < len(a)
	todo.list[index] = task
}

func (todo *Todo) RemoveTask(index int) {
	if index >= len(todo.list) {
		return
	}
	if len(todo.list) < 1 {
		todo.list = []*Task{}
		return
	}
	todo.list = append(todo.list[:index], todo.list[index+1:]...)
}
func (todo *Todo) Save() {
	data, err := json.Marshal(todo.list)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func (todo *Todo) Render() {
	todo.input = CLIInput()
	data, _ := ioutil.ReadFile("data.json")
	json.Unmarshal(data, &todo.list)
	ResetStyle()
	switch todo.input {
	case "help":
		PrintBold(Yellow("- new: "))
		PrintUnderline("add a new task into the list\n")
		PrintBold(Yellow("- del: "))
		PrintUnderline("delete the task by put in the position\n")
		PrintBold(Yellow("- show: "))
		PrintUnderline("display all the task\n")
		PrintBold(Yellow("- do: "))
		PrintUnderline("making the undone task into done task or making the done task into undone task\n")
		PrintBold(Yellow("- quit: "))
		PrintUnderline("quit the program\n")
	case "new":
		todo.input = GetInput("task name: ")
		todo.NewTask(todo.input)
		fmt.Print(CHECKMARK)
		msg := fmt.Sprintf(" sucess to add new '%v' task from the list\n", todo.input)
		PrintSucess(msg)
		ResetStyle()
	case "del":
		todo.input = GetInput("position of the task: ")
		position, _ := strconv.Atoi(todo.input)
		if len(todo.list) > position {
			msg := fmt.Sprintf("%v sucess to delete '%v' task from the list\n", CHECKMARK, todo.list[position].GetName())
			todo.RemoveTask(position)
			PrintSucess(msg)
			ResetStyle()
		} else {
			fmt.Print(CROSSMARK)
			msg := fmt.Sprintf(" cannot delete the task,because the task doesn't exist\n")
			PrintWarning(msg)
		}

	case "show":
		fmt.Println("--- Tasks ---")
		if len(todo.list) < 1 {
			fmt.Println("nothing here")
		}
		for i, v := range todo.list {
			symbol := func() string {
				if v.IsDone() {
					return CHECKMARK
				}
				return CROSSMARK
			}()
			msg := fmt.Sprintf("%v %v. %v\n", symbol, i, v.GetName())
			PrintSucess(msg)
		}
		fmt.Println("-------------")
	case "do":
		todo.input = GetInput("position of the task: ")
		position, _ := strconv.Atoi(todo.input)
		task := todo.list[position]
		if task.IsDone() {
			task.ReDo()
			fmt.Print(CHECKMARK)
			msg := fmt.Sprintf(" sucess to change the '%v' task into undone\n", task.GetName())
			PrintSucess(msg)
		} else {
			task.Done()
			fmt.Print(CHECKMARK)
			msg := fmt.Sprintf(" sucess to change the '%v' task into done\n", task.GetName())
			PrintSucess(msg)
		}
		ResetStyle()
	case "quit":
		fmt.Print(CHECKMARK)
		PrintSucess(" quitted sucessfully\n")
		return
	default:
		fmt.Print(CROSSMARK)
		PrintWarning(" invalid command\n")
	}
	todo.Save()
	todo.Render()
}
