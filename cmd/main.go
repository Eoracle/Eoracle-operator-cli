package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/eodata/operator-cli/cmd/commands"
	"github.com/eodata/operator-cli/cmd/flags"
	"github.com/urfave/cli/v3"
)

const (
	// clientIdentifier to advertise over the network.
	clientIdentifier = "eo-operator-cli"
)

//go:embed VERSION
var version string

func main() {
	ctx := context.Background()
	logger := flags.Logger

	app := new(cli.Command)
	app.Name = clientIdentifier
	app.EnableShellCompletion = true
	app.Version = version
	app.Usage = "The EO AVS operator command line"
	app.Copyright = "Copyright 2025 The EO Authors"
	app.Commands = []*cli.Command{
		commands.NewSenderEcdsaEncryptCommand(ctx, app.Version),
		commands.NewAliasEcdsaEncryptCommand(ctx, app.Version),
		commands.NewBlsEncryptCommand(ctx, app.Version),
		commands.NewDecryptCommand(ctx, app.Version),
		commands.NewRegisterCommand(ctx, app.Version),
		// commands.NewDeregisterCommand(ctx,app.Version),
	}

	if err := app.Run(ctx, os.Args); err != nil {
		logger.Fatal("Error: ", err)
	}
}
