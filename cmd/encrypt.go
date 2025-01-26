package cmd

import (
	"fmt"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func NewEncryptCommand() *cli.Command {
	return &cli.Command{
		Name:        "encrypt",
		Description: "Encrypt the ecdsa and bls private keys",
		Action:      runEncrypt,
		Flags: []cli.Flag{
			EcdsaPrivateKeyFlag,
			BlsPrivateKeyFlag,
			PassphraseFlag,
			KeyStorePathFlag,
		},
	}
}

func runEncrypt(c *cli.Context) error {
	// Encrypt the ecdsa private key and save it to a file
	ecdsaPair, err := getECDSAPrivateKey(c)
	if err != nil {
		utils.Fatalf("Invalid ECDSA private key %v", err)
	}

	if err = eigensdkecdsa.WriteKey(filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaEncryptedWallet.json"), ecdsaPair, c.String(PassphraseFlag.Name)); err != nil {
		utils.Fatalf("Error writing the ecdsaEncryptedWallet.json file %v", err)
	}

	fmt.Println("ecdsa address ", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "saved")

	// Encrypt the bls private key and save it to a file
	blsKeyPair, err := getBLSPrivateKey(c)
	if err != nil {
		utils.Fatalf("Invalid BLS private key %v", err)
	}

	if err = blsKeyPair.SaveToFile(filepath.Join(c.String(KeyStorePathFlag.Name), "blsEncryptedWallet.json"), c.String(PassphraseFlag.Name)); err != nil {
		utils.Fatalf("Error writing the blsEncryptedWallet.json file %v", err)
	}

	fmt.Println("bls address G1, G2 ", blsKeyPair.GetPubKeyG1().String(), ", ", blsKeyPair.GetPubKeyG2().String(), "saved")

	return nil
}
