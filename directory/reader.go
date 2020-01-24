package directory

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func OpenAndReadFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s", fileName)
	}

	lines, err := readFile(file)
	if err != nil {
		fmt.Printf("Failed to read file: %s", fileName)
	}
	return lines, err
}

func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return lines, err
}