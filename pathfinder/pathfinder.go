package pathfinder

import (
	"fmt"
	"os"
	"path/filepath"
)

func getFilePath(filename string) (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Construct the file path by joining the current working directory and the filename
	filePath := filepath.Join(cwd, filename)

	return filePath, nil
}

func Fpath() string {
	filename := "jobs.csv"
	filePath, err := getFilePath(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return err.Error()
	}
	return filePath 
}
