package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_file_size(fileName string) int64 {
	file, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("File does not exist")
	}
	return file.Size()
}

func main() {
	argsWithProg := os.Args
	args := argsWithProg[1:]
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	if args[0] == "-c" {
		fmt.Println(get_file_size(args[1]), args[1])
	} else if args[0] == "-l" {
		lineCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineCount++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(lineCount, args[1])
	} else if args[0] == "-w" {
		wordCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.Fields(scanner.Text())
			for _, w := range line {
				if w != "" {
					wordCount++
				}
			}
		}
		fmt.Println(wordCount, args[1])
	} else if args[0] == "-m" {
		charCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Add the number of characters in the line to the count
			// fmt.Println(scanner.Text())
			charCount += len(scanner.Text())
		}

		fmt.Println(charCount, args[1])
	}
}
