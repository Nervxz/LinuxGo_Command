package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input your command: ")
	repl(os.Stdin, os.Stdout)
}

// read input
func repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	prompt := ">"

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "" {
			continue
		}
		if err := inputCommand(line, out); err != nil {
			fmt.Fprintln(out, err)
		}
	}
}

// ls,cd,pwd

func inputCommand(line string, out io.Writer) error {

	if line == "exit" || line == "quit" {
		os.Exit(0)
	}

	args := strings.Fields(line)
	switch args[0] {
	case "ls":
		return ls(out)
	case "pwd":
		return pwd(out)
	case "cd":
		return cd(args[1:])
	default:
		return fmt.Errorf("Please Reinput Command")
	}
}

func ls(out io.Writer) error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintln(out, file.Name())
	}
	return nil
}

func pwd(out io.Writer) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, pwd)
	return nil
}

func cd(args []string) error {
	if len(args) < 1 {
		return errors.New("cd: path")
	}
	return os.Chdir(args[0])
}

/*
func inputCommand(line string, out io.Writer) error {

	if line == "exit" || line == "quit" {
		os.Exit(0)
	}

	args := strings.Fields(line)
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("cd: path")
		}
		return os.Chdir(args[1])

	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = out

		return cmd.Run()
	}
}
*/
