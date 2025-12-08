package system

import (
	"bufio"
	"os"
	"strings"
)

func OpenFile(filePath string) *os.File {
    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    return file
}

func CloseFile(file *os.File) {
    err := file.Close()
    if err != nil {
        panic(err)
    }
}

func readLineScanner(file *os.File) *bufio.Scanner {
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    return fileScanner
}

func ReadLines(path string) []string {
    var linesBuffer []string

    file := OpenFile(path)
    scanner := readLineScanner(file)

    for scanner.Scan() {
        linesBuffer = append(linesBuffer, scanner.Text())
    }

    CloseFile(file)

    return linesBuffer
}

func ReadAll(path string) string {
	file := OpenFile(path)
	scanner := bufio.NewScanner(file)
	
	str := []string{}

	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	return strings.Join(str, "")
}