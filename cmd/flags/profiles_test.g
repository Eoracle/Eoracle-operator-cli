package flags

import (
	flaglib "flag"
	"testing"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
)

func TestFromArgs(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedError   bool
		expectedProfile *NetworkProfile
	}{
		{
			name:            "valid mainnet profile",
			args:            []string{"--" + ProfileFlag.Name, MainnetProfileName},
			expectedError:   false,
			expectedProfile: &MainnetProfile,
		},
		{
			name:            "valid testnet profile",
			args:            []string{"--" + ProfileFlag.Name, TestnetProfileName},
			expectedError:   false,
			expectedProfile: &TestnetProfile,
		},
		{
			name:          "invalid profile",
			args:          []string{"--" + ProfileFlag.Name, "invalid"},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := flaglib.NewFlagSet("test", flaglib.ContinueOnError)
			set.String(ProfileFlag.Name, "", "profile name")
			c := cli.NewContext(nil, set, nil)
			err := set.Parse(tt.args)
			assert.NoError(t, err)
			err = SetProfile(c)
			assert.NoError(t, err)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedProfile, profile)
			}
		})
	}
}

func TestOverrideAddress(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		flagName        string
		expectedError   bool
		expectedAddress gethcommon.Address
	}{
		{
			name:            "valid address",
			args:            []string{"--" + EOConfigAddressFlag.Name, "0xf735Ad57952906a672eEEaDbef3bC69ECD24E50C"},
			flagName:        EOConfigAddressFlag.Name,
			expectedError:   false,
			expectedAddress: gethcommon.HexToAddress("0xf735Ad57952906a672eEEaDbef3bC69ECD24E50C"),
		},
		{
			name:          "invalid address",
			args:          []string{"--" + EOConfigAddressFlag.Name, "invalid"},
			flagName:      EOConfigAddressFlag.Name,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := flaglib.NewFlagSet("test", flaglib.ContinueOnError)
			set.String(tt.flagName, "", "address")
			set.String(ProfileFlag.Name, "testnet", "profile name")
			c := cli.NewContext(nil, set, nil)
			err := set.Parse(tt.args)
			assert.NoError(t, err)

			err = setProfile(c)
			assert.NoError(t, err)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedAddress, profile.EOConfigAddress)
			}
		})
	}
}

func TestOverrideURL(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		flagName      string
		expectedError bool
		expectedURL   string
	}{
		{
			name:          "valid URL",
			args:          []string{"--" + EOChainRPCFlag.Name, "https://rpc.eoracle.network"},
			flagName:      EOChainRPCFlag.Name,
			expectedError: false,
			expectedURL:   "https://rpc.eoracle.network",
		},
		{
			name:          "invalid URL",
			args:          []string{"--" + EOChainRPCFlag.Name, "invalid"},
			flagName:      EOChainRPCFlag.Name,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := flaglib.NewFlagSet("test", flaglib.ContinueOnError)
			set.String(tt.flagName, "", "URL")
			set.String(ProfileFlag.Name, "testnet", "profile name")
			c := cli.NewContext(nil, set, nil)
			err := set.Parse(tt.args)
			assert.NoError(t, err)
			err = setProfile(c)
			assert.NoError(t, err)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedURL, profile.EOChainRPCEndpoint)
			}
		})
	}
}
