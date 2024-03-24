// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractEORegistryCoordinator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IEOBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IEOBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	ChainValidatorSignature     BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IEORegistryCoordinatorOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEORegistryCoordinatorOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// IEORegistryCoordinatorOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IEORegistryCoordinatorOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// IEORegistryCoordinatorOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IEORegistryCoordinatorOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IEORegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IEORegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// IEOStakeRegistryStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IEOStakeRegistryStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractEORegistryCoordinatorMetaData contains all meta data concerning the ContractEORegistryCoordinator contract.
var ContractEORegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIEOStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIEOBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIEOIndexRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EOChainManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOChainManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEORegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIEORegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEOStakeRegistry.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEORegistryCoordinator.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEORegistryCoordinator.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEORegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEORegistryCoordinator.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEORegistryCoordinator.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEORegistryCoordinator.OperatorSetParam[]\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"_minimumStakes\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[][]\",\"internalType\":\"structIEOStakeRegistry.StrategyParams[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIEOBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"chainValidatorSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIEOBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"chainValidatorSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEORegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEOChainManager\",\"inputs\":[{\"name\":\"_EOChainManager\",\"type\":\"address\",\"internalType\":\"contractIEOChainManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIEORegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEORegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101c06040523480156200001257600080fd5b506040516200645238038062006452833981016040819052620000359162000254565b604080518082018252601c81527f656f7261636c65454f5265676973747279436f6f7264696e61746f720000000060208083019182528351808501909452600684526576302e302e3160d01b908401528151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a05287938793879387939192917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620001358184846040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6080523060c05261012052505050506001600160a01b039384166101405291831661018052821661016052166101a0526200016f62000179565b50505050620002bc565b600054610100900460ff1615620001e65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000239576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200025157600080fd5b50565b600080600080608085870312156200026b57600080fd5b845162000278816200023b565b60208601519094506200028b816200023b565b60408601519093506200029e816200023b565b6060860151909250620002b1816200023b565b939692955090935050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a05161608e620003c4600039600081816106550152818161122e01528181611fbc01528181612bb4015281816135660152613d8701526000818161059a01528181611f47015281816123ef01528181612b0b015281816134e60152818161394b0152613d0601526000818161056001528181610f4c01528181611f85015281816125c50152818161263b01528181612a8b015281816134680152613e03015260008181610491015281816129d301526133be015260006141ca01526000614219015260006141f40152600061414d01526000614177015260006141a1015261608e6000f3fe608060405234801561001057600080fd5b506004361061029f5760003560e01c80635df45946116101675780639feab859116100ce578063d75b4c8811610087578063d75b4c8814610713578063dd8283f314610726578063e65797ad14610739578063f2fde38b146107dc578063fabc1cbc146107ef578063fd39105a1461080257600080fd5b80639feab85914610677578063a65f39e11461069e578063c391425e146106b1578063ca0de882146106d1578063ca4f2d97146106f8578063d72d8dd61461070b57600080fd5b8063871ef04911610120578063871ef049146105ea578063886f1195146105fd5780638da5cb5b14610616578063993252b41461061e5780639aa1653d146106315780639e9923c21461065057600080fd5b80635df459461461055b5780636347c9001461058257806368304835146105955780636e3b17db146105bc578063715018a6146105cf57806384ca5213146105d757600080fd5b8063296bb0641161020b5780635865c60c116101c45780635865c60c146104e6578063595c6a67146105065780635ac86ab71461050e5780635b0b829f1461052d5780635b658e23146105405780635c975abb1461055357600080fd5b8063296bb0641461045357806329d1e0c3146104665780632cdd1e86146104795780633998fdd31461048c5780633c2a7f4c146104b35780635140a548146104d357600080fd5b8063136439dd1161025d578063136439dd1461037e5780631478851f146103915780631615700a146103c45780631eb812da146103d7578063249a0c421461042057806328f61b311461044057600080fd5b8062cf2ab5146102a457806303fd3492146102b957806304ec6351146102ec578063054310e61461031757806310d67a2f1461034257806313542a4e14610355575b600080fd5b6102b76102b2366004614b0a565b61083e565b005b6102d96102c7366004614b4b565b60009081526098602052604090205490565b6040519081526020015b60405180910390f35b6102ff6102fa366004614b76565b610954565b6040516001600160c01b0390911681526020016102e3565b609d5461032a906001600160a01b031681565b6040516001600160a01b0390911681526020016102e3565b6102b7610350366004614bd3565b610b4e565b6102d9610363366004614bd3565b6001600160a01b031660009081526099602052604090205490565b6102b761038c366004614b4b565b610c01565b6103b461039f366004614b4b565b609a6020526000908152604090205460ff1681565b60405190151581526020016102e3565b6102b76103d2366004614d8c565b610d45565b6103ea6103e5366004614e0b565b610ea2565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b0316908201526060016102e3565b6102d961042e366004614e3e565b609b6020526000908152604090205481565b609e5461032a906001600160a01b031681565b61032a610461366004614b4b565b610f33565b6102b7610474366004614bd3565b610fbf565b6102b7610487366004614bd3565b610fd0565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6104c66104c1366004614bd3565b610fe1565b6040516102e39190614e59565b6102b76104e1366004614e70565b611060565b6104f96104f4366004614bd3565b6115f3565b6040516102e39190614f13565b6102b7611667565b6103b461051c366004614e3e565b6001805460ff9092161b9081161490565b6102b761053b366004614f98565b611733565b60c85461032a906001600160a01b031681565b6001546102d9565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b61032a610590366004614b4b565b6117ca565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6102b76105ca366004614fcc565b6117f4565b6102b76118b4565b6102d96105e5366004615083565b6118c8565b6102ff6105f8366004614b4b565b611912565b60005461032a906201000090046001600160a01b031681565b61032a61191d565b6102b761062c366004615150565b611936565b60965461063e9060ff1681565b60405160ff90911681526020016102e3565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6102d97f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6102b76106ac366004614bd3565b611c41565b6106c46106bf366004615251565b611c6b565b6040516102e391906152f6565b6102d97f3bfe5390d1688d2e056e0ee4b68c5419f42d9bd0c792f01acb83886714adeb5681565b6102b7610706366004615340565b611d24565b609c546102d9565b6102b7610721366004615426565b611d8b565b6102b76107343660046155d9565b611d9e565b6107a8610747366004614e3e565b60408051606080820183526000808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff9081169183019190915292820151909216908201526060016102e3565b6102b76107ea366004614bd3565b6120aa565b6102b76107fd366004614b4b565b612120565b610831610810366004614bd3565b6001600160a01b031660009081526099602052604090206001015460ff1690565b6040516102e391906156ad565b600154600290600490811614156108705760405162461bcd60e51b8152600401610867906156bb565b60405180910390fd5b60005b8281101561094e57600084848381811061088f5761088f6156f2565b90506020020160208101906108a49190614bd3565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff1660028111156108ef576108ef614edb565b600281111561090057610900614edb565b905250805190915060006109138261227c565b90506000610929826001600160c01b03166122e5565b90506109368585836123b1565b505050505080806109469061571e565b915050610873565b50505050565b6000838152609860205260408120805482919084908110610977576109776156f2565b600091825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b03169183019190915290925085161015610a735760405162461bcd60e51b815260206004820152606760248201527f454f5265676973747279436f6f7264696e61746f722e67657451756f72756d4260448201527f69746d61704174426c6f636b4e756d6265724279496e6465783a2071756f727560648201527f6d4269746d61705570646174652069732066726f6d20616674657220626c6f6360848201526635a73ab6b132b960c91b60a482015260c401610867565b602081015163ffffffff161580610a995750806020015163ffffffff168463ffffffff16105b610b425760405162461bcd60e51b815260206004820152606860248201527f454f5265676973747279436f6f7264696e61746f722e67657451756f72756d4260448201527f69746d61704174426c6f636b4e756d6265724279496e6465783a2071756f727560648201527f6d4269746d61705570646174652069732066726f6d206265666f726520626c6f60848201526731b5a73ab6b132b960c11b60a482015260c401610867565b60400151949350505050565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ba1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc59190615739565b6001600160a01b0316336001600160a01b031614610bf55760405162461bcd60e51b815260040161086790615756565b610bfe8161249e565b50565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610c4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7291906157a0565b610c8e5760405162461bcd60e51b8152600401610867906157c2565b60015481811614610d075760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610867565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600180546000919081161415610d6d5760405162461bcd60e51b8152600401610867906156bb565b6000610d7933856125a3565b90506000610d9c3383898988610d97368c90038c0160408d0161580a565b6126db565b51905060005b86811015610e98576000888883818110610dbe57610dbe6156f2565b919091013560f81c600081815260976020526040902054855191935063ffffffff169150849084908110610df457610df46156f2565b602002602001015163ffffffff161115610e855760405162461bcd60e51b815260206004820152604660248201527f454f5265676973747279436f6f7264696e61746f722e72656769737465724f7060448201527f657261746f723a206f70657261746f7220636f756e742065786365656473206d6064820152656178696d756d60d01b608482015260a401610867565b5080610e908161571e565b915050610da2565b5050505050505050565b60408051606081018252600080825260208201819052918101919091526000838152609860205260409020805483908110610edf57610edf6156f2565b600091825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610f9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2d9190615739565b610fc7612d84565b610bfe81612de3565b610fd8612d84565b610bfe81612e4c565b6040805180820190915260008082526020820152610f2d61105b7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de6846040516020016110409291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120612eb5565b612f03565b600154600290600490811614156110895760405162461bcd60e51b8152600401610867906156bb565b60006110d184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff169150612f939050565b90506110dc81613045565b61114e5760405162461bcd60e51b8152602060048201526049602482015260008051602061603983398151915260448201527f61746f7273466f7251756f72756d3a20736f6d652071756f72756d7320646f206064820152681b9bdd08195e1a5cdd60ba1b608482015260a401610867565b8483146111bf5760405162461bcd60e51b8152602060048201526045602482015260008051602061603983398151915260448201527f61746f7273466f7251756f72756d3a20696e707574206c656e677468206d69736064820152640dac2e8c6d60db1b608482015260a401610867565b60005b838110156115ea5760008585838181106111de576111de6156f2565b919091013560f81c915036905060008989858181106111ff576111ff6156f2565b90506020028101906112119190615858565b6040516379a0849160e11b815260ff8616600482015291935091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f341092290602401602060405180830381865afa15801561127d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a191906158a1565b63ffffffff16811461133f5760405162461bcd60e51b8152602060048201526067602482015260008051602061603983398151915260448201527f61746f7273466f7251756f72756d3a206e756d626572206f662075706461746560648201527f64206f70657261746f727320646f6573206e6f74206d617463682071756f72756084820152661b481d1bdd185b60ca1b60a482015260c401610867565b6000805b8281101561158957600084848381811061135f5761135f6156f2565b90506020020160208101906113749190614bd3565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff1660028111156113bf576113bf614edb565b60028111156113d0576113d0614edb565b905250805190915060006113e38261227c565b905060016001600160c01b03821660ff8b161c8116146114685760405162461bcd60e51b8152602060048201526046602482015260008051602061603983398151915260448201527f61746f7273466f7251756f72756d3a206f70657261746f72206e6f7420696e2060648201526571756f72756d60d01b608482015260a401610867565b856001600160a01b0316846001600160a01b0316116115155760405162461bcd60e51b8152602060048201526069602482015260008051602061603983398151915260448201527f61746f7273466f7251756f72756d3a206f70657261746f72732061727261792060648201527f6d75737420626520736f7274656420696e20617363656e64696e67206164647260848201526832b9b99037b93232b960b91b60a482015260c401610867565b5061157383838f8f8d908e600161152c91906158be565b92611539939291906158d6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506123b192505050565b5090925061158290508161571e565b9050611343565b5060ff84166000818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a250505050806115e39061571e565b90506111c2565b50505050505050565b60408051808201909152600080825260208201526001600160a01b0382166000908152609960209081526040918290208251808401909352805483526001810154909183019060ff16600281111561164d5761164d614edb565b600281111561165e5761165e614edb565b90525092915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156116b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d891906157a0565b6116f45760405162461bcd60e51b8152600401610867906157c2565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b61173b612d84565b609654829060ff908116908216106117bb5760405162461bcd60e51b815260206004820152603960248201527f454f5265676973747279436f6f7264696e61746f722e71756f72756d4578697360448201527f74733a2071756f72756d20646f6573206e6f74206578697374000000000000006064820152608401610867565b6117c58383613078565b505050565b609c81815481106117da57600080fd5b6000918252602090912001546001600160a01b0316905081565b609e546001600160a01b031633146118745760405162461bcd60e51b815260206004820152603c60248201527f454f5265676973747279436f6f7264696e61746f722e6f6e6c79456a6563746f60448201527f723a2063616c6c6572206973206e6f742074686520656a6563746f72000000006064820152608401610867565b6117c58383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061312592505050565b6118bc612d84565b6118c6600061363f565b565b60006119087f3bfe5390d1688d2e056e0ee4b68c5419f42d9bd0c792f01acb83886714adeb56878787878760405160200161104096959493929190615900565b9695505050505050565b6000610f2d8261227c565b60006119316064546001600160a01b031690565b905090565b60018054600091908116141561195e5760405162461bcd60e51b8152600401610867906156bb565b8387146119e25760405162461bcd60e51b815260206004820152604660248201527f454f5265676973747279436f6f7264696e61746f722e72656769737465724f7060448201527f657261746f7257697468436875726e3a20696e707574206c656e677468206d696064820152650e6dac2e8c6d60d31b608482015260a401610867565b60006119ee33886125a3565b9050611a4e33828888808060200260200160405190810160405280939291908181526020016000905b82821015611a4357611a3460408302860136819003810190615985565b81526020019060010190611a17565b505050505087613691565b6000611a6c33838c8c888d604001803603810190610d97919061580a565b905060005b89811015611c34576000609760008d8d85818110611a9157611a916156f2565b919091013560f81c82525060208082019290925260409081016000208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b90910490931691810191909152845180519193509084908110611afe57611afe6156f2565b602002602001015163ffffffff161115611c2157611b9f8c8c84818110611b2757611b276156f2565b9050013560f81c60f81b60f81c84604001518481518110611b4a57611b4a6156f2565b60200260200101513386602001518681518110611b6957611b696156f2565b60200260200101518d8d88818110611b8357611b836156f2565b905060400201803603810190611b999190615985565b86613822565b611c21898984818110611bb457611bb46156f2565b9050604002016020016020810190611bcc9190614bd3565b8d848e611bda8260016158be565b92611be7939291906158d6565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061312592505050565b5080611c2c8161571e565b915050611a71565b5050505050505050505050565b611c49612d84565b60c880546001600160a01b0319166001600160a01b0392909216919091179055565b6060600082516001600160401b03811115611c8857611c88614c4a565b604051908082528060200260200182016040528015611cb1578160200160208202803683370190505b50905060005b8351811015611d1c57611ce385858381518110611cd657611cd66156f2565b6020026020010151613b08565b828281518110611cf557611cf56156f2565b63ffffffff9092166020928302919091019091015280611d148161571e565b915050611cb7565b509392505050565b6001805460029081161415611d4b5760405162461bcd60e51b8152600401610867906156bb565b6117c53384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061312592505050565b611d93612d84565b6117c5838383613c46565b600054610100900460ff1615808015611dbe5750600054600160ff909116105b80611dd85750303b158015611dd8575060005460ff166001145b611e3b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610867565b6000805460ff191660011790558015611e5e576000805461ff0019166101001790555b82518451148015611e70575081518351145b611ee25760405162461bcd60e51b815260206004820152603760248201527f454f5265676973747279436f6f7264696e61746f722e696e697469616c697a6560448201527f3a20696e707574206c656e677468206d69736d617463680000000000000000006064820152608401610867565b611eeb8961363f565b611ef58686613e65565b611efe88612de3565b611f0787612e4c565b609c80546001818101835560008381527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c92830180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925585548085018755850180547f0000000000000000000000000000000000000000000000000000000000000000841690831617905585549384019095559190920180547f000000000000000000000000000000000000000000000000000000000000000090921691909316179091555b845181101561205857612046858281518110612005576120056156f2565b602002602001015185838151811061201f5761201f6156f2565b6020026020010151858481518110612039576120396156f2565b6020026020010151613c46565b806120508161571e565b915050611fe7565b50801561209f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b6120b2612d84565b6001600160a01b0381166121175760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610867565b610bfe8161363f565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612173573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121979190615739565b6001600160a01b0316336001600160a01b0316146121c75760405162461bcd60e51b815260040161086790615756565b6001541981196001541916146122455760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610867565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610d3a565b600081815260986020526040812054806122995750600092915050565b60008381526098602052604090206122b26001836159a1565b815481106122c2576122c26156f2565b600091825260209091200154600160401b90046001600160c01b03169392505050565b60606000806122f384613f55565b61ffff166001600160401b0381111561230e5761230e614c4a565b6040519080825280601f01601f191660200182016040528015612338576020820181803683370190505b5090506000805b825182108015612350575061010081105b156123a7576001811b935085841615612397578060f81b838381518110612379576123796156f2565b60200101906001600160f81b031916908160001a9053508160010191505b6123a08161571e565b905061233f565b5090949350505050565b6001826020015160028111156123c9576123c9614edb565b146123d357505050565b81516040516333567f7f60e11b81526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe9061242890889086908890600401615a05565b6020604051808303816000875af1158015612447573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061246b9190615a35565b90506001600160c01b038116156124975761249785612492836001600160c01b03166122e5565b613125565b5050505050565b6001600160a01b03811661252c5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610867565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b6040516309aa152760e11b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa15801561260e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126329190615a5e565b905080610f2d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637454b9f1848461267387610fe1565b6040518463ffffffff1660e01b815260040161269193929190615a77565b6020604051808303816000875af11580156126b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126d49190615a5e565b9392505050565b6126ff60405180606001604052806060815260200160608152602001606081525090565b600061274786868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff169150612f939050565b905060006127548861227c565b90506001600160c01b0382166127d25760405162461bcd60e51b815260206004820152603b60248201527f454f5265676973747279436f6f7264696e61746f722e5f72656769737465724f60448201527f70657261746f723a206269746d61702063616e6e6f74206265203000000000006064820152608401610867565b6127db82613045565b6128585760405162461bcd60e51b815260206004820152604260248201527f454f5265676973747279436f6f7264696e61746f722e5f72656769737465724f60448201527f70657261746f723a20736f6d652071756f72756d7320646f206e6f74206578696064820152611cdd60f21b608482015260a401610867565b8082166001600160c01b0316156129105760405162461bcd60e51b815260206004820152606a60248201527f454f5265676973747279436f6f7264696e61746f722e5f72656769737465724f60448201527f70657261746f723a206f70657261746f7220616c72656164792072656769737460648201527f6572656420666f7220736f6d652071756f72756d73206265696e67207265676960848201526939ba32b932b2103337b960b11b60a482015260c401610867565b6001600160c01b03818116908316176129298982613f80565b60016001600160a01b038b1660009081526099602052604090206001015460ff16600281111561295b5761295b614edb565b14612a74576040805180820182528a8152600160208083018281526001600160a01b038f166000908152609990925293902082518155925183820180549394939192909160ff1916908360028111156129b6576129b6614edb565b021790555050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169150639926ee7d90612a0b908d908a90600401615b12565b600060405180830381600087803b158015612a2557600080fd5b505af1158015612a39573d6000803e3d6000fd5b50506040518b92506001600160a01b038d1691507fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe90600090a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290612ac4908d908c908c90600401615b86565b600060405180830381600087803b158015612ade57600080fd5b505af1158015612af2573d6000803e3d6000fd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150612b48908d908d908d908d90600401615bab565b6000604051808303816000875af1158015612b67573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612b8f9190810190615c37565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90612bec908c908c908c90600401615c9a565b6000604051808303816000875af1158015612c0b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612c339190810190615cb4565b845260c8546001600160a01b031615612d7757845160208601518115801590612c5b57508015155b15612cfa576000604051806040016040528084815260200183815250905060c860009054906101000a90046001600160a01b03166001600160a01b031663ff2acee08e8960200151846040518463ffffffff1660e01b8152600401612cc293929190615d91565b600060405180830381600087803b158015612cdc57600080fd5b505af1158015612cf0573d6000803e3d6000fd5b5050505050612d74565b60c860009054906101000a90046001600160a01b03166001600160a01b0316639e5a13748d88602001516040518363ffffffff1660e01b8152600401612d41929190615de9565b600060405180830381600087803b158015612d5b57600080fd5b505af1158015612d6f573d6000803e3d6000fd5b505050505b50505b5050509695505050505050565b33612d8d61191d565b6001600160a01b0316146118c65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610867565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610f2d612ec2614140565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b604080518082019091526000808252602082015260008080612f33600080516020615ff983398151915286615e2b565b90505b612f3f81614267565b9093509150600080516020615ff9833981519152828309831415612f79576040805180820190915290815260208101919091529392505050565b600080516020615ff9833981519152600182089050612f36565b600080612f9f846142e9565b905080156126d4578260ff168460018651612fba91906159a1565b81518110612fca57612fca6156f2565b016020015160f81c106126d45760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610867565b609654600090819061305f9060019060ff1681901b6159a1565b90506126d46001600160c01b0384811690831681161490565b60ff8216600081815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac9060600160405180910390a25050565b6001600160a01b0382166000908152609960205260409020805460018083015460ff16600281111561315957613159614edb565b146131c85760405162461bcd60e51b81526020600482015260456024820152600080516020615fd983398151915260448201527f724f70657261746f723a206f70657261746f72206973206e6f742072656769736064820152641d195c995960da1b608482015260a401610867565b6096546000906131dc90859060ff16612f93565b905060006131e98361227c565b90506001600160c01b0382166132555760405162461bcd60e51b815260206004820152603d6024820152600080516020615fd983398151915260448201527f724f70657261746f723a206269746d61702063616e6e6f7420626520300000006064820152608401610867565b61325e82613045565b6132cc5760405162461bcd60e51b815260206004820152604460248201819052600080516020615fd9833981519152908201527f724f70657261746f723a20736f6d652071756f72756d7320646f206e6f7420656064820152631e1a5cdd60e21b608482015260a401610867565b6132e36001600160c01b0383811690831681161490565b6133695760405162461bcd60e51b815260206004820152605b6024820152600080516020615fd983398151915260448201527f724f70657261746f723a206f70657261746f72206973206e6f7420726567697360648201527f746572656420666f72207370656369666965642071756f72756d730000000000608482015260a401610867565b6001600160c01b03828116198216166133828482613f80565b6001600160c01b0381166134515760018501805460ff191660021790556040516351b27a6d60e11b81526001600160a01b0388811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b15801561340257600080fd5b505af1158015613416573d6000803e3d6000fd5b50506040518692506001600160a01b038a1691507f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe59061349f908a908a90600401615e3f565b600060405180830381600087803b1580156134b957600080fd5b505af11580156134cd573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061351f9087908a90600401615e63565b600060405180830381600087803b15801561353957600080fd5b505af115801561354d573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061359f9087908a90600401615e63565b600060405180830381600087803b1580156135b957600080fd5b505af11580156135cd573d6000803e3d6000fd5b505060c8546001600160a01b03161591506115ea90505760c8546040516393835a5b60e01b81526001600160a01b038981166004830152909116906393835a5b90602401600060405180830381600087803b15801561362b57600080fd5b505af1158015611c34573d6000803e3d6000fd5b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6020808201516000908152609a909152604090205460ff16156137395760405162461bcd60e51b815260206004820152605460248201527f454f5265676973747279436f6f7264696e61746f722e5f76657269667943687560448201527f726e417070726f7665725369676e61747572653a20636875726e417070726f76606482015273195c881cd85b1d08185b1c9958591e481d5cd95960621b608482015260a401610867565b42816040015110156137d05760405162461bcd60e51b815260206004820152605460248201527f454f5265676973747279436f6f7264696e61746f722e5f76657269667943687560448201527f726e417070726f7665725369676e61747572653a20636875726e417070726f76606482015273195c881cda59db985d1d5c9948195e1c1a5c995960621b608482015260a401610867565b602080820180516000908152609a909252604091829020805460ff19166001179055609d5490519183015161094e926001600160a01b039092169161381b91889188918891906118c8565b8351614476565b6020808301516001600160a01b0380821660008181526099909452604090932054919290871614156138aa5760405162461bcd60e51b8152602060048201526037602482015260008051602061601983398151915260448201527f6875726e3a2063616e6e6f7420636875726e2073656c660000000000000000006064820152608401610867565b8760ff16846000015160ff16146139295760405162461bcd60e51b8152602060048201526049602482015260008051602061601983398151915260448201527f6875726e3a2071756f72756d4e756d626572206e6f74207468652073616d6520606482015268185cc81cda59db995960ba1b608482015260a401610867565b604051635401ed2760e01b81526004810182905260ff891660248201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa15801561399a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139be9190615e7c565b90506139ca8185614630565b6001600160601b0316866001600160601b031611613a645760405162461bcd60e51b8152602060048201526058602482015260008051602061601983398151915260448201527f6875726e3a20696e636f6d696e67206f70657261746f722068617320696e737560648201527f6666696369656e74207374616b6520666f7220636875726e0000000000000000608482015260a401610867565b613a6e8885614654565b6001600160601b0316816001600160601b03161061209f5760405162461bcd60e51b815260206004820152605e602482015260008051602061601983398151915260448201527f6875726e3a2063616e6e6f74206b69636b206f70657261746f7220776974682060648201527f6d6f7265207468616e206b69636b424950734f66546f74616c5374616b650000608482015260a401610867565b600081815260986020526040812054815b81811015613b9a576001613b2d82846159a1565b613b3791906159a1565b92508463ffffffff16609860008681526020019081526020016000208463ffffffff1681548110613b6a57613b6a6156f2565b60009182526020909120015463ffffffff1611613b88575050610f2d565b80613b928161571e565b915050613b19565b5060405162461bcd60e51b815260206004820152606e60248201527f454f5265676973747279436f6f7264696e61746f722e67657451756f72756d4260448201527f69746d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d60648201527f61702075706461746520666f756e6420666f72206f70657261746f724964206160848201526d3a10313637b1b590373ab6b132b960911b60a482015260c401610867565b60965460ff1660c08110613cc25760405162461bcd60e51b815260206004820152603760248201527f454f5265676973747279436f6f7264696e61746f722e63726561746551756f7260448201527f756d3a206d61782071756f72756d7320726561636865640000000000000000006064820152608401610867565b613ccd816001615e99565b6096805460ff191660ff9290921691909117905580613cec8186613078565b60405160016296b58960e01b031981526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ff694a7790613d3f90849088908890600401615ebe565b600060405180830381600087803b158015613d5957600080fd5b505af1158015613d6d573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b158015613dd557600080fd5b505af1158015613de9573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b158015613e5157600080fd5b505af115801561209f573d6000803e3d6000fd5b6000546201000090046001600160a01b0316158015613e8c57506001600160a01b03821615155b613f0e5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610867565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613f518261249e565b5050565b6000805b8215610f2d57613f6a6001846159a1565b9092169180613f7881615f37565b915050613f59565b60008281526098602052604090205480614025576000838152609860209081526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055505050565b600083815260986020526040812061403e6001846159a1565b8154811061404e5761404e6156f2565b600091825260209091200180549091504363ffffffff908116911614156140925780546001600160401b0316600160401b6001600160c01b0385160217815561094e565b805463ffffffff438116600160201b81810267ffffffff0000000019909416939093178455600087815260986020908152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff199093169390941692909217179190911691909117905550505050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561419957507f000000000000000000000000000000000000000000000000000000000000000046145b156141c357507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60008080600080516020615ff98339815191526003600080516020615ff983398151915286600080516020615ff98339815191528889090908905060006142dd827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615ff983398151915261466e565b91959194509092505050565b6000610100825111156143725760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610867565b815161438057506000919050565b60008083600081518110614396576143966156f2565b0160200151600160f89190911c81901b92505b845181101561446d578481815181106143c4576143c46156f2565b0160200151600160f89190911c1b91508282116144595760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610867565b918117916144668161571e565b90506143a9565b50909392505050565b6001600160a01b0383163b1561459057604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906144b69086908690600401615e63565b602060405180830381865afa1580156144d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144f79190615f59565b6001600160e01b031916146117c55760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610867565b826001600160a01b03166145a4838361471d565b6001600160a01b0316146117c55760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610867565b60208101516000906127109061464a9061ffff1685615f83565b6126d49190615fb2565b60408101516000906127109061464a9061ffff1685615f83565b600080614679614a8a565b614681614aa8565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156146c2576146c4565bfe5b50826147125760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610867565b505195945050505050565b600080600061472c8585614739565b91509150611d1c816147a9565b6000808251604114156147705760208301516040840151606085015160001a61476487828585614964565b945094505050506147a2565b82516040141561479a576020830151604084015161478f868383614a51565b9350935050506147a2565b506000905060025b9250929050565b60008160048111156147bd576147bd614edb565b14156147c65750565b60018160048111156147da576147da614edb565b14156148285760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610867565b600281600481111561483c5761483c614edb565b141561488a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610867565b600381600481111561489e5761489e614edb565b14156148f75760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610867565b600481600481111561490b5761490b614edb565b1415610bfe5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610867565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561499b5750600090506003614a48565b8460ff16601b141580156149b357508460ff16601c14155b156149c45750600090506004614a48565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614a18573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614a4157600060019250925050614a48565b9150600090505b94509492505050565b6000806001600160ff1b03831681614a6e60ff86901c601b6158be565b9050614a7c87828885614964565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60008083601f840112614ad857600080fd5b5081356001600160401b03811115614aef57600080fd5b6020830191508360208260051b85010111156147a257600080fd5b60008060208385031215614b1d57600080fd5b82356001600160401b03811115614b3357600080fd5b614b3f85828601614ac6565b90969095509350505050565b600060208284031215614b5d57600080fd5b5035919050565b63ffffffff81168114610bfe57600080fd5b600080600060608486031215614b8b57600080fd5b833592506020840135614b9d81614b64565b929592945050506040919091013590565b6001600160a01b0381168114610bfe57600080fd5b8035614bce81614bae565b919050565b600060208284031215614be557600080fd5b81356126d481614bae565b60008083601f840112614c0257600080fd5b5081356001600160401b03811115614c1957600080fd5b6020830191508360208285010111156147a257600080fd5b60006101408284031215614c4457600080fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614c8257614c82614c4a565b60405290565b604080519081016001600160401b0381118282101715614c8257614c82614c4a565b604051601f8201601f191681016001600160401b0381118282101715614cd257614cd2614c4a565b604052919050565b600060608284031215614cec57600080fd5b614cf4614c60565b905081356001600160401b0380821115614d0d57600080fd5b818401915084601f830112614d2157600080fd5b8135602082821115614d3557614d35614c4a565b614d47601f8301601f19168201614caa565b92508183528681838601011115614d5d57600080fd5b818185018285013760008183850101528285528086013581860152505050506040820135604082015292915050565b6000806000806101808587031215614da357600080fd5b84356001600160401b0380821115614dba57600080fd5b614dc688838901614bf0565b9096509450849150614ddb8860208901614c31565b9350610160870135915080821115614df257600080fd5b50614dff87828801614cda565b91505092959194509250565b60008060408385031215614e1e57600080fd5b50508035926020909101359150565b803560ff81168114614bce57600080fd5b600060208284031215614e5057600080fd5b6126d482614e2d565b815181526020808301519082015260408101610f2d565b60008060008060408587031215614e8657600080fd5b84356001600160401b0380821115614e9d57600080fd5b614ea988838901614ac6565b90965094506020870135915080821115614ec257600080fd5b50614ecf87828801614bf0565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b60038110614f0f57634e487b7160e01b600052602160045260246000fd5b9052565b815181526020808301516040830191614f2e90840182614ef1565b5092915050565b803561ffff81168114614bce57600080fd5b600060608284031215614f5957600080fd5b614f61614c60565b90508135614f6e81614b64565b8152614f7c60208301614f35565b6020820152614f8d60408301614f35565b604082015292915050565b60008060808385031215614fab57600080fd5b614fb483614e2d565b9150614fc38460208501614f47565b90509250929050565b600080600060408486031215614fe157600080fd5b8335614fec81614bae565b925060208401356001600160401b0381111561500757600080fd5b61501386828701614bf0565b9497909650939450505050565b60006001600160401b0382111561503957615039614c4a565b5060051b60200190565b60006040828403121561505557600080fd5b61505d614c88565b905061506882614e2d565b8152602082013561507881614bae565b602082015292915050565b600080600080600060a0868803121561509b57600080fd5b85356150a681614bae565b945060208681013594506040808801356001600160401b038111156150ca57600080fd5b8801601f81018a136150db57600080fd5b80356150ee6150e982615020565b614caa565b81815260069190911b8201840190848101908c83111561510d57600080fd5b928501925b82841015615133576151248d85615043565b82529284019290850190615112565b999c989b5098996060810135995060800135979650505050505050565b60008060008060008060006101c0888a03121561516c57600080fd5b87356001600160401b038082111561518357600080fd5b61518f8b838c01614bf0565b90995097508791506151a48b60208c01614c31565b96506101608a01359150808211156151bb57600080fd5b818a0191508a601f8301126151cf57600080fd5b8135818111156151de57600080fd5b8b60208260061b85010111156151f357600080fd5b602083019650809550506101808a013591508082111561521257600080fd5b61521e8b838c01614cda565b93506101a08a013591508082111561523557600080fd5b506152428a828b01614cda565b91505092959891949750929550565b6000806040838503121561526457600080fd5b823561526f81614b64565b91506020838101356001600160401b0381111561528b57600080fd5b8401601f8101861361529c57600080fd5b80356152aa6150e982615020565b81815260059190911b820183019083810190888311156152c957600080fd5b928401925b828410156152e7578335825292840192908401906152ce565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b8181101561533457835163ffffffff1683529284019291840191600101615312565b50909695505050505050565b6000806020838503121561535357600080fd5b82356001600160401b0381111561536957600080fd5b614b3f85828601614bf0565b6001600160601b0381168114610bfe57600080fd5b600082601f83011261539b57600080fd5b813560206153ab6150e983615020565b82815260069290921b840181019181810190868411156153ca57600080fd5b8286015b8481101561541b57604081890312156153e75760008081fd5b6153ef614c88565b81356153fa81614bae565b81528185013561540981615375565b818601528352918301916040016153ce565b509695505050505050565b600080600060a0848603121561543b57600080fd5b6154458585614f47565b9250606084013561545581615375565b915060808401356001600160401b0381111561547057600080fd5b61547c8682870161538a565b9150509250925092565b600082601f83011261549757600080fd5b813560206154a76150e983615020565b828152606092830285018201928282019190878511156154c657600080fd5b8387015b858110156154e9576154dc8982614f47565b84529284019281016154ca565b5090979650505050505050565b600082601f83011261550757600080fd5b813560206155176150e983615020565b82815260059290921b8401810191818101908684111561553657600080fd5b8286015b8481101561541b57803561554d81615375565b835291830191830161553a565b600082601f83011261556b57600080fd5b8135602061557b6150e983615020565b82815260059290921b8401810191818101908684111561559a57600080fd5b8286015b8481101561541b5780356001600160401b038111156155bd5760008081fd5b6155cb8986838b010161538a565b84525091830191830161559e565b600080600080600080600080610100898b0312156155f657600080fd5b6155ff89614bc3565b975061560d60208a01614bc3565b965061561b60408a01614bc3565b955061562960608a01614bc3565b94506080890135935060a08901356001600160401b038082111561564c57600080fd5b6156588c838d01615486565b945060c08b013591508082111561566e57600080fd5b61567a8c838d016154f6565b935060e08b013591508082111561569057600080fd5b5061569d8b828c0161555a565b9150509295985092959890939650565b60208101610f2d8284614ef1565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561573257615732615708565b5060010190565b60006020828403121561574b57600080fd5b81516126d481614bae565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156157b257600080fd5b815180151581146126d457600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60006040828403121561581c57600080fd5b604051604081018181106001600160401b038211171561583e5761583e614c4a565b604052823581526020928301359281019290925250919050565b6000808335601e1984360301811261586f57600080fd5b8301803591506001600160401b0382111561588957600080fd5b6020019150600581901b36038213156147a257600080fd5b6000602082840312156158b357600080fd5b81516126d481614b64565b600082198211156158d1576158d1615708565b500190565b600080858511156158e657600080fd5b838611156158f357600080fd5b5050820193919092039150565b600060c08201888352602060018060a01b03808a16828601526040898187015260c0606087015283895180865260e088019150848b01955060005b81811015615965578651805160ff168452860151851686840152958501959183019160010161593b565b505060808701989098525050505060a09091019190915250949350505050565b60006040828403121561599757600080fd5b6126d48383615043565b6000828210156159b3576159b3615708565b500390565b6000815180845260005b818110156159de576020818501810151868301820152016159c2565b818111156159f0576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b0384168152826020820152606060408201526000615a2c60608301846159b8565b95945050505050565b600060208284031215615a4757600080fd5b81516001600160c01b03811681146126d457600080fd5b600060208284031215615a7057600080fd5b5051919050565b6001600160a01b03841681526101a08101615a9f602083018580358252602090810135910152565b615ab9606083016040860180358252602090810135910152565b615ad360a083016080860180358252602090810135910152565b604060c0850160e08401376101208201600081526040610100860182375060006101608301908152835190526020909201516101809091015292915050565b60018060a01b0383168152604060208201526000825160606040840152615b3c60a08401826159b8565b90506020840151606084015260408401516080840152809150509392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0384168152604060208201819052600090615a2c9083018486615b5d565b60018060a01b0385168152836020820152606060408201526000611908606083018486615b5d565b600082601f830112615be457600080fd5b81516020615bf46150e983615020565b82815260059290921b84018101918181019086841115615c1357600080fd5b8286015b8481101561541b578051615c2a81615375565b8352918301918301615c17565b60008060408385031215615c4a57600080fd5b82516001600160401b0380821115615c6157600080fd5b615c6d86838701615bd3565b93506020850151915080821115615c8357600080fd5b50615c9085828601615bd3565b9150509250929050565b838152604060208201526000615a2c604083018486615b5d565b60006020808385031215615cc757600080fd5b82516001600160401b03811115615cdd57600080fd5b8301601f81018513615cee57600080fd5b8051615cfc6150e982615020565b81815260059190911b82018301908381019087831115615d1b57600080fd5b928401925b82841015615d42578351615d3381614b64565b82529284019290840190615d20565b979650505050505050565b600081518084526020808501945080840160005b83811015615d865781516001600160601b031687529582019590820190600101615d61565b509495945050505050565b6001600160a01b038416815260806020808301829052600091615db690840186615d4d565b9150604083018460005b6002811015615ddd57815183529183019190830190600101615dc0565b50505050949350505050565b6001600160a01b0383168152604060208201819052600090615e0d90830184615d4d565b949350505050565b634e487b7160e01b600052601260045260246000fd5b600082615e3a57615e3a615e15565b500690565b6001600160a01b0383168152604060208201819052600090615e0d908301846159b8565b828152604060208201526000615e0d60408301846159b8565b600060208284031215615e8e57600080fd5b81516126d481615375565b600060ff821660ff84168060ff03821115615eb657615eb6615708565b019392505050565b60006060820160ff8616835260206001600160601b03808716828601526040606081870152838751808652608088019150848901955060005b81811015615f2757865180516001600160a01b031684528601518516868401529585019591830191600101615ef7565b50909a9950505050505050505050565b600061ffff80831681811415615f4f57615f4f615708565b6001019392505050565b600060208284031215615f6b57600080fd5b81516001600160e01b0319811681146126d457600080fd5b60006001600160601b0380831681851681830481118215151615615fa957615fa9615708565b02949350505050565b60006001600160601b0380841680615fcc57615fcc615e15565b9216919091049291505056fe454f5265676973747279436f6f7264696e61746f722e5f64657265676973746530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47454f5265676973747279436f6f7264696e61746f722e5f76616c696461746543454f5265676973747279436f6f7264696e61746f722e7570646174654f706572a26469706673582212204c59a8de628371e49f7824b47a8a668875f8f052b5e9ec37e4a6f8f4e886209b64736f6c634300080c0033",
}

// ContractEORegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractEORegistryCoordinatorMetaData.ABI instead.
var ContractEORegistryCoordinatorABI = ContractEORegistryCoordinatorMetaData.ABI

// ContractEORegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractEORegistryCoordinatorMetaData.Bin instead.
var ContractEORegistryCoordinatorBin = ContractEORegistryCoordinatorMetaData.Bin

// DeployContractEORegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractEORegistryCoordinator to it.
func DeployContractEORegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _serviceManager common.Address, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address) (common.Address, *types.Transaction, *ContractEORegistryCoordinator, error) {
	parsed, err := ContractEORegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractEORegistryCoordinatorBin), backend, _serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractEORegistryCoordinator{ContractEORegistryCoordinatorCaller: ContractEORegistryCoordinatorCaller{contract: contract}, ContractEORegistryCoordinatorTransactor: ContractEORegistryCoordinatorTransactor{contract: contract}, ContractEORegistryCoordinatorFilterer: ContractEORegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractEORegistryCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractEORegistryCoordinatorMethods interface {
	ContractEORegistryCoordinatorCalls
	ContractEORegistryCoordinatorTransacts
	ContractEORegistryCoordinatorFilters
}

// ContractEORegistryCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractEORegistryCoordinatorCalls interface {
	EOChainManager(opts *bind.CallOpts) (common.Address, error)

	OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error)

	ChurnApprover(opts *bind.CallOpts) (common.Address, error)

	Ejector(opts *bind.CallOpts) (common.Address, error)

	GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetOperator(opts *bind.CallOpts, operator common.Address) (IEORegistryCoordinatorOperatorInfo, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IEORegistryCoordinatorOperatorSetParam, error)

	GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error)

	GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error)

	GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error)

	GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IEORegistryCoordinatorQuorumBitmapUpdate, error)

	IndexRegistry(opts *bind.CallOpts) (common.Address, error)

	IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	NumRegistries(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error)

	Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)

	ServiceManager(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractEORegistryCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractEORegistryCoordinatorTransacts interface {
	CreateQuorum(opts *bind.TransactOpts, operatorSetParams IEORegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IEOStakeRegistryStrategyParams) (*types.Transaction, error)

	DeregisterOperator(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error)

	EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IEORegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IEOStakeRegistryStrategyParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error)

	SetEOChainManager(opts *bind.TransactOpts, _EOChainManager common.Address) (*types.Transaction, error)

	SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IEORegistryCoordinatorOperatorSetParam) (*types.Transaction, error)

	SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error)

	UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractEORegistryCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractEORegistryCoordinatorFilters interface {
	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractEORegistryCoordinatorChurnApproverUpdated, error)

	FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorEjectorUpdatedIterator, error)
	WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorEjectorUpdated) (event.Subscription, error)
	ParseEjectorUpdated(log types.Log) (*ContractEORegistryCoordinatorEjectorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractEORegistryCoordinatorInitialized, error)

	FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractEORegistryCoordinatorOperatorDeregisteredIterator, error)
	WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorDeregistered(log types.Log) (*ContractEORegistryCoordinatorOperatorDeregistered, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractEORegistryCoordinatorOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractEORegistryCoordinatorOperatorRegistered, error)

	FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator, error)
	WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseOperatorSetParamsUpdated(log types.Log) (*ContractEORegistryCoordinatorOperatorSetParamsUpdated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractEORegistryCoordinatorOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractEORegistryCoordinatorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractEORegistryCoordinatorPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractEORegistryCoordinatorPaused, error)

	FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorPauserRegistrySetIterator, error)
	WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorPauserRegistrySet) (event.Subscription, error)
	ParsePauserRegistrySet(log types.Log) (*ContractEORegistryCoordinatorPauserRegistrySet, error)

	FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator, error)
	WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumBlockNumberUpdated(log types.Log) (*ContractEORegistryCoordinatorQuorumBlockNumberUpdated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractEORegistryCoordinatorUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractEORegistryCoordinatorUnpaused, error)
}

// ContractEORegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractEORegistryCoordinator struct {
	ContractEORegistryCoordinatorCaller     // Read-only binding to the contract
	ContractEORegistryCoordinatorTransactor // Write-only binding to the contract
	ContractEORegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractEORegistryCoordinator implements the ContractEORegistryCoordinatorMethods interface.
var _ ContractEORegistryCoordinatorMethods = (*ContractEORegistryCoordinator)(nil)

// ContractEORegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractEORegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEORegistryCoordinatorCaller implements the ContractEORegistryCoordinatorCalls interface.
var _ ContractEORegistryCoordinatorCalls = (*ContractEORegistryCoordinatorCaller)(nil)

// ContractEORegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractEORegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEORegistryCoordinatorTransactor implements the ContractEORegistryCoordinatorTransacts interface.
var _ ContractEORegistryCoordinatorTransacts = (*ContractEORegistryCoordinatorTransactor)(nil)

// ContractEORegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractEORegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEORegistryCoordinatorFilterer implements the ContractEORegistryCoordinatorFilters interface.
var _ ContractEORegistryCoordinatorFilters = (*ContractEORegistryCoordinatorFilterer)(nil)

// ContractEORegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractEORegistryCoordinatorSession struct {
	Contract     *ContractEORegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractEORegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractEORegistryCoordinatorCallerSession struct {
	Contract *ContractEORegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractEORegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractEORegistryCoordinatorTransactorSession struct {
	Contract     *ContractEORegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractEORegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractEORegistryCoordinatorRaw struct {
	Contract *ContractEORegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractEORegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractEORegistryCoordinatorCallerRaw struct {
	Contract *ContractEORegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractEORegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractEORegistryCoordinatorTransactorRaw struct {
	Contract *ContractEORegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractEORegistryCoordinator creates a new instance of ContractEORegistryCoordinator, bound to a specific deployed contract.
func NewContractEORegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractEORegistryCoordinator, error) {
	contract, err := bindContractEORegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinator{ContractEORegistryCoordinatorCaller: ContractEORegistryCoordinatorCaller{contract: contract}, ContractEORegistryCoordinatorTransactor: ContractEORegistryCoordinatorTransactor{contract: contract}, ContractEORegistryCoordinatorFilterer: ContractEORegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractEORegistryCoordinatorCaller creates a new read-only instance of ContractEORegistryCoordinator, bound to a specific deployed contract.
func NewContractEORegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractEORegistryCoordinatorCaller, error) {
	contract, err := bindContractEORegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractEORegistryCoordinatorTransactor creates a new write-only instance of ContractEORegistryCoordinator, bound to a specific deployed contract.
func NewContractEORegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractEORegistryCoordinatorTransactor, error) {
	contract, err := bindContractEORegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractEORegistryCoordinatorFilterer creates a new log filterer instance of ContractEORegistryCoordinator, bound to a specific deployed contract.
func NewContractEORegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractEORegistryCoordinatorFilterer, error) {
	contract, err := bindContractEORegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractEORegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractEORegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractEORegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractEORegistryCoordinator.Contract.ContractEORegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.ContractEORegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.ContractEORegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractEORegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// EOChainManager is a free data retrieval call binding the contract method 0x5b658e23.
//
// Solidity: function EOChainManager() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) EOChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "EOChainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EOChainManager is a free data retrieval call binding the contract method 0x5b658e23.
//
// Solidity: function EOChainManager() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) EOChainManager() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.EOChainManager(&_ContractEORegistryCoordinator.CallOpts)
}

// EOChainManager is a free data retrieval call binding the contract method 0x5b658e23.
//
// Solidity: function EOChainManager() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) EOChainManager() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.EOChainManager(&_ContractEORegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractEORegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractEORegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractEORegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractEORegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.BlsApkRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.BlsApkRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractEORegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractEORegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.ChurnApprover(&_ContractEORegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.ChurnApprover(&_ContractEORegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.Ejector(&_ContractEORegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.Ejector(&_ContractEORegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractEORegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractEORegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IEORegistryCoordinatorOperatorInfo, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IEORegistryCoordinatorOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEORegistryCoordinatorOperatorInfo)).(*IEORegistryCoordinatorOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetOperator(operator common.Address) (IEORegistryCoordinatorOperatorInfo, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperator(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetOperator(operator common.Address) (IEORegistryCoordinatorOperatorInfo, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperator(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorFromId(&_ContractEORegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorFromId(&_ContractEORegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorId(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorId(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IEORegistryCoordinatorOperatorSetParam, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IEORegistryCoordinatorOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IEORegistryCoordinatorOperatorSetParam)).(*IEORegistryCoordinatorOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (IEORegistryCoordinatorOperatorSetParam, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorSetParams(&_ContractEORegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (IEORegistryCoordinatorOperatorSetParam, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorSetParams(&_ContractEORegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorStatus(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractEORegistryCoordinator.Contract.GetOperatorStatus(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractEORegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractEORegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractEORegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractEORegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractEORegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractEORegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IEORegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(IEORegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IEORegistryCoordinatorQuorumBitmapUpdate)).(*IEORegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IEORegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractEORegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IEORegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractEORegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractEORegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.IndexRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.IndexRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractEORegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractEORegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractEORegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractEORegistryCoordinator.CallOpts, arg0)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.NumRegistries(&_ContractEORegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.NumRegistries(&_ContractEORegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.Owner(&_ContractEORegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.Owner(&_ContractEORegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractEORegistryCoordinator.Contract.Paused(&_ContractEORegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractEORegistryCoordinator.Contract.Paused(&_ContractEORegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.Paused0(&_ContractEORegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.Paused0(&_ContractEORegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.PauserRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.PauserRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractEORegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractEORegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractEORegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractEORegistryCoordinator.Contract.QuorumCount(&_ContractEORegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractEORegistryCoordinator.Contract.QuorumCount(&_ContractEORegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractEORegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractEORegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractEORegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.Registries(&_ContractEORegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.Registries(&_ContractEORegistryCoordinator.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) ServiceManager() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.ServiceManager(&_ContractEORegistryCoordinator.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) ServiceManager() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.ServiceManager(&_ContractEORegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEORegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.StakeRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractEORegistryCoordinator.Contract.StakeRegistry(&_ContractEORegistryCoordinator.CallOpts)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) CreateQuorum(opts *bind.TransactOpts, operatorSetParams IEORegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IEOStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "createQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) CreateQuorum(operatorSetParams IEORegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IEOStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.CreateQuorum(&_ContractEORegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) CreateQuorum(operatorSetParams IEORegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IEOStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.CreateQuorum(&_ContractEORegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "deregisterOperator", quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) DeregisterOperator(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.DeregisterOperator(&_ContractEORegistryCoordinator.TransactOpts, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) DeregisterOperator(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.DeregisterOperator(&_ContractEORegistryCoordinator.TransactOpts, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.EjectOperator(&_ContractEORegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.EjectOperator(&_ContractEORegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xdd8283f3.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IEORegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IEOStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xdd8283f3.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IEORegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IEOStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.Initialize(&_ContractEORegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xdd8283f3.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IEORegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IEOStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.Initialize(&_ContractEORegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.Pause(&_ContractEORegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.Pause(&_ContractEORegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.PauseAll(&_ContractEORegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.PauseAll(&_ContractEORegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x1615700a.
//
// Solidity: function registerOperator(bytes quorumNumbers, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "registerOperator", quorumNumbers, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x1615700a.
//
// Solidity: function registerOperator(bytes quorumNumbers, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) RegisterOperator(quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.RegisterOperator(&_ContractEORegistryCoordinator.TransactOpts, quorumNumbers, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x1615700a.
//
// Solidity: function registerOperator(bytes quorumNumbers, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) RegisterOperator(quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.RegisterOperator(&_ContractEORegistryCoordinator.TransactOpts, quorumNumbers, params, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x993252b4.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "registerOperatorWithChurn", quorumNumbers, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x993252b4.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) RegisterOperatorWithChurn(quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractEORegistryCoordinator.TransactOpts, quorumNumbers, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x993252b4.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) RegisterOperatorWithChurn(quorumNumbers []byte, params IEOBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IEORegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractEORegistryCoordinator.TransactOpts, quorumNumbers, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.RenounceOwnership(&_ContractEORegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.RenounceOwnership(&_ContractEORegistryCoordinator.TransactOpts)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetChurnApprover(&_ContractEORegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetChurnApprover(&_ContractEORegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEOChainManager is a paid mutator transaction binding the contract method 0xa65f39e1.
//
// Solidity: function setEOChainManager(address _EOChainManager) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) SetEOChainManager(opts *bind.TransactOpts, _EOChainManager common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "setEOChainManager", _EOChainManager)
}

// SetEOChainManager is a paid mutator transaction binding the contract method 0xa65f39e1.
//
// Solidity: function setEOChainManager(address _EOChainManager) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) SetEOChainManager(_EOChainManager common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetEOChainManager(&_ContractEORegistryCoordinator.TransactOpts, _EOChainManager)
}

// SetEOChainManager is a paid mutator transaction binding the contract method 0xa65f39e1.
//
// Solidity: function setEOChainManager(address _EOChainManager) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) SetEOChainManager(_EOChainManager common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetEOChainManager(&_ContractEORegistryCoordinator.TransactOpts, _EOChainManager)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetEjector(&_ContractEORegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetEjector(&_ContractEORegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IEORegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IEORegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetOperatorSetParams(&_ContractEORegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IEORegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetOperatorSetParams(&_ContractEORegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetPauserRegistry(&_ContractEORegistryCoordinator.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.SetPauserRegistry(&_ContractEORegistryCoordinator.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.TransferOwnership(&_ContractEORegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.TransferOwnership(&_ContractEORegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.Unpause(&_ContractEORegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.Unpause(&_ContractEORegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.UpdateOperators(&_ContractEORegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.UpdateOperators(&_ContractEORegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractEORegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractEORegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractEORegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// ContractEORegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractEORegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorChurnApproverUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorChurnApproverUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractEORegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorChurnApproverUpdated)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractEORegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractEORegistryCoordinatorChurnApproverUpdated)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractEORegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorEjectorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorEjectorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorEjectorUpdatedIterator{contract: _ContractEORegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorEjectorUpdated)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractEORegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractEORegistryCoordinatorEjectorUpdated)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorInitializedIterator struct {
	Event *ContractEORegistryCoordinatorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorInitialized represents a Initialized event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorInitializedIterator{contract: _ContractEORegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorInitialized)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractEORegistryCoordinatorInitialized, error) {
	event := new(ContractEORegistryCoordinatorInitialized)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractEORegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorOperatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorOperatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractEORegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractEORegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorOperatorDeregistered)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractEORegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractEORegistryCoordinatorOperatorDeregistered)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractEORegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractEORegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorOperatorRegisteredIterator{contract: _ContractEORegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorOperatorRegistered)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractEORegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractEORegistryCoordinatorOperatorRegistered)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractEORegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorOperatorSetParamsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorOperatorSetParamsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IEORegistryCoordinatorOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractEORegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractEORegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractEORegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractEORegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractEORegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorOwnershipTransferredIterator{contract: _ContractEORegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorOwnershipTransferred)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractEORegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractEORegistryCoordinatorOwnershipTransferred)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorPausedIterator struct {
	Event *ContractEORegistryCoordinatorPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorPaused represents a Paused event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractEORegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorPausedIterator{contract: _ContractEORegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorPaused)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractEORegistryCoordinatorPaused, error) {
	event := new(ContractEORegistryCoordinatorPaused)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorPauserRegistrySetIterator struct {
	Event *ContractEORegistryCoordinatorPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorPauserRegistrySet represents a PauserRegistrySet event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractEORegistryCoordinatorPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorPauserRegistrySetIterator{contract: _ContractEORegistryCoordinator.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorPauserRegistrySet)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParsePauserRegistrySet(log types.Log) (*ContractEORegistryCoordinatorPauserRegistrySet, error) {
	event := new(ContractEORegistryCoordinatorPauserRegistrySet)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractEORegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorQuorumBlockNumberUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorQuorumBlockNumberUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractEORegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractEORegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractEORegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEORegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorUnpausedIterator struct {
	Event *ContractEORegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEORegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEORegistryCoordinatorUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEORegistryCoordinatorUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEORegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEORegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEORegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractEORegistryCoordinator contract.
type ContractEORegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractEORegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractEORegistryCoordinatorUnpausedIterator{contract: _ContractEORegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractEORegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEORegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEORegistryCoordinatorUnpaused)
				if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractEORegistryCoordinator *ContractEORegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractEORegistryCoordinatorUnpaused, error) {
	event := new(ContractEORegistryCoordinatorUnpaused)
	if err := _ContractEORegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
