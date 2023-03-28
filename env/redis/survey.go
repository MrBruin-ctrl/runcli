package redis

import "github.com/AlecAivazis/survey/v2"

var BaseQs = []*survey.Question{
	{
		Name:     "Name",
		Prompt:   &survey.Input{Message: "What is your redis name?", Default: "redis-cli"},
		Validate: survey.Required,
	},
	{
		Name:     "Host",
		Prompt:   &survey.Input{Message: "What is your cli host?", Default: "localhost"},
		Validate: survey.Required,
	},
	{
		Name:     "Port",
		Prompt:   &survey.Input{Message: "What is your cli port?", Default: "6379"},
		Validate: survey.Required,
	},

	{
		Name:   "Password",
		Prompt: &survey.Password{Message: "What is your cli password?"},
	},
}

var SaveAndConnQs = []*survey.Question{
	{
		Name:     "IsSave",
		Prompt:   &survey.Confirm{Message: "save your conn conf?"},
		Validate: survey.Required,
	},
	{
		Name:     "IsConn",
		Prompt:   &survey.Confirm{Message: "is conn?"},
		Validate: survey.Required,
	},
}
