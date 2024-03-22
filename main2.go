package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
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
		if err := inputCommand(line, out); err != nil {
			fmt.Fprintln(out, err)
		}
	}
}

// ls,cd,pwd,etc...
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
