package main

import (
  "fmt"
  "os"
)

func main() {
  argsWithProg := os.Args
  args := argsWithProg[1:]
  if args[0] == "-c" {
    file, err := os.Stat(args[1])
    if err != nil {
      fmt.Println("File does not exist")
    }
    fmt.Println(file.Size())
  }
}
