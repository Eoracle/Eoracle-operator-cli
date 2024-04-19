package operatorcli

import (
	"crypto/ecdsa"
	"fmt"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

type mangeKey struct {
	keyStorePath           string
	operatorKeystore       string
	operatorECDSPassphrase string
	aliasECDSA             string
	aliasKeystore          string
	aliasECDSAPassphrase   string
}

type priKeys struct {
	operatorECDSA *ecdsa.PrivateKey
	aliasECDSA    *ecdsa.PrivateKey
}

func NewMangeKey(
	keyStorePath string,
	operatorECDSAFile string,
	operatorECDSPassphrase string,
	aliasECDSA string,
	aliasECDSAFile string,
	aliasECDSAPassphrase string,
) *mangeKey {
	return &mangeKey{
		keyStorePath:           keyStorePath,
		operatorKeystore:       operatorECDSAFile,
		operatorECDSPassphrase: operatorECDSPassphrase,
		aliasECDSA:             aliasECDSA,
		aliasKeystore:          aliasECDSAFile,
		aliasECDSAPassphrase:   aliasECDSAPassphrase,
	}
}

func (m *mangeKey) getPrivateKeys() (priKeys, error) {
	operatorPK, err := m.getOperatorPrivateKey()
	if err != nil {
		return priKeys{}, err
	}

	aliasPK, err := m.getAliasPrivateKey()
	if err != nil {
		return priKeys{}, err
	}

	return priKeys{
		operatorECDSA: operatorPK,
		aliasECDSA:    aliasPK,
	}, nil
}

func (m *mangeKey) getOperatorPrivateKey() (*ecdsa.PrivateKey, error) {
	fmt.Println("keyPath:", m.keyStorePath)

	operatorECDSAPair, err := eigensdkecdsa.ReadKey(filepath.Join(m.keyStorePath, m.operatorKeystore),
		m.operatorECDSPassphrase)
	if err != nil {
		return nil, cli.Exit(fmt.Sprintf("Failed to read %s file %v", m.operatorKeystore, err), 1)
	}
	return operatorECDSAPair, nil
}

func (m *mangeKey) getAliasPrivateKey() (*ecdsa.PrivateKey, error) {
	aliasPrivateKey, err := eigensdkecdsa.ReadKey(filepath.Join(m.keyStorePath, m.aliasKeystore), m.aliasECDSAPassphrase)
	if err != nil {
		var ecdsaPair *ecdsa.PrivateKey
		// generate
		if m.aliasECDSA != "" {
			if m.aliasECDSA[:2] == "0x" || m.aliasECDSA[:2] == "0X" {
				m.aliasECDSA = m.aliasECDSA[2:]
			}
			ecdsaPair, err = crypto.HexToECDSA(m.aliasECDSA)
		} else {
			ecdsaPair, err = crypto.GenerateKey()
		}

		if err != nil {
			return nil, err
		}

		err = eigensdkecdsa.WriteKey(filepath.Join(m.keyStorePath, m.aliasKeystore), ecdsaPair, m.aliasECDSAPassphrase)
		if err != nil {
			return nil, fmt.Errorf("error writing the %s file %v", m.aliasKeystore, err)
		}

		fmt.Println("alias ecdsa address", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "saved")
		return ecdsaPair, nil
	}
	fmt.Println("alias Private Key already exists")
	return aliasPrivateKey, nil
}
