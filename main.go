package main

import "github.com/aixoio/guc/guc"

func main() {
  code := guc.GetCommand("PRINT Hello, world!", 1)
  code.Run()
}

