package main

import (
	"fmt"
	"os"

	"github.com/Eoracle/core-go/internal/flag"
	"github.com/Eoracle/core-go/internal/operatorcli"
	"github.com/urfave/cli/v2"
)

var commandEncrypt = &cli.Command{
	Name:        "encrypt",
	Description: "Encrypt the ecdsa and bls private keys",
	Action:      runEncrypt,
	Flags: []cli.Flag{
		flag.EcdsaPrivateKeyFlag,
		flag.BlsPrivateKeyFlag,
		flag.PassphraseFlag,
		flag.KeyStorePathFlag,
	},
}
var commandDecrypt = &cli.Command{
	Name:        "decrypt",
	Description: "Decrypt the ecdsa and bls private keys",
	Action:      runDecrypt,
	Flags: []cli.Flag{
		flag.PassphraseFlag,
		flag.KeyStorePathFlag,
	},
}

var commandRegister = &cli.Command{
	Name:        "register",
	Description: "Register the operator",
	Action:      runRegister,
	Flags: []cli.Flag{
		flag.EthRPCFlag,
		flag.RegistryCoordinatorFlag,
		flag.PassphraseFlag,
		flag.KeyStorePathFlag,
		flag.SaltFlag,
		flag.ExpiryFlag,
		flag.EcdsaPrivateKeyFlag,
		flag.BlsPrivateKeyFlag,
		flag.ChainValidatorG1PointSignatureFlag,
		flag.ValidatorRoleFlag,
		flag.QuorumNumberFlag,
	},
}

var commandDeregister = &cli.Command{
	Name:        "deregister",
	Description: "Deregister the operator",
	Action:      runDeregister,
	Flags: []cli.Flag{
		flag.EthRPCFlag,
		flag.RegistryCoordinatorFlag,
		flag.PassphraseFlag,
		flag.KeyStorePathFlag,
		flag.EcdsaPrivateKeyFlag,
		flag.QuorumNumberFlag,
	},
}

var commandPrintStatus = &cli.Command{
	Name:        "print-status",
	Description: "Print the operator status",
	Action:      runPrintStatus,
	Flags: []cli.Flag{
		flag.EthRPCFlag,
		flag.RegistryCoordinatorFlag,
		flag.PassphraseFlag,
		flag.KeyStorePathFlag,
		flag.EcdsaPrivateKeyFlag,
		flag.QuorumNumberFlag,
	},
}

var commandGenerateBLSKey = &cli.Command{
	Name:        "generate-bls-key",
	Description: "Generate the BLS key",
	Action:      runGenerateBLSKey,
	Flags:       []cli.Flag{},
}

var commandEOChainSetAlias = &cli.Command{
	Name:        "eochain-set-alias",
	Description: "Create or Import a ECDSA private key only for oracle chain. The command can also update the eochain with the alias",
	Action:      runEOChainSetAlias,
	Flags:       []cli.Flag{
		flag.EcdsaPrivateKeyFlag,
		flag.PassphraseFlag,
		flag.KeyStorePathFlag,
		flag.EOChainEthRPCFlag,
		flag.EncryptOnlyFlag,
		flag.OverrideFlag,
		flag.EOConfigAddressFlag,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "operatorCli"
	app.Description = "Eoracle generate operator signature and message for registration"
	app.Usage = "Used to create operator signature and message for registration"
	app.Commands = []*cli.Command{
		commandEncrypt,
		commandDecrypt,
		commandRegister,
		commandDeregister,
		commandPrintStatus,
		commandGenerateBLSKey,
		commandEOChainSetAlias,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func runEncrypt(c *cli.Context) error {
	return operatorcli.RunEncrypt(c)
}

func runDecrypt(c *cli.Context) error {
	return operatorcli.RunDecrypt(c)
}

func runRegister(c *cli.Context) error {
	return operatorcli.RunRegister(c)
}

func runDeregister(c *cli.Context) error {
	return operatorcli.RunDeregister(c)
}

func runPrintStatus(c *cli.Context) error {
	return operatorcli.RunPrintStatus(c)
}

func runGenerateBLSKey(c *cli.Context) error {
	return operatorcli.RunGenerateBLSKey(c)
}

func runEOChainSetAlias(c *cli.Context) error {
	return operatorcli.RunEOChainSetAlias(c)
}
