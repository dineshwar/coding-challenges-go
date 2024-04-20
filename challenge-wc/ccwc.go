package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  argsWithProg := os.Args
  args := argsWithProg[1:]
  file, err := os.Stat(args[1])
  if err != nil {
    fmt.Println("File does not exist")
  }
  if args[0] == "-c" { 
        fmt.Println(file.Size(), args[1])
  } else if (args[0] == "-l") {
    file, err := os.Open(args[1])
    if err != nil {
		  fmt.Println("Error:", err)
		  return
	  }
	  defer file.Close()

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
  } else if (args[0] == "-w") {

  }
}
