package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
			fmt.Printf("%s has been writed\n", target)
		}
	case 2:
		read(scanner, target)
	case 3:
		writeNew(target)
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

	errC := ioutil.WriteFile(file, []byte(" "), 0644)
	isError(errC)

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
	if exists(file) {
		if confirmation("File exist, overwrite ", file) {
			createFile(file)
			return true
		} else {
			return false
		}
	} else {
		createFile(file)
		return true
	}
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

	for i, v := range idx {
		if i != len(idx) {
			fmt.Fprintln(f, v)
		} else {
			fmt.Fprint(f, v)
		}
		isError(err)
	}
}

func writeNew(target string) {
	f, err := os.OpenFile(target, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	isError(err)
	defer f.Close()
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	afterEdit, err := ui.Ask("new line", &input.Options{})

	isError(err)

	_, ers := f.Write([]byte("\n" + afterEdit))

	isError(ers)

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
	help := "\n---------------\ncommand :\n1. new    - create new file\n2. edit   - write or edit on file\n3. write  - write new line on file\n4. read   - view file\n5. delete - delete file\n\n./goText <file name> <command>"
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
		if newFile(command[1]) {
			fmt.Printf("%s has been created\n", command[1])
		}
	case "edit":
		rw(command[1], 1)
	case "write":
		rw(command[1], 3)
	case "read":
		rw(command[1], 2)
	case "delete":
		if exists(command[1]) {
			if confirmation("Delete ", command[1]) {
				if !(deleteFile(command[1])) {
					fmt.Printf("%s has been deleted\n", command[1])
				}
			}
		} else {
			fmt.Printf("%s not found\n", command[1])
		}
	default:
		invalid("\nUnknown command :" + command[2])
	}
}
