package library

import (
	"github.com/taubyte/tau-cli/cli/common"
)

func (l link) Push() common.Command {
	return l.cmd.PushCmd()
}
