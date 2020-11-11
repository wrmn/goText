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
	return (err != nil)
}

func rw(target string, argu int) error {
	file, err := os.OpenFile(target, os.O_RDWR, 0644)

	if isError(err) {
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

func newFile(file string) error {
	return nil
}

func write(file *bufio.Scanner, target string) error {
	var idx []string

	f, err := os.Create(target + ".temp")

	for file.Scan() {
		idx = append(idx, edit(file.Text()))
	}
	for _, v := range idx {
		fmt.Fprintln(f, v)
		if isError(err) {
			return err
		}
	}
	return nil
}

func read(text *bufio.Scanner, target string) error {

	for text.Scan() {
		fmt.Println(text.Text())
	}

	if isError(text.Err()) {
		return nil
	}

	return nil

}

func deleteFile(file string) error {
	return nil
}

func edit(line string) string {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	afterEdit, err := ui.Ask("", &input.Options{
		Default: ": " + line,
	})

	if isError(err) {
		return ""
	}

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

	switch command[1] {
	case "new":
		newFile(command[2])
	case "write":
		rw(command[2], 1)
	case "read":
		rw(command[2], 2)
	case "delete":
		deleteFile(command[2])
	default:
		fmt.Print("Invalid command")
	}
}
