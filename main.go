package main

import (
	_ "github.com/MrBruin-ctrl/runcli/env"
	"github.com/MrBruin-ctrl/runcli/internal/command"
	"github.com/c-bata/go-prompt"
)

func main() {
	command.ChooseCliDriver()
	p := prompt.New(
		command.GetActiveCmdExecutor,
		command.GetActiveCmdCompleter,
		prompt.OptionAddKeyBind(prompt.KeyBind{Key: prompt.F2, Fn: func(buffer *prompt.Buffer) {
			//重新执行
			command.ChooseCliDriver()
		}}),
		// register hot key for select active env
		// register live prefix that will be change automatically when env changed
		prompt.OptionLivePrefix(command.CommonLivePrefix),
		prompt.OptionTitle("Run-CLI: Interactive Cloud-Native Environment KubeClient"),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(8),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.LightGray),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionTextColor(prompt.White),
		prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
		prompt.OptionSelectedDescriptionBGColor(prompt.DarkBlue),
		prompt.OptionScrollbarBGColor(prompt.LightGray),
		prompt.OptionScrollbarThumbColor(prompt.Blue),
	)
	p.Run()

}
