package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Input your command: ")
	repl(os.Stdin, os.Stdout)

}

// Read input
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

// Implement command
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
	case "find":
		return find(args[1:], out)

	default:
		return errors.New("please re_input command") // replace fmt.Errorf with errors.New
	}
}

// Ls command
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

// Pwd command
func pwd(out io.Writer) error {

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, pwd)
	return nil
}

// Cd command
func cd(args []string) error {
	if len(args) < 1 {
		return errors.New("cd: path")
	}
	return os.Chdir(args[0])
}

/*
 use find + directory + * to search all file in directory ex : find /home/ *
// use find + directory + extension ex :  find home/example .txt
// use find + directory + name of search file ex : find /home/example main2.go

*/
// Find command
func find(args []string, out io.Writer) error {
	if len(args) < 2 {
		return errors.New("re-enter: find <path> <expression>")
	}

	root := args[0]
	expression := args[1]

	return filepath.Walk(root, func(path string, _ os.FileInfo, err error) error { // rename info os.FileInfo to _ due to --presets=linter lint
		if err != nil {
			return err
		}

		// Check if filename or path
		if strings.Contains(strings.ToLower(path), strings.ToLower(expression)) {
			fmt.Fprintln(out, path)
		}

		return nil
	})
}
