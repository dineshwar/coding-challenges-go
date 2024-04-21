package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		file, err := os.Stat(args[1])
		if err != nil {
			fmt.Println("File does not exist")
		}
		fmt.Println(file.Size(), args[1])
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
		fmt.Println(charCount, args[1])
	}
}
