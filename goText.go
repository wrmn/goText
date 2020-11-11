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
		delFile, renFile := os.Remove(target), os.Rename("."+target+".temp", target)
		if isError(renFile) && isError(delFile) {
			fmt.Printf("%s has been writed", target)
		}
	case 2:
		read(scanner, target)
	}

	defer file.Close()

	return nil
}

func confirmation(msg, file string) bool {
	var input string
	fmt.Print(msg + file + " ? (y/N)")
	fmt.Scanln(&input)
	if input == "y" || input == "Y" {
		return true
	}
	return false
}

func createFile(file string) {
	_, err := os.Create(file)
	isError(err)
}

func exists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func newFile(file string) bool {
	fmt.Printf("%t", exists(file))
	if exists(file) {
		if confirmation("File exist, overwrite ", file) {
			createFile(file)
		}
	} else {
		createFile(file)
	}
	return false
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
	help := "\n---------------\ncommand :\n1. new    - create new file\n2. edit   - write or edit on file\n3. read   - view file\n4. delete - delete file\n\n./goText <file name> <command>"
	fmt.Printf(inv + help)
}

func main() {
	command := os.Args

	if len(command) < 2 {
		invalid("\nNeed Argument")
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
		if confirmation("Delete ", command[1]) {
			if !(deleteFile(command[1])) {
				fmt.Printf("%s has been deleted", command[1])
			}
		}
	default:
		invalid("\nUnknown command :" + command[2])
	}
}
