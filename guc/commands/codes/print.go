package codes

import "fmt"

type PrintCode struct {
  Line string
  Index int
}

func (p *PrintCode) Run() bool {
  fmt.Println(p.Line[6:])
  return true
}

