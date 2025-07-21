package commands

import (
	"context"
	"fmt"

	"github.com/eodata/operator-cli/cmd/flags"
	"github.com/eodata/operator-cli/internal/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v3"
)

func NewAliasEcdsaEncryptCommand(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:        "alias-ecdsa-encrypt",
		Description: "Encrypt the ecdsa private key used as operator alias",
		Version:     version,
		HideVersion: true,
		Action:      runAliasEcdsaEncrypt,
		Flags: []cli.Flag{
			flags.EcdsaPrivateKeyFlag,
			flags.PassphraseFlag,
			flags.EcdsaPassphraseFlag,
			flags.KeyStorePathFlag,
			flags.GenerateKeyPairFlag,
		},
	}
}

func runAliasEcdsaEncrypt(ctx context.Context, c *cli.Command) error {
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

	_, err = keystore.GetECDSAPrivateKey(passphrase, keystorePath, address.Hex())
	if err == nil && !c.Bool(flags.OverrideFlag.Name) {
		logger.Info("alias ECDSA private key already exists", "address", address.Hex())
		return nil
	}

	if c.Bool(flags.GenerateKeyPairFlag.Name) {
		ecdsaPair, err = keystore.GenerateEcdsaKeyPair()
		if err != nil {
			return fmt.Errorf("error generating ECDSA key %v", err)
		}
	}
	address = crypto.PubkeyToAddress(ecdsaPair.PublicKey)

	if err = keystore.SaveECDSAPrivateKey(passphrase, keystorePath, address.Hex(), ecdsaPair); err != nil {
		return fmt.Errorf(
			"error writing the alias ecdsa private key to %s file %v",
			fmt.Sprintf(keystore.EcdsaEncryptedWallet, address.Hex()),
			err,
		)
	}
	logger.Info("alias ecdsa address ", "address", address.Hex())

	return nil
}
