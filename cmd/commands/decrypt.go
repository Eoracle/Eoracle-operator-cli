package commands

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/eodata/operator-cli/cmd/flags"
	"github.com/eodata/operator-cli/internal/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v3"
)

func NewDecryptCommand(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:        "decrypt",
		Description: "Decrypt the ecdsa and bls private keys",
		Version:     version,
		HideVersion: true,
		Action:      runDecrypt,
		Flags: []cli.Flag{
			flags.PassphraseFlag,
			flags.EcdsaPassphraseFlag,
			flags.BlsPassphraseFlag,
			flags.KeyStorePathFlag,
		},
	}
}

func runDecrypt(ctx context.Context, c *cli.Command) error {
	logger := flags.Logger

	var ecdsaPassphrase string
	var blsPassphrase string

	if c.IsSet(flags.PassphraseFlag.Name) {
		ecdsaPassphrase = c.String(flags.PassphraseFlag.Name)
		blsPassphrase = c.String(flags.PassphraseFlag.Name)
	}

	if c.IsSet(flags.EcdsaPassphraseFlag.Name) {
		if ecdsaPassphrase != "" {
			return fmt.Errorf("either common passphrase or ecdsa passphrases should be set")
		}
		ecdsaPassphrase = c.String(flags.EcdsaPassphraseFlag.Name)
	}

	if c.IsSet(flags.BlsPassphraseFlag.Name) {
		if blsPassphrase != "" {
			return fmt.Errorf("either common passphrase or bls passphrases should be set")
		}
		blsPassphrase = c.String(flags.BlsPassphraseFlag.Name)
	}

	keystorePath := c.String(flags.KeyStorePathFlag.Name)

	senderEcdsaKeyPair, err := keystore.GetECDSAPrivateKey(ecdsaPassphrase, keystorePath, "")
	if err != nil {
		return fmt.Errorf("error reading the sender ECDSA private key %v", err)
	}

	logger.Info(
		"sender ecdsa info",
		"address", crypto.PubkeyToAddress(senderEcdsaKeyPair.PublicKey),
		"private key", hex.EncodeToString(senderEcdsaKeyPair.D.Bytes()),
	)

	blsKeyPair, err := keystore.GetBLSPrivateKey(blsPassphrase, keystorePath)
	if err != nil {
		return fmt.Errorf("error reading the BLS private key %v", err)
	}

	logger.Info(
		"BLS info",
		"G1", blsKeyPair.GetPubKeyG1().String(),
		"G2", blsKeyPair.GetPubKeyG2().String(),
		"private key", blsKeyPair.PrivKey.String(),
	)

	ecdsaAddresses, err := keystore.ListEcdsaAddresses(keystorePath)
	if err != nil {
		return fmt.Errorf("error listing the ECDSA keys %v", err)
	}

	for _, ecdsaAddress := range ecdsaAddresses {
		ecdsaKeyPair, err := keystore.GetECDSAPrivateKey(ecdsaPassphrase, keystorePath, ecdsaAddress)
		if err != nil {
			return fmt.Errorf("error reading the ECDSA private key of %s %v", ecdsaAddress, err)
		}

		logger.Info(
			"ECDSA info",
			"address", crypto.PubkeyToAddress(ecdsaKeyPair.PublicKey),
			"private key", hex.EncodeToString(ecdsaKeyPair.D.Bytes()),
		)
	}

	return nil
}
