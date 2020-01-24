package directory

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {

	file, err := os.Open("../testdata/grades.csv")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
	}

	content, err := readFile(file)
	if err != nil {
		t.Error("Failed to read csv data")
	}
	fmt.Print(content)
}

func TestOpenAndReadFile(t *testing.T) {
	file := "../testdata/grades.csv"

	content, err := OpenAndReadFile(file)
	if err != nil {
		t.Error("Failed to read csv data")
	}

	fmt.Print(content)
}