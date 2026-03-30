package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func WriteJson(path string, data any) {
	os.Create(path)

}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return nil, err
	}
	return lines, nil
}
