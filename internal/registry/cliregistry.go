package registry

import (
	"runcli/internal/driver"
)

var cliFactory = map[string]driver.CliDriver{}

var cliNames = []string{}

type New func() driver.CliDriver

func RegisterCli(cli New) {
	cliDriver := cli()
	name := cliDriver.CliName()
	cliFactory[name] = cliDriver
	cliNames = append(cliNames, name)
}

func GetCliNames() []string {
	return cliNames
}

func GetCliByName(cliname string) driver.CliDriver {
	return cliFactory[cliname]
}
