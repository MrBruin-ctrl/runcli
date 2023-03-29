package utils

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

func AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) {
	err := survey.AskOne(p, response, opts...)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func Ask(qs []*survey.Question, response interface{}, opts ...survey.AskOpt) {
	err := survey.Ask(qs, response, opts...)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
