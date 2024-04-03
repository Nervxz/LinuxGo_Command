package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLs(t *testing.T) {
	// Create a temporary directory
	testDir := t.TempDir()

	// Create some files
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, file := range files {
		f, err := os.Create(filepath.Join(testDir, file))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
	}

	// Change current directory to the test directory
	if err := os.Chdir(testDir); err != nil {
		t.Fatal(err)
	}
	// Ensure to check and handle the error when deferring the change back to the original directory
	defer func() {
		if err := os.Chdir(".."); err != nil {
			t.Fatal(err)
		}
	}()

	// Capture output
	var out bytes.Buffer
	err := ls(&out)
	if err != nil {
		t.Fatal(err)
	}

	// Check if all files are listed
	for _, file := range files {
		if !strings.Contains(out.String(), file) {
			t.Errorf("ls did not list file: %s", file)
		}
	}
}

func TestPwd(t *testing.T) {
	// Create a temporary directory
	testDir := t.TempDir()

	// Change current directory to the test directory
	if err := os.Chdir(testDir); err != nil {
		t.Fatal(err)
	}
	// Ensure to check and handle the error when deferring the change back to the original directory
	defer func() {
		if err := os.Chdir(".."); err != nil {
			t.Fatal(err)
		}
	}()

	// Capture output
	var out bytes.Buffer
	err := pwd(&out)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the current directory is the expected test directory
	expectedDir := testDir
	if out.String() != expectedDir+"\n" {
		t.Errorf("pwd did not return the correct directory. Expected: %s, Got: %s", expectedDir, out.String())
	}
}

func TestCd(t *testing.T) {
	// Create a temporary directory
	testDir := t.TempDir()

	// Create a directory within the temporary directory
	if err := os.Mkdir(filepath.Join(testDir, "test_dir"), 0755); err != nil {
		t.Fatal(err)
	}

	// Change current directory to the test directory
	if err := os.Chdir(filepath.Join(testDir, "test_dir")); err != nil {
		t.Fatal(err)
	}
	// Ensure to check and handle the error when deferring the change back to the original directory
	defer func() {
		if err := os.Chdir(".."); err != nil {
			t.Fatal(err)
		}
	}()

	// Capture output
	var out bytes.Buffer
	err := pwd(&out)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the current directory is the expected test directory
	expectedDir := filepath.Join(testDir, "test_dir")
	if out.String() != expectedDir+"\n" {
		t.Errorf("cd did not change directory properly. Expected: %s, Got: %s", expectedDir, out.String())
	}
}

func TestFind(t *testing.T) {
	// Create a temporary directory
	testDir := t.TempDir()

	// Create some files
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, file := range files {
		f, err := os.Create(filepath.Join(testDir, file))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
	}

	// Capture output
	var out bytes.Buffer
	err := find([]string{testDir, "file2"}, &out)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the expected file is found
	expectedFilePath := filepath.Join(testDir, "file2.txt")
	if !strings.Contains(out.String(), expectedFilePath) {
		t.Errorf("find did not find the expected file: %s", expectedFilePath)
	}
}

func TestInvalidCommand(t *testing.T) {
	// Capture output
	var out bytes.Buffer
	err := inputCommand("invalid_command", &out)
	if err == nil {
		t.Error("Expected error for invalid command, but got nil")
	}
}

func TestNonExistentDirectoryCd(t *testing.T) {
	// Change to a non-existent directory
	err := cd([]string{"non_existent_dir"})
	if err == nil {
		t.Error("Expected error for non-existent directory, but got nil")
	}
}

/*
go test
go test -cover
go test -cover -coverprofile=coverage.prof
go tool cover -html=coverage.prof -o coverage.html
*/
