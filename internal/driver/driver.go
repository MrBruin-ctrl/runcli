package driver

import (
	"context"
	"github.com/c-bata/go-prompt"
)

type CliDriver interface {
	Prompt
	// CliName 客户端名称，用于前排提示
	CliName() string
	SurveyConfig()
	Conn(context.Context) error
	Exit()
}

type Prompt interface {
	LivePrefix() string
	Executor(cmd string)
	Completer(document prompt.Document) []prompt.Suggest
}
