package profile

import (
	"encoding/hex"
	"fmt"
	"github.com/eoracle/eoracle-operator-cli/internal/flag"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"net/url"
	"strings"
)

const (
	MainnetProfileName = "mainnet"
	TestnetProfileName = "testnet"
)

var (
	MainnetProfile = NetworkProfile{
		NetworkName: MainnetProfileName,

		EOConfigAddress:            gethcommon.HexToAddress("0x05a6f762f64Ac2ccE0588677317a0Ed8af9d0c16"),
		RegistryCoordinatorAddress: gethcommon.HexToAddress("0x757E6f572AfD8E111bD913d35314B5472C051cA8"),

		EOChainRPCEndpoint: "https://rpc.eoracle.network",
		EthRPCEndpoint:     "https://rpc.flashbots.net",
	}

	TestnetProfile = NetworkProfile{
		NetworkName: TestnetProfileName,

		EOConfigAddress:            gethcommon.HexToAddress("0xf735Ad57952906a672eEEaDbef3bC69ECD24E50C"),
		RegistryCoordinatorAddress: gethcommon.HexToAddress("0xc4A6E362e8Bd89F28Eb405F9Aa533784884B9c4F"),

		EOChainRPCEndpoint: "https://rpc.testnet.eoracle.network",
		EthRPCEndpoint:     "https://ethereum-holesky-rpc.publicnode.com",
	}
)

type NetworkProfile struct {
	NetworkName string

	EOConfigAddress            gethcommon.Address
	RegistryCoordinatorAddress gethcommon.Address

	EOChainRPCEndpoint string
	EthRPCEndpoint     string
}

func FromArgs(c *cli.Context) (*NetworkProfile, error) {
	profileName := c.String(flag.ProfileFlag.Name)

	var networkProfile *NetworkProfile
	switch profileName {
	case MainnetProfileName:
		networkProfile = &MainnetProfile
	case TestnetProfileName:
		networkProfile = &TestnetProfile
	default:
		return nil, fmt.Errorf("invalid profile name: %s", profileName)
	}

	if err := overrideAddress(c, flag.EOConfigAddressFlag.Name, &networkProfile.EOConfigAddress); err != nil {
		return nil, err
	}

	if err := overrideAddress(c, flag.RegistryCoordinatorFlag.Name, &networkProfile.RegistryCoordinatorAddress); err != nil {
		return nil, err
	}

	if err := overrideURL(c, flag.EOChainRPCFlag.Name, &networkProfile.EOChainRPCEndpoint); err != nil {
		return nil, err
	}

	if err := overrideURL(c, flag.EthRPCFlag.Name, &networkProfile.EthRPCEndpoint); err != nil {
		return nil, err
	}

	return networkProfile, nil
}

func overrideAddress(c *cli.Context, flagName string, address *gethcommon.Address) error {
	if c.IsSet(flagName) {
		addressOverride := c.String(flagName)
		if err := isValidAddress(addressOverride); err != nil {
			return fmt.Errorf("invalid %s: %v", flagName, err)
		}
		*address = gethcommon.HexToAddress(addressOverride)
	}
	return nil
}

func overrideURL(c *cli.Context, flagName string, urlStr *string) error {
	if c.IsSet(flagName) {
		urlOverride := c.String(flagName)
		if err := isValidHttpURL(urlOverride); err != nil {
			return fmt.Errorf("invalid %s: %v", flagName, err)
		}
		*urlStr = urlOverride
	}
	return nil
}

// IsValidAddress checks if provided string is a valid Ethereum address
func isValidAddress(address string) error {
	// remove 0x prefix if it exists
	if strings.HasPrefix(address, "0x") {
		address = address[2:]
	}

	// decode the address
	decodedAddress, err := hex.DecodeString(address)
	if err != nil {
		return fmt.Errorf("address %s contains invalid characters", address)
	}

	// check if the address has the correct length
	if len(decodedAddress) != gethcommon.AddressLength {
		return fmt.Errorf("address %s has invalid length", string(decodedAddress))
	}

	return nil
}

func isValidHttpURL(u string) error {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("invalid scheme: %s", parsedURL.Scheme)
	}

	return nil
}
