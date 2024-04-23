package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFileSize(fileName string) int64 {
	file, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("File does not exist")
	}
	return file.Size()
}

func getLineCount(fileName os.File) int {
	lineCount := 0
	scanner := bufio.NewScanner(&fileName)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return lineCount
}

func getWordCount(fileName os.File) int {
	wordCount := 0
	scanner := bufio.NewScanner(&fileName)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		for _, w := range line {
			if w != "" {
				wordCount++
			}
		}
	}
	return wordCount
}

func main() {
	argsWithProg := os.Args
	args := argsWithProg[1:]
	if args[0] == "-c" {
		fmt.Println(getFileSize(args[1]), args[1])
	} else if args[0] == "-l" {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		fmt.Println(getLineCount(*file), args[1])
	} else if args[0] == "-w" {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		fmt.Println(getWordCount(*file), args[1])
	} else if args[0] == "-m" {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		charCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Add the number of characters in the line to the count
			// fmt.Println(scanner.Text())
			charCount += len(scanner.Text())
		}

		fmt.Println(charCount, args[1])
	} else {
		if strings.Contains(args[0], ".") {
			file, err := os.Open(args[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer file.Close()
			fmt.Println(getLineCount(*file), getFileSize(args[0]), getWordCount(*file), args[0])
		}
	}
}
