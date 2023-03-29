package command

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/MrBruin-ctrl/runcli/internal/driver"
	"github.com/MrBruin-ctrl/runcli/internal/registry"
	"github.com/MrBruin-ctrl/runcli/internal/utils"
	"github.com/c-bata/go-prompt"
)

var CurrentDriver driver.CliDriver

func ChooseCliDriver() {
	closeConnectIfExist()
	//选择实现的版本
	chooseCli := ""
	names := registry.GetCliNames()
	prompt := &survey.Select{
		Message: "Choose a client:",
		Options: names,
	}
	utils.AskOne(prompt, &chooseCli)
	cliDriver := registry.GetCliByName(chooseCli)
	CurrentDriver = cliDriver
	cliConnect()
}

func cliConnect() {
	ctx := context.Background()
	//配置参数
	for {
		CurrentDriver.InitConfig()
		CurrentDriver.SurveyConfig()
		err := CurrentDriver.Conn(ctx)
		if err != nil {
			fmt.Println(err.Error())
			isKeep := true
			confirm := &survey.Confirm{
				Message: "do you want retry?",
			}
			utils.AskOne(confirm, &isKeep)
			if !isKeep {
				break
			}
			continue
		}
		break
	}
}

func closeConnectIfExist() {
	if CurrentDriver != nil {
		CurrentDriver.Exit()
		CurrentDriver = nil
	}
}

func GetActiveCmdExecutor(cmd string) {
	CurrentDriver.Executor(cmd)
}

func GetActiveCmdCompleter(document prompt.Document) []prompt.Suggest {
	return CurrentDriver.Completer(document)
}

func GetActiveKeyBind() []prompt.KeyBind {
	return []prompt.KeyBind{
		{
			Key: prompt.ControlR, Fn: func(buffer *prompt.Buffer) {
				//重新执行
				ChooseCliDriver()
			}},
		//{
		//	Key: prompt.ControlM, Fn: func(buffer *prompt.Buffer) {
		//		closeConnectIfExist()
		//		cliConnect()
		//	}},
	}
}

func CommonLivePrefix() (prefix string, useLivePrefix bool) {
	return CurrentDriver.LivePrefix(), true
}
