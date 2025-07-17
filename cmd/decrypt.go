package cmd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func NewDecryptCommand() *cli.Command {
	return &cli.Command{
		Name:        "decrypt",
		Description: "Decrypt the ecdsa and bls private keys",
		Action:      runDecrypt,
		Flags: []cli.Flag{
			PassphraseFlag,
			KeyStorePathFlag,
		},
	}
}

func runDecrypt(c *cli.Context) error {
	ecdsaPair, err := getECDSAPrivateKey(c)
	if err != nil {
		utils.Fatalf("Error reading the ecdsaEncryptedWallet.json file %v", err)
	}
	fmt.Println("ecdsa address ", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "private key", hex.EncodeToString(ecdsaPair.D.Bytes()))

	ecdsaEOChainPair, err := eigensdkecdsa.ReadKey(filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaAliasedEncryptedWallet.json"), c.String(PassphraseFlag.Name))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("eochain alias was not set in the system")
			return nil
		}
		utils.Fatalf("Error reading the ecdsaAliasedEncryptedWallet.json file %v", err)
	}
	fmt.Println("EOChain ecdsa address ", crypto.PubkeyToAddress(ecdsaEOChainPair.PublicKey), "private key", hex.EncodeToString(ecdsaEOChainPair.D.Bytes()))

	blsKeyPair, err := getBLSPrivateKey(c)
	if err != nil {
		utils.Fatalf("Error reading the blsEncryptedWallet.json file %v", err)
	}

	fmt.Println("bls address G1, G2 ", blsKeyPair.GetPubKeyG1().String(), ", ", blsKeyPair.GetPubKeyG2().String(), "private key", blsKeyPair.PrivKey.String())

	return nil
}
