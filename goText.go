package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tcnksm/go-input"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err == nil)
}

func rw(target string, argu int) error {
	file, err := os.OpenFile(target, os.O_RDWR, 0644)

	if !(isError(err)) {
		return nil
	} else {
		fmt.Println("Opening ", target)
		fmt.Printf("================\n\n")
	}

	scanner := bufio.NewScanner(file)

	switch argu {
	case 1:
		write(scanner, target)
	case 2:
		read(scanner, target)
	}

	defer file.Close()

	return nil
}

func newFile(file string) bool {
	_, err := os.Create(file)

	if isError(err) {
		return false
	}

	return true
}

func write(file *bufio.Scanner, target string) {
	var idx []string
	f, err := os.Create("." + target + ".temp")
	i := 1

	for file.Scan() {
		fmt.Println("line ", i)
		idx = append(idx, edit(file.Text()))
		i++
	}

	for _, v := range idx {
		fmt.Fprintln(f, v)
		isError(err)
	}
}

func read(text *bufio.Scanner, target string) {
	for text.Scan() {
		fmt.Println(text.Text())
	}

	isError(text.Err())
}

func deleteFile(file string) bool {
	err := os.Remove(file)

	if isError(err) {
		return false
	}

	return true
}

func edit(line string) string {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	afterEdit, err := ui.Ask("", &input.Options{
		Default: ": " + line,
	})

	isError(err)

	return afterEdit
}

func invalid(inv string) {
	help := ", use \n---------------\n1. new    - create new file\n2. edit  - write or edit on file\n3. read   - view file\n4. delete - delete file\n\n/.goText <file name> <command>"
	fmt.Printf(inv + help)
}

func main() {
	command := os.Args

	if len(command) < 2 {
		invalid("Need Argument")
		return
	}

	switch command[2] {
	case "new":
		if !(newFile(command[1])) {
			fmt.Printf("%s has been created", command[1])
		}
	case "edit":
		rw(command[1], 1)
	case "read":
		rw(command[1], 2)
	case "delete":
		if !(deleteFile(command[1])) {
			fmt.Printf("%s has been deleted", command[1])
		}
	default:
		invalid("Incomplete command")
	}
}
