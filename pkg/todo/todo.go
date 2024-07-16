package todo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sharon-xa/todo/pkg/file"
	"github.com/sharon-xa/todo/pkg/pprint"
)

type Todo struct {
	FilePath string
}

func Init() *Todo {
	t := new(Todo)

	configDir := filepath.Join(os.Getenv("HOME"), ".config", "todo")
	if !file.DirectoryExists(configDir) {
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			pprint.Perror(fmt.Sprintf("Can't create config directory, %s", err))
			return nil
		}
	}

	t.FilePath = filepath.Join(configDir, "todo.txt")
	if !file.FileExists(t.FilePath) {
		file, err := os.OpenFile(t.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		defer file.Close()
		if err != nil {
			pprint.Perror(fmt.Sprintf("Can't create todo.txt file, %s", err))
			return nil
		}
	}

	return t
}

func (t *Todo) OpenFileWithDefaultEditor() {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		pprint.Perror("Can't open editor [no $EDITOR env]")
		return
	}

	cmd := exec.Command(editor, t.FilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		pprint.Perror("Failed to open editor")
		pprint.Perror(err.Error())
	}
}
