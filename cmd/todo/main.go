package main

import (
	"flag"

	"github.com/sharon-xa/todo/pkg/todo"
)

func main() {
	t := todo.Init()

	newTaskPtr := flag.String("a", "", "Add a task")
	removeTaskNumPtr := flag.Int("r", 0, "Remove a task")
	toggledTaskNumPtr := flag.Int("t", 0, "Toggle done for a task")
	editFilePtr := flag.Bool("e", false, "Edit todo file")

	flag.Parse()

	newTask := *newTaskPtr
	removeTaskNum := *removeTaskNumPtr
	toggledTaskNum := *toggledTaskNumPtr
	editFile := *editFilePtr

	switch {
	case removeTaskNum > 0:
		t.RemoveTask(removeTaskNum)

	case newTask != "":
		t.AddTask(newTask)

	case toggledTaskNum > 0:
		t.ToggleTask(toggledTaskNum)

	case editFile:
		t.OpenFileWithDefaultEditor()

	default:
		t.PrintTasks()
	}
}
