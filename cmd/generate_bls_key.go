package cmd

import (
	"fmt"
	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/urfave/cli/v2"
)

func NewGenerateBLSKeyCommand() *cli.Command {
	return &cli.Command{
		Name:        "generate-bls-key",
		Description: "Generate the BLS key",
		Action:      runGenerateBLSKey,
		Flags:       []cli.Flag{},
	}
}

func runGenerateBLSKey(_ *cli.Context) error {
	keyPair, err := eigensdkbls.GenRandomBlsKeys()
	if err != nil {
		utils.Fatalf("Failed to generate BLS key pair %v", err)
	}
	fmt.Println("BLS private key", keyPair.PrivKey.String())
	fmt.Println("BLS public key G1", keyPair.GetPubKeyG1().String())
	fmt.Println("BLS public key G2", keyPair.GetPubKeyG2().String())
	return nil
}
