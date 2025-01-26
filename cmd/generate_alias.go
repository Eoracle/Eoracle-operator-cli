package cmd

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func NewGenerateAliasCommand() *cli.Command {
	return &cli.Command{
		Name:        "generate-alias",
		Description: "Create or Import an ECDSA private key only for oracle chain",
		Action:      runGenerateAlias,
		Flags: []cli.Flag{
			EcdsaPrivateKeyFlag,
			PassphraseFlag,
			KeyStorePathFlag,
			OverrideFlag,
		},
	}
}

func runGenerateAlias(c *cli.Context) error {
	if !c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name) {
		utils.Fatalf("passphrase and keystore-path are required")
	}

	aliasEcdsaPair, err := getAliasECDSAPrivateKey(c)
	if err != nil {
		return err
	}

	if err = eigensdkecdsa.WriteKey(
		filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaAliasedEncryptedWallet.json"),
		aliasEcdsaPair,
		c.String(PassphraseFlag.Name),
	); err != nil {
		utils.Fatalf("Error writing the ecdsaAliasedEncryptedWallet.json file %v", err)
	}

	fmt.Println("alias ecdsa address ", crypto.PubkeyToAddress(aliasEcdsaPair.PublicKey), "encrypted and saved")
	return nil
}

func getAliasECDSAPrivateKey(c *cli.Context) (*ecdsa.PrivateKey, error) {
	aliasEcdsaPair, err := eigensdkecdsa.ReadKey(filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaAliasedEncryptedWallet.json"), c.String(PassphraseFlag.Name))
	if err != nil {
		if c.String(EcdsaPrivateKeyFlag.Name) != "" {
			aliasEcdsaPair, err = crypto.HexToECDSA(c.String(EcdsaPrivateKeyFlag.Name))
			if err != nil {
				utils.Fatalf("Invalid ECDSA private key %v", err)
			}
		} else {
			aliasEcdsaPair, err = crypto.GenerateKey()
			if err != nil {
				utils.Fatalf("Failed to generate ECDSA key %v", err)
			}
			fmt.Println("a new alias ecdsa was generated address ", crypto.PubkeyToAddress(aliasEcdsaPair.PublicKey), "private key", hex.EncodeToString(aliasEcdsaPair.D.Bytes()))
		}
	} else {
		if c.String(EcdsaPrivateKeyFlag.Name) != "" {
			if !c.Bool(OverrideFlag.Name) {
				utils.Fatalf("alias already exists, use --override to replace it")
			}
			aliasEcdsaPair, err = crypto.HexToECDSA(c.String(EcdsaPrivateKeyFlag.Name))
			if err != nil {
				utils.Fatalf("Invalid ECDSA private key %v", err)
			}
		}
	}
	return aliasEcdsaPair, nil
}
