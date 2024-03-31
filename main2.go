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

// implement command
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
		return fmt.Errorf("Please Reinput Command")
	}
}

// ls command
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

// pwd command
func pwd(out io.Writer) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, pwd)
	return nil
}

// cd command
func cd(args []string) error {
	if len(args) < 1 {
		return errors.New("cd: path")
	}
	return os.Chdir(args[0])
}

// find command
func find(args []string, out io.Writer) error {
	if len(args) < 2 {
		return errors.New("Reinput: find <path> <expression>")
	}
	root := args[0]
	expression := args[1]

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Base(path) == expression || filepath.Ext(path) == expression {
			fmt.Fprintln(out, path)
			return filepath.SkipDir
		}

		return nil
	})
}

func match(path string, expression string) bool {
	if expression == "*" {
		// Match any file or directory
		return true
	}

	if strings.HasPrefix(expression, ".") {
		// Match extension
		return filepath.Ext(path) == expression
	}

	// Match filename
	return filepath.Base(path) == expression
}

/*
 use find + directory + * to search all file in directory ex : find /home/ *
 use find + directory + extension ex :  find home/nervx .txt
 use find + directory + name of search file ex : find /home/nervx main2.go

*/
