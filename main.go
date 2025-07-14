package main

import (
	"fmt"
	"os"

	"github.com/eodata/operator-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "operatorCli"
	app.Description = "Eoracle generate operator signature and message for registration"
	app.Usage = "Used to create operator signature and message for registration"
	app.Commands = []*cli.Command{
		cmd.NewEncryptCommand(),
		cmd.NewDecryptCommand(),
		cmd.NewRegisterCommand(),
		cmd.NewDeregisterCommand(),
		cmd.NewPrintStatusCommand(),
		cmd.NewGenerateBLSKeyCommand(),
		cmd.NewGenerateAliasCommand(),
		cmd.NewDeclareAliasCommand(),
		cmd.NewResetConfigurationCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
