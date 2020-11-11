package main

import (
	"bufio"
	"fmt"
	"io"
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

func write(file *bufio.Scanner, target string) error {
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

	return nil
}

func read(text *bufio.Scanner, target string) error {

	for text.Scan() {
		fmt.Println(text.Text())
	}

	isError(text.Err())

	return nil

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

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func main() {
	command := os.Args

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
		fmt.Printf("Invalid command, use \n---------------\n1. new    - create new file\n2. edit  - write or edit on file\n3. read   - view file\n4. delete - delete file\n")

	}
}
