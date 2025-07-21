package flags

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

var (
	profile   *NetworkProfile
	Logger, _ = logging.NewZapLogger(logging.Production)
)

type NetworkProfile struct {
	NetworkName string

	EOConfigAddress            gethcommon.Address
	RegistryCoordinatorAddress gethcommon.Address

	EOChainRPCEndpoint string
	EthRPCEndpoint     string
}
