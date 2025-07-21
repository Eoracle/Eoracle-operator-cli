package commands

import (
	"context"
	"fmt"

	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/eodata/operator-cli/cmd/flags"
	"github.com/eodata/operator-cli/internal/keystore"
	"github.com/urfave/cli/v3"
)

func NewBlsEncryptCommand(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:        "bls-encrypt",
		Description: "Encrypt the bls private key used as operator alias",
		Version:     version,
		HideVersion: true,
		Action:      runBlsEncrypt,
		Flags: []cli.Flag{
			flags.BlsPrivateKeyFlag,
			flags.PassphraseFlag,
			flags.BlsPassphraseFlag,
			flags.KeyStorePathFlag,
			flags.OverrideFlag,
			flags.GenerateKeyPairFlag,
		},
	}
}

func runBlsEncrypt(ctx context.Context, c *cli.Command) error {
	logger := flags.Logger

	var passphrase string
	var blsKeyPair *eigensdkbls.KeyPair
	var err error

	if c.IsSet(flags.PassphraseFlag.Name) {
		passphrase = c.String(flags.PassphraseFlag.Name)
	}

	if c.IsSet(flags.BlsPassphraseFlag.Name) {
		if passphrase != "" {
			return fmt.Errorf("either common passphrase or specific passphrases should be set")
		}
		passphrase = c.String(flags.BlsPassphraseFlag.Name)
	}

	if passphrase == "" {
		return fmt.Errorf("either common passphrase or specific passphrases should be set")
	}

	keystorePath := c.String(flags.KeyStorePathFlag.Name)

	blsKeyPair, err = keystore.GetBLSPrivateKey(passphrase, keystorePath)
	if err == nil && !c.Bool(flags.OverrideFlag.Name) {
		logger.Info(
			"BLS private key already exists, override flag not set",
			"G1", blsKeyPair.GetPubKeyG1().String(),
			"G2", blsKeyPair.GetPubKeyG2().String(),
		)
		return nil
	}

	if c.Bool(flags.GenerateKeyPairFlag.Name) {
		blsKeyPair, err = keystore.GenerateBlsKeyPair()
		if err != nil {
			return fmt.Errorf("error generating BLS key %v", err)
		}
	}

	if err = keystore.SaveBLSPrivateKey(passphrase, keystorePath, blsKeyPair); err != nil {
		return fmt.Errorf("error writing the BLS private key to %s file %v", keystore.BlsEncryptedWallet, err)
	}
	logger.Info(
		"BLS private key saved",
		"G1", blsKeyPair.GetPubKeyG1().String(),
		"G2", blsKeyPair.GetPubKeyG2().String(),
	)

	return nil
}
