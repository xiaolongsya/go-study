package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Error creating file")
	}
	defer file.Close()

	time.Sleep(3 * time.Second)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// file.Close()
		return errors.New("Error encoding JSON")
	}
	// file.Close()
	return nil
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		// file.Close()
		return nil, err
	}
	return lines, nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
