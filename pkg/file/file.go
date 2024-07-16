package file

import (
	"fmt"
	"os"

	"github.com/sharon-xa/todo/pkg/pprint"
)

func LoadFileToRead(filePath string) *os.File {
	file, err := os.Open(filePath) // Open the file in read-only mode
	if err != nil {
		pprint.Perror(err.Error())
	}
	return file
}

func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func AppendStringToFile(filePath string, content string) error {
	// Open the file with append-only, create if not exists, with permissions 0644
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		pprint.Perror(fmt.Sprintf("Can't write in file, %s", err.Error()))
		return err
	}

	return nil
}

func ReplaceFileContent(filePath string, content string) error {
	// Open the file with write-only, create if not exists, truncate if exists, with permissions 0644
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
