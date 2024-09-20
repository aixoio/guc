package codes

type NothingCode struct {
  Line string
  Index int
}

func (p *NothingCode) Run() bool {
  return true
}

