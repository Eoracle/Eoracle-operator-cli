package flags

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v3"
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

		EOChainRPCEndpoint: "https://rpc.eo.app",
		EthRPCEndpoint:     "https://rpc.flashbots.net",
	}

	TestnetProfile = NetworkProfile{
		NetworkName: TestnetProfileName,

		EOConfigAddress:            gethcommon.HexToAddress("0xf735Ad57952906a672eEEaDbef3bC69ECD24E50C"),
		RegistryCoordinatorAddress: gethcommon.HexToAddress("0xc4A6E362e8Bd89F28Eb405F9Aa533784884B9c4F"),

		EOChainRPCEndpoint: "https://rpc.testnet.eo.app",
		EthRPCEndpoint:     "https://holesky.gateway.tenderly.com",
	}
)

func SetProfile(context context.Context, c *cli.Command) error {
	profileName := c.String(ProfileFlag.Name)

	switch profileName {
	case MainnetProfileName:
		profile = &MainnetProfile
	case TestnetProfileName:
		profile = &TestnetProfile
	default:
		return fmt.Errorf("invalid profile name: %s", profileName)
	}

	if err := OverrideAddress(c, EOConfigAddressFlag.Name, &profile.EOConfigAddress); err != nil {
		return err
	}

	if err := OverrideAddress(c, RegistryCoordinatorFlag.Name, &profile.RegistryCoordinatorAddress); err != nil {
		return err
	}

	if err := OverrideURL(c, EOChainRPCFlag.Name, &profile.EOChainRPCEndpoint); err != nil {
		return err
	}

	if err := OverrideURL(c, EthRPCFlag.Name, &profile.EthRPCEndpoint); err != nil {
		return err
	}

	return nil
}

func OverrideAddress(c *cli.Command, flagName string, address *gethcommon.Address) error {
	if c.IsSet(flagName) {
		addressOverride := c.String(flagName)
		if err := isValidAddress(addressOverride); err != nil {
			return fmt.Errorf("invalid %s: %v", flagName, err)
		}
		*address = gethcommon.HexToAddress(addressOverride)
	}
	return nil
}

func OverrideURL(c *cli.Command, flagName string, urlStr *string) error {
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
	address = strings.TrimPrefix(address, "0x")

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
