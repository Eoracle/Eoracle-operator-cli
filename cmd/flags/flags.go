package flags

import (
	"github.com/urfave/cli/v3"
)

var (
	KeyStorePathFlag = &cli.StringFlag{
		Name:     "keystore-path",
		Usage:    "location of the keystore folder",
		Sources:  cli.EnvVars("EO_KEYSTORE_PATH"),
		Value:    ".keystore",
		Required: true,
	}
	EthRPCFlag = &cli.StringFlag{
		Name:     "eth-rpc-endpoint",
		Usage:    "ethereum rpc url",
		Sources:  cli.EnvVars("ETH_RPC_ENDPOINT"),
		Required: true,
	}
	OperatorAddressFlag = &cli.StringFlag{
		Name:     "operator-address",
		Usage:    "operator address",
		Sources:  cli.EnvVars("EO_OPERATOR_ADDRESS"),
		Required: true,
	}
	RegistryCoordinatorFlag = &cli.StringFlag{
		Name:     "registry-coordinator",
		Usage:    "registry coordinator contract address",
		Sources:  cli.EnvVars("EO_REGISTRY_COORDINATOR"),
		Required: true,
	}
	EcdsaPrivateKeyFlag = &cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "ecdsa private key",
		Sources:  cli.EnvVars("EO_ECDSA_PRIVATE_KEY"),
		Required: true,
	}
	BlsPrivateKeyFlag = &cli.StringFlag{
		Name:     "bls-private-key",
		Usage:    "bls private key",
		Sources:  cli.EnvVars("EO_BLS_PRIVATE_KEY"),
		Required: true,
	}
	SaltFlag = &cli.StringFlag{
		Name:    "salt",
		Usage:   "salt",
		Sources: cli.EnvVars("EO_SALT"),
		Value:   "0x01",
	}
	ExpiryFlag = &cli.StringFlag{
		Name:    "expiry",
		Usage:   "expiry",
		Sources: cli.EnvVars("EO_EXPIRY"),
		Value:   "115792089237316195423570985008687907853269984665640564039457584007913129639935",
	}
	PassphraseFlag = &cli.StringFlag{
		Name:    "passphrase",
		Usage:   "passphrase to open the encrypted private key",
		Sources: cli.EnvVars("EO_PASSPHRASE"),
	}
	EcdsaPassphraseFlag = &cli.StringFlag{
		Name:    "ecdsa-passphrase",
		Usage:   "ecdsa passphrase",
		Sources: cli.EnvVars("EO_ECDSA_PASSPHRASE"),
	}
	BlsPassphraseFlag = &cli.StringFlag{
		Name:    "bls-passphrase",
		Usage:   "bls passphrase",
		Sources: cli.EnvVars("EO_BLS_PASSPHRASE"),
	}
	ChainValidatorG1PointSignatureFlag = &cli.StringSliceFlag{
		Name:    "chain-validator-g1-point-signature",
		Usage:   "G1Point signature of the chain operator",
		Sources: cli.EnvVars("EO_CHAIN_VALIDATOR_G1_POINT_SIGNATURE"),
	}
	QuorumNumberFlag = &cli.IntFlag{
		Name:    "quorum-number",
		Usage:   "quorum number",
		Sources: cli.EnvVars("EO_QUORUM_NUMBER"),
		Value:   0,
	}
	EOChainRPCFlag = &cli.StringFlag{
		Name:    "eochain-rpc-endpoint",
		Usage:   "eochain rpc url",
		Sources: cli.EnvVars("EO_CHAIN_RPC_ENDPOINT"),
	}
	GenerateKeyPairFlag = &cli.BoolFlag{
		Name:    "generate-key-pair",
		Usage:   "Generate a new key pair",
		Sources: cli.EnvVars("EO_GENERATE_KEY_PAIR"),
	}
	OverrideFlag = &cli.BoolFlag{
		Name:    "alias-override",
		Usage:   "Indication if a new alias key should be created",
		Sources: cli.EnvVars("EO_ALIAS_OVERRIDE"),
	}
	EOConfigAddressFlag = &cli.StringFlag{
		Name:    "eoconfig-address",
		Usage:   "eoconfig contract address",
		Sources: cli.EnvVars("EO_CONFIG_ADDRESS"),
	}
	ProfileFlag = &cli.StringFlag{
		Name:    "profile",
		Usage:   "Network to use: mainnet or testnet.",
		Sources: cli.EnvVars("EO_PROFILE"),
		Value:   TestnetProfileName,
	}
	OperatorSetIDsFlag = &cli.StringSliceFlag{
		Name:     "operator-set-ids",
		Usage:    "operator set ids",
		Sources:  cli.EnvVars("EO_OPERATOR_SET_IDS"),
		Required: true,
	}
)
