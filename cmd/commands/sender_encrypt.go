package commands

import (
	"context"
	"fmt"

	"github.com/eodata/operator-cli/cmd/flags"
	"github.com/eodata/operator-cli/internal/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v3"
)

func NewSenderEcdsaEncryptCommand(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:        "sender-ecdsa-encrypt",
		Description: "Encrypt the ecdsa private key used to send transactions",
		Version:     version,
		HideVersion: true,
		Action:      runSenderEcdsaEncrypt,
		Flags: []cli.Flag{
			flags.EcdsaPrivateKeyFlag,
			flags.PassphraseFlag,
			flags.EcdsaPassphraseFlag,
			flags.KeyStorePathFlag,
			flags.OverrideFlag,
		},
	}
}

func runSenderEcdsaEncrypt(ctx context.Context, c *cli.Command) error {
	logger := flags.Logger

	var passphrase string

	if c.IsSet(flags.PassphraseFlag.Name) {
		passphrase = c.String(flags.PassphraseFlag.Name)
	}

	if c.IsSet(flags.EcdsaPassphraseFlag.Name) {
		if passphrase != "" {
			return fmt.Errorf("either common passphrase or specific passphrases should be set")
		}
		passphrase = c.String(flags.EcdsaPassphraseFlag.Name)
	}

	if passphrase == "" {
		return fmt.Errorf("either common passphrase or specific passphrases should be set")
	}

	keystorePath := c.String(flags.KeyStorePathFlag.Name)

	ecdsaPair, err := crypto.HexToECDSA(c.String(flags.EcdsaPrivateKeyFlag.Name))
	if err != nil {
		return fmt.Errorf("invalid ECDSA private key %v", err)
	}
	address := crypto.PubkeyToAddress(ecdsaPair.PublicKey)

	_, err = keystore.GetECDSAPrivateKey(passphrase, keystorePath, "")
	if err == nil && !c.Bool(flags.OverrideFlag.Name) {
		logger.Info("sender ECDSA private key already exists, override flag not set", "address", address.Hex())
		return nil
	}

	if err = keystore.SaveECDSAPrivateKey(passphrase, keystorePath, "", ecdsaPair); err != nil {
		return fmt.Errorf("error writing the sender ecdsa private key to %s file %v", keystore.EcdsaEncryptedWallet, err)
	}
	logger.Info("sender ecdsa address ", "address", address.Hex())

	return nil
}
