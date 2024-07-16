package pprint

import (
	"fmt"

	"github.com/fatih/color"
)

func Perror(msg string) {
	color.Set(color.FgRed, color.Bold)
	fmt.Print("Error: ")
	color.Set(color.Reset)
	fmt.Println(msg)
}

// make this function take color and
func Pprint(msg string, c color.Attribute) {
	color.Set(c, color.Bold)
	fmt.Print(msg)
	color.Set(color.Reset)
}

func Pdone(msg string) {
	color.Set(color.FgGreen, color.Bold)
	fmt.Print("Done: ")
	color.Set(color.Reset)
	println(msg)
}
