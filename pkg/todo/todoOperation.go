package todo

import (
	"bufio"
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
	f := file.LoadFileToRead(t.FilePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	newLines := make([]string, 0)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if i != taskId {
			newLines = append(newLines, line)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		pprint.Perror("Can't read file")
	}

	newFileContent := strings.Join(newLines, "\n")
	err := file.ReplaceFileContent(t.FilePath, newFileContent+"\n")
	// I added the "\n" char because the newFileContent is joined by \n
	// and when a save the new file content to the file
	// I get a all the lines normally but the list line doesn't have a new line char
	// so when we append the new task it will be appended to the last line and not to a new line
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
	f := file.LoadFileToRead(t.FilePath)
	scanner := bufio.NewScanner(f)

	newLines := make([]string, 0)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if i == taskId {
			if strings.Contains(line, "[X]") {
				line = strings.Replace(line, "[X", "[", 1)
			} else {
				line = strings.Replace(line, "[", "[X", 1)
			}
		}
		newLines = append(newLines, line)
		i++
	}

	if err := scanner.Err(); err != nil {
		pprint.Perror("Can't read file")
		return
	}

	newFileContent := strings.Join(newLines, "\n")
	err := file.ReplaceFileContent(t.FilePath, newFileContent)
	if err != nil {
		pprint.Perror(fmt.Sprintf("couldn't write to the file, %s", err.Error()))
		return
	}

	t.PrintTasks()
}
