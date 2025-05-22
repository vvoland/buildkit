package instructions

import (
	"fmt"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/instructions/custom"
)

func tryMagicRun(cmd string, req parseRequest) (Command, error) {
	shortcut, ok := custom.Runs[strings.ToUpper(cmd)]
	if !ok {
		return nil, nil
	}

	req.command = "RUN"
	req.args = []string{
		fmt.Sprintf(shortcut.CmdFmt, strings.Join(req.args, " ")),
	}
	if req.attributes != nil {
		req.attributes["json"] = false
	}
	req.flags.Args = shortcut.Flags

	return parseRun(req)
}
