package fileHelper

import (
	"bufio"
	"os"
)

func ReadFile(filePath string) (error, []string) {

	var fileLines []string
	readFile, err := os.Open(filePath)

	if err != nil {
		return err, fileLines
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return nil, fileLines
}
