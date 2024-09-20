package codes

type NothingCode struct {
  Line string
  Index int
  Success bool
}

func (p *NothingCode) Run() bool {
  return p.Success
}

