package operatorcli

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/Eoracle/core-go/internal/flag"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"path/filepath"
)

type manageKey struct {
	keyStorePath     string
	operatorKeystore string
	passphrase       string
	aliasECDSA       string
	aliasKeystore    string
	blsEncryptedFile string
}

type privateKeys struct {
	operatorECDSA    *ecdsa.PrivateKey
	aliasECDSA       *ecdsa.PrivateKey
	blsEncryptedFile *bls.KeyPair
}

func NewMangeKey(
	keyStorePath string,
	operatorECDSAFile string,
	passphrase string,
	aliasECDSA string,
	aliasECDSAFile string,
	blsEncryptedFile string,
) *manageKey {
	return &manageKey{
		keyStorePath:     keyStorePath,
		operatorKeystore: operatorECDSAFile,
		passphrase:       passphrase,
		aliasECDSA:       aliasECDSA,
		aliasKeystore:    aliasECDSAFile,
		blsEncryptedFile: blsEncryptedFile,
	}
}

func (m *manageKey) getPrivateKeys() (privateKeys, error) {
	operatorPrivateKey, err := m.getOperatorPrivateKey()
	if err != nil {
		return privateKeys{}, fmt.Errorf("failed read private Key for the operator, erorr: %s", err)
	}

	aliasPrivateKey, err := m.getAliasPrivateKey()
	if err != nil {
		return privateKeys{}, fmt.Errorf("failed read private key for the alias, error: %s", err)
	}

	blsPrivateKey, err := m.blsPrivateKey()
	if err != nil {
		return privateKeys{}, fmt.Errorf("failed read bls private error: %s", err)
	}

	return privateKeys{
		operatorECDSA:    operatorPrivateKey,
		aliasECDSA:       aliasPrivateKey,
		blsEncryptedFile: blsPrivateKey,
	}, nil
}

func (m *manageKey) getOperatorPrivateKey() (*ecdsa.PrivateKey, error) {
	fmt.Println("keyPath:", m.keyStorePath)

	operatorECDSAPair, err := eigensdkecdsa.ReadKey(filepath.Join(m.keyStorePath, m.operatorKeystore),
		m.passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s file %v", m.operatorKeystore, err)
	}
	return operatorECDSAPair, nil
}

func (m *manageKey) getAliasPrivateKey() (*ecdsa.PrivateKey, error) {
	aliasPrivateKey, err := eigensdkecdsa.ReadKey(filepath.Join(m.keyStorePath, m.aliasKeystore), m.passphrase)
	if err != nil {
		var ecdsaPair *ecdsa.PrivateKey
		// generate
		if m.aliasECDSA != "" {
			if len(m.aliasECDSA) >= 64 && (m.aliasECDSA[:2] == "0x" || m.aliasECDSA[:2] == "0X") {
				m.aliasECDSA = m.aliasECDSA[2:]
			}
			ecdsaPair, err = crypto.HexToECDSA(m.aliasECDSA)
		} else {
			ecdsaPair, err = crypto.GenerateKey()
		}

		if err != nil {
			return nil, err
		}

		err = eigensdkecdsa.WriteKey(filepath.Join(m.keyStorePath, m.aliasKeystore), ecdsaPair, m.passphrase)
		if err != nil {
			return nil, fmt.Errorf("error writing the %s file %v", m.aliasKeystore, err)
		}

		fmt.Println("alias ecdsa address", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "saved")
		return ecdsaPair, nil
	}
	fmt.Println("alias Private Key already exists")
	return aliasPrivateKey, nil
}

func (m *manageKey) blsPrivateKey() (*bls.KeyPair, error) {
	fmt.Println("bls private key path:", m.keyStorePath)

	blsKeyPair, err := eigensdkbls.ReadPrivateKeyFromFile(filepath.Join(m.keyStorePath, m.blsEncryptedFile), m.passphrase)
	if err != nil {
		return nil, fmt.Errorf("error reading the %s file %v", flag.EncryptedBLSFile, err)
	}
	fmt.Println("bls address G1, G2 ", blsKeyPair.GetPubKeyG1().String(), ", ", blsKeyPair.GetPubKeyG2().String(),
		"private key", blsKeyPair.PrivKey.String())
	return blsKeyPair, nil
}
