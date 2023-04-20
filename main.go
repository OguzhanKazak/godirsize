package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go {directory}")
		os.Exit(1)
	}

	path := os.Args[1]
	err := showTree(path, " ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func showTree(root, indent string) error {
	// Read the root directory
	files, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	// Traverse file/folders
	for i, file := range files {

		path := filepath.Join(root, file.Name())

		// calculate indentations
		var newIndent string
		if i == len(files)-1 {
			newIndent = indent + "└── "
		} else {
			newIndent = indent + "├── "
		}

		fileInfo, err := file.Info()
		if err != nil {
			return err
		}

		sizeString := prepareSizeString(fileInfo.Size())

		// print out the file/folder name and the size
		fmt.Printf("%s%s (%s)\n", newIndent, file.Name(), sizeString)

		// DFS recursive call for deeper file/folders
		if file.IsDir() {
			err := showTree(path, indent+"│   ")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func prepareSizeString(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	unitIndex := 0
	for size > 1024 && unitIndex < len(units)-1 {
		size = size / 1024
		unitIndex++
	}

	return fmt.Sprintf("%d %s", size, units[unitIndex])
}
