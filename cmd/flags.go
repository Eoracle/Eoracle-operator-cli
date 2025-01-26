package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	KeyStorePathFlag = &cli.StringFlag{
		Name:    "keystore-path",
		Usage:   "location of the keystore folder",
		EnvVars: []string{"EO_KEYSTORE_PATH"},
		Value:   ".keystore",
	}
	EthRPCFlag = &cli.StringFlag{
		Name:    "eth-rpc-endpoint",
		Usage:   "ethereum rpc url",
		EnvVars: []string{"ETH_RPC_ENDPOINT"},
	}
	RegistryCoordinatorFlag = &cli.StringFlag{
		Name:    "registry-coordinator",
		Usage:   "registry coordinator contract address",
		EnvVars: []string{"EO_REGISTRY_COORDINATOR"},
	}
	EcdsaPrivateKeyFlag = &cli.StringFlag{
		Name:    "ecdsa-private-key",
		Usage:   "ecdsa private key",
		EnvVars: []string{"EO_ECDSA_PRIVATE_KEY"},
	}
	BlsPrivateKeyFlag = &cli.StringFlag{
		Name:    "bls-private-key",
		Usage:   "bls private key",
		EnvVars: []string{"EO_BLS_PRIVATE_KEY"},
	}
	SaltFlag = &cli.StringFlag{
		Name:    "salt",
		Usage:   "salt",
		EnvVars: []string{"EO_SALT"},
		Value:   "0x01",
	}
	ExpiryFlag = &cli.StringFlag{
		Name:    "expiry",
		Usage:   "expiry",
		EnvVars: []string{"EO_EXPIRY"},
		Value:   "115792089237316195423570985008687907853269984665640564039457584007913129639935",
	}
	PassphraseFlag = &cli.StringFlag{
		Name:    "passphrase",
		Usage:   "passphrase to open the encrypted private key",
		EnvVars: []string{"EO_PASSPHRASE"},
	}
	ValidatorRoleFlag = &cli.StringFlag{
		Name:    "validator-role",
		Usage:   "role of the operator",
		EnvVars: []string{"EO_VALIDATOR_ROLE"},
		Value:   "DATA_VALIDATOR",
	}
	ChainValidatorG1PointSignatureFlag = &cli.StringSliceFlag{
		Name:    "chain-validator-g1-point-signature",
		Usage:   "G1Point signature of the chain operator",
		EnvVars: []string{"EO_CHAIN_VALIDATOR_G1_POINT_SIGNATURE"},
	}
	QuorumNumberFlag = &cli.IntFlag{
		Name:    "quorum-number",
		Usage:   "quorum number",
		EnvVars: []string{"EO_QUORUM_NUMBER"},
		Value:   0,
	}
	EOChainRPCFlag = &cli.StringFlag{
		Name:    "eochain-rpc-endpoint",
		Usage:   "eochain rpc url",
		EnvVars: []string{"EO_CHAIN_RPC_ENDPOINT"},
	}
	OverrideFlag = &cli.BoolFlag{
		Name:    "alias-override",
		Usage:   "Indication if a new alias key should be created",
		EnvVars: []string{"EO_ALIAS_OVERRIDE"},
	}
	EOConfigAddressFlag = &cli.StringFlag{
		Name:    "eoconfig-address",
		Usage:   "eoconfig contract address",
		EnvVars: []string{"EO_CONFIG_ADDRESS"},
	}
	ProfileFlag = &cli.StringFlag{
		Name:    "profile",
		Usage:   "Network to use: mainnet or testnet.",
		EnvVars: []string{"EO_PROFILE"},
		Value:   TestnetProfileName,
	}
)
