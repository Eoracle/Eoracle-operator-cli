package keystore

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	Suffix               = ".json"
	EcdsaPrefix          = "ecdsaEncryptedWallet"
	EcdsaEncryptedWallet = EcdsaPrefix + "%s" + Suffix
	BlsEncryptedWallet   = "blsEncryptedWallet" + Suffix
)

func GetECDSAPrivateKey(passphrase string, keystorePath string, address string) (*ecdsa.PrivateKey, error) {
	filePath := filepath.Join(keystorePath, fmt.Sprintf(EcdsaEncryptedWallet, address))
	return eigensdkecdsa.ReadKey(filePath, passphrase)
}

func SaveECDSAPrivateKey(passphrase string, keystorePath string, address string, key *ecdsa.PrivateKey) error {
	filePath := filepath.Join(keystorePath, fmt.Sprintf(EcdsaEncryptedWallet, address))
	return eigensdkecdsa.WriteKey(filePath, key, passphrase)
}

func GenerateEcdsaKeyPair() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func ListEcdsaAddresses(keystorePath string) ([]string, error) {
	files, err := os.ReadDir(keystorePath)
	if err != nil {
		return nil, err
	}

	ecdsaKeys := make([]string, 0)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), Suffix) {
			ecdsaKeys = append(
				ecdsaKeys,
				strings.TrimPrefix(strings.TrimSuffix(file.Name(), Suffix), EcdsaPrefix),
			)
		}
	}
	return ecdsaKeys, nil
}

func GetBLSPrivateKey(passphrase string, keystorePath string) (*eigensdkbls.KeyPair, error) {
	filePath := filepath.Join(keystorePath, BlsEncryptedWallet)
	return eigensdkbls.ReadPrivateKeyFromFile(filePath, passphrase)
}

func SaveBLSPrivateKey(passphrase string, keystorePath string, key *eigensdkbls.KeyPair) error {
	filePath := filepath.Join(keystorePath, BlsEncryptedWallet)
	return key.SaveToFile(filePath, passphrase)
}

func GenerateBlsKeyPair() (*eigensdkbls.KeyPair, error) {
	return eigensdkbls.GenRandomBlsKeys()
}
