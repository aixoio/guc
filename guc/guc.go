package guc

import (
	"regexp"

	"github.com/aixoio/guc/guc/commands"
	"github.com/aixoio/guc/guc/commands/codes"
)

func GetCommand(line string, index int) commands.Command {
  if regexp.MustCompile("^PRINT ").Match([]byte(line)) {
    return &codes.PrintCode{
      Line: line,
      Index: index,
    }
  }

  

  return &codes.NothingCode{
    Success: true,
  }
}

