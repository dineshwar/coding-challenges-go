package main

import (
  "fmt"
  "os"
)

func main() {
  argsWithProg := os.Args
  args := argsWithProg[1:]
  fmt.Println(args)
}
