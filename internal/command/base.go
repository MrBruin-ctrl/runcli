package command

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/c-bata/go-prompt"
	"runcli/internal/driver"
	"runcli/internal/registry"
)

var CurrentDriver driver.CliDriver

func ChooseCliDriver() {
	if CurrentDriver != nil {
		CurrentDriver.Exit()
		CurrentDriver = nil
	}
	//选择实现的版本
	chooseCli := ""
	names := registry.GetCliNames()
	prompt := &survey.Select{
		Message: "Choose a client:",
		Options: names,
	}
	survey.AskOne(prompt, &chooseCli)
	cliDriver := registry.GetCliByName(chooseCli)
	CurrentDriver = cliDriver
	ctx := context.Background()
	//配置参数
	for {
		CurrentDriver.SurveyConfig()
		err := CurrentDriver.Conn(ctx)
		if err != nil {
			fmt.Println(err.Error())
			isKeep := true
			confirm := &survey.Confirm{
				Message: "do you want retry?",
			}
			survey.AskOne(confirm, &isKeep)
			if !isKeep {
				break
			}
			continue
		}
		break
	}

}

func GetActiveCmdExecutor(cmd string) {
	CurrentDriver.Executor(cmd)
}

func GetActiveCmdCompleter(document prompt.Document) []prompt.Suggest {
	return CurrentDriver.Completer(document)
}

func CommonLivePrefix() (prefix string, useLivePrefix bool) {
	return fmt.Sprint(CurrentDriver.LivePrefix()), true
}
