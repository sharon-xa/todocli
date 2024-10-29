package todo

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/sharon-xa/todo/pkg/file"
	"github.com/sharon-xa/todo/pkg/pprint"
)

func (t *Todo) PrintTasks() {
	file := file.LoadFileToRead(t.FilePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	color.Set(color.FgMagenta)
	fmt.Print("")
	color.Set(color.BgMagenta, color.FgBlack, color.Bold)
	fmt.Print("      ToDo List      ")
	color.Set(color.Reset)
	color.Set(color.FgMagenta)
	fmt.Print("")
	fmt.Print("\n\n")
	color.Set(color.Reset)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		color.Set(color.FgWhite, color.Faint)
		fmt.Printf("%2d  ", i)
		color.Set(color.Reset)
		if strings.Contains(line, "[X]") {
			color.Set(color.FgGreen)
			fmt.Print("  ")
			color.Set(color.Reset)
			color.Set(color.Bold, color.CrossedOut)
			fmt.Println(line[4:])
			color.Set(color.Reset)
		} else {
			fmt.Print("  ")
			color.Set(color.Bold)
			fmt.Println(line[3:])
			color.Set(color.Reset)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		pprint.Perror(fmt.Sprintf("Can't read file, %s", scanner.Err().Error()))
	}
}

func (t *Todo) RemoveTask(taskId int) {
	newLines, err := t.getFileContentMap()
	if err != nil {
		pprint.Perror(err.Error())
		return
	}

	if len(*newLines) == 0 {
		pprint.Pdone("No Tasks to remove.")
		return
	}

	delete(*newLines, taskId)

	var newFileContent strings.Builder
	for _, line := range *newLines {
		newFileContent.WriteString(line)
		newFileContent.WriteString("\n")
	}

	err = file.ReplaceFileContent(t.FilePath, newFileContent.String())
	if err != nil {
		pprint.Perror(fmt.Sprintf("couldn't write to the file, %s", err.Error()))
		return
	}

	t.PrintTasks()
}

func (t *Todo) AddTask(task string) {
	ftask := fmt.Sprintf("[] %s\n", task)
	err := file.AppendStringToFile(t.FilePath, ftask)
	if err != nil {
		pprint.Perror(fmt.Sprintf("couldn't append text to file, %s", err.Error()))
		return
	}
	t.PrintTasks()
}

func (t *Todo) ToggleTask(taskId int) {
	newLines, err := t.getFileContentMap()
	if err != nil {
		pprint.Perror(err.Error())
		return
	}

	if len(*newLines) == 0 {
		pprint.Perror("No Tasks to toggle.")
		return
	}

	if strings.Contains((*newLines)[taskId], "[X]") {
		(*newLines)[taskId] = strings.Replace((*newLines)[taskId], "[X", "[", 1)
	} else if strings.Contains((*newLines)[taskId], "[]") {
		(*newLines)[taskId] = strings.Replace((*newLines)[taskId], "[]", "[X]", 1)
	}

	var newFileContent strings.Builder
	for _, line := range *newLines {
		newFileContent.WriteString(line)
		newFileContent.WriteString("\n")
	}

	err = file.ReplaceFileContent(t.FilePath, newFileContent.String())
	if err != nil {
		pprint.Perror(fmt.Sprintf("couldn't write to the file, %s", err.Error()))
		return
	}

	t.PrintTasks()
}

func (t *Todo) getFileContentMap() (*map[int]string, error) {
	f := file.LoadFileToRead(t.FilePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	newLines := make(map[int]string, 0)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		newLines[i] = line
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Can't read file")
	}

	return &newLines, nil
}
