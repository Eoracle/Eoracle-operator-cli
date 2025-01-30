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

// IBLSApkRegistryTypesPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	ChainValidatorSignature     BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ISlashingRegistryCoordinatorTypesOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// ISlashingRegistryCoordinatorTypesOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// ISlashingRegistryCoordinatorTypesOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// IStakeRegistryTypesStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryTypesStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// EORegistryCoordinatorMetaData contains all meta data concerning the EORegistryCoordinator contract.
var EORegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIEOStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIEOBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIEOIndexRegistry\"},{\"name\":\"_socketRegistry\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOChainManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createSlashableStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTotalDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableM2QuorumRegistration\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableOperatorSets\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_accountIdentifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isM2Quorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"m2QuorumsDisabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"chainValidatorSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"chainValidatorSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAccountIdentifier\",\"inputs\":[{\"name\":\"_accountIdentifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainManager\",\"inputs\":[{\"name\":\"newChainManager\",\"type\":\"address\",\"internalType\":\"contractIEOChainManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"socketRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainManagerUpdated\",\"inputs\":[{\"name\":\"prevChainManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChainManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"M2QuorumsDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetsEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegisteredForQuorums\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapUpdateIsAfterBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotChurnSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotKickOperatorAboveThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotReregisterYet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChurnApproverSaltUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientStakeForChurn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRegistrationType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"M2QuorumsAlreadyDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxQuorumsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NextBitmapUpdateIsBeforeBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegisteredForQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEjector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsAlreadyEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumOperatorCountMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]}]",
}

// EORegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use EORegistryCoordinatorMetaData.ABI instead.
var EORegistryCoordinatorABI = EORegistryCoordinatorMetaData.ABI

// EORegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type EORegistryCoordinator struct {
	EORegistryCoordinatorCaller     // Read-only binding to the contract
	EORegistryCoordinatorTransactor // Write-only binding to the contract
	EORegistryCoordinatorFilterer   // Log filterer for contract events
}

// EORegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type EORegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EORegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EORegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EORegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EORegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EORegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EORegistryCoordinatorSession struct {
	Contract     *EORegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EORegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EORegistryCoordinatorCallerSession struct {
	Contract *EORegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// EORegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EORegistryCoordinatorTransactorSession struct {
	Contract     *EORegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// EORegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type EORegistryCoordinatorRaw struct {
	Contract *EORegistryCoordinator // Generic contract binding to access the raw methods on
}

// EORegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EORegistryCoordinatorCallerRaw struct {
	Contract *EORegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// EORegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EORegistryCoordinatorTransactorRaw struct {
	Contract *EORegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEORegistryCoordinator creates a new instance of EORegistryCoordinator, bound to a specific deployed contract.
func NewEORegistryCoordinator(address common.Address, backend bind.ContractBackend) (*EORegistryCoordinator, error) {
	contract, err := bindEORegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinator{EORegistryCoordinatorCaller: EORegistryCoordinatorCaller{contract: contract}, EORegistryCoordinatorTransactor: EORegistryCoordinatorTransactor{contract: contract}, EORegistryCoordinatorFilterer: EORegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewEORegistryCoordinatorCaller creates a new read-only instance of EORegistryCoordinator, bound to a specific deployed contract.
func NewEORegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*EORegistryCoordinatorCaller, error) {
	contract, err := bindEORegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorCaller{contract: contract}, nil
}

// NewEORegistryCoordinatorTransactor creates a new write-only instance of EORegistryCoordinator, bound to a specific deployed contract.
func NewEORegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*EORegistryCoordinatorTransactor, error) {
	contract, err := bindEORegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorTransactor{contract: contract}, nil
}

// NewEORegistryCoordinatorFilterer creates a new log filterer instance of EORegistryCoordinator, bound to a specific deployed contract.
func NewEORegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*EORegistryCoordinatorFilterer, error) {
	contract, err := bindEORegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorFilterer{contract: contract}, nil
}

// bindEORegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindEORegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EORegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EORegistryCoordinator *EORegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EORegistryCoordinator.Contract.EORegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EORegistryCoordinator *EORegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.EORegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EORegistryCoordinator *EORegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.EORegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EORegistryCoordinator *EORegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EORegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _EORegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_EORegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _EORegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_EORegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _EORegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_EORegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _EORegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_EORegistryCoordinator.CallOpts)
}

// AccountIdentifier is a free data retrieval call binding the contract method 0x0764cb93.
//
// Solidity: function accountIdentifier() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) AccountIdentifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "accountIdentifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountIdentifier is a free data retrieval call binding the contract method 0x0764cb93.
//
// Solidity: function accountIdentifier() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) AccountIdentifier() (common.Address, error) {
	return _EORegistryCoordinator.Contract.AccountIdentifier(&_EORegistryCoordinator.CallOpts)
}

// AccountIdentifier is a free data retrieval call binding the contract method 0x0764cb93.
//
// Solidity: function accountIdentifier() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) AccountIdentifier() (common.Address, error) {
	return _EORegistryCoordinator.Contract.AccountIdentifier(&_EORegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) AllocationManager() (common.Address, error) {
	return _EORegistryCoordinator.Contract.AllocationManager(&_EORegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) AllocationManager() (common.Address, error) {
	return _EORegistryCoordinator.Contract.AllocationManager(&_EORegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.BlsApkRegistry(&_EORegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.BlsApkRegistry(&_EORegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _EORegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_EORegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _EORegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_EORegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) ChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "chainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) ChainManager() (common.Address, error) {
	return _EORegistryCoordinator.Contract.ChainManager(&_EORegistryCoordinator.CallOpts)
}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) ChainManager() (common.Address, error) {
	return _EORegistryCoordinator.Contract.ChainManager(&_EORegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _EORegistryCoordinator.Contract.ChurnApprover(&_EORegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _EORegistryCoordinator.Contract.ChurnApprover(&_EORegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) EjectionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "ejectionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) EjectionCooldown() (*big.Int, error) {
	return _EORegistryCoordinator.Contract.EjectionCooldown(&_EORegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) EjectionCooldown() (*big.Int, error) {
	return _EORegistryCoordinator.Contract.EjectionCooldown(&_EORegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _EORegistryCoordinator.Contract.Ejector(&_EORegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _EORegistryCoordinator.Contract.Ejector(&_EORegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_EORegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_EORegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesOperatorInfo)).(*ISlashingRegistryCoordinatorTypesOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetOperator(operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	return _EORegistryCoordinator.Contract.GetOperator(&_EORegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetOperator(operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	return _EORegistryCoordinator.Contract.GetOperator(&_EORegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _EORegistryCoordinator.Contract.GetOperatorFromId(&_EORegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _EORegistryCoordinator.Contract.GetOperatorFromId(&_EORegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _EORegistryCoordinator.Contract.GetOperatorId(&_EORegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _EORegistryCoordinator.Contract.GetOperatorId(&_EORegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesOperatorSetParam)).(*ISlashingRegistryCoordinatorTypesOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	return _EORegistryCoordinator.Contract.GetOperatorSetParams(&_EORegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	return _EORegistryCoordinator.Contract.GetOperatorSetParams(&_EORegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _EORegistryCoordinator.Contract.GetOperatorStatus(&_EORegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _EORegistryCoordinator.Contract.GetOperatorStatus(&_EORegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_EORegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_EORegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_EORegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_EORegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_EORegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_EORegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate)).(*ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_EORegistryCoordinator *EORegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_EORegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	return _EORegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_EORegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.IndexRegistry(&_EORegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.IndexRegistry(&_EORegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _EORegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_EORegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _EORegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_EORegistryCoordinator.CallOpts, arg0)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 quorumNumber) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) IsM2Quorum(opts *bind.CallOpts, quorumNumber uint8) (bool, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "isM2Quorum", quorumNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 quorumNumber) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) IsM2Quorum(quorumNumber uint8) (bool, error) {
	return _EORegistryCoordinator.Contract.IsM2Quorum(&_EORegistryCoordinator.CallOpts, quorumNumber)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 quorumNumber) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) IsM2Quorum(quorumNumber uint8) (bool, error) {
	return _EORegistryCoordinator.Contract.IsM2Quorum(&_EORegistryCoordinator.CallOpts, quorumNumber)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "lastEjectionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.LastEjectionTimestamp(&_EORegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.LastEjectionTimestamp(&_EORegistryCoordinator.CallOpts, arg0)
}

// M2QuorumsDisabled is a free data retrieval call binding the contract method 0xb2d8678d.
//
// Solidity: function m2QuorumsDisabled() view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) M2QuorumsDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "m2QuorumsDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// M2QuorumsDisabled is a free data retrieval call binding the contract method 0xb2d8678d.
//
// Solidity: function m2QuorumsDisabled() view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) M2QuorumsDisabled() (bool, error) {
	return _EORegistryCoordinator.Contract.M2QuorumsDisabled(&_EORegistryCoordinator.CallOpts)
}

// M2QuorumsDisabled is a free data retrieval call binding the contract method 0xb2d8678d.
//
// Solidity: function m2QuorumsDisabled() view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) M2QuorumsDisabled() (bool, error) {
	return _EORegistryCoordinator.Contract.M2QuorumsDisabled(&_EORegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _EORegistryCoordinator.Contract.NumRegistries(&_EORegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _EORegistryCoordinator.Contract.NumRegistries(&_EORegistryCoordinator.CallOpts)
}

// OperatorSetsEnabled is a free data retrieval call binding the contract method 0x81f936d2.
//
// Solidity: function operatorSetsEnabled() view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) OperatorSetsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "operatorSetsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSetsEnabled is a free data retrieval call binding the contract method 0x81f936d2.
//
// Solidity: function operatorSetsEnabled() view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) OperatorSetsEnabled() (bool, error) {
	return _EORegistryCoordinator.Contract.OperatorSetsEnabled(&_EORegistryCoordinator.CallOpts)
}

// OperatorSetsEnabled is a free data retrieval call binding the contract method 0x81f936d2.
//
// Solidity: function operatorSetsEnabled() view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) OperatorSetsEnabled() (bool, error) {
	return _EORegistryCoordinator.Contract.OperatorSetsEnabled(&_EORegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Owner() (common.Address, error) {
	return _EORegistryCoordinator.Contract.Owner(&_EORegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _EORegistryCoordinator.Contract.Owner(&_EORegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _EORegistryCoordinator.Contract.Paused(&_EORegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _EORegistryCoordinator.Contract.Paused(&_EORegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _EORegistryCoordinator.Contract.Paused0(&_EORegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _EORegistryCoordinator.Contract.Paused0(&_EORegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.PauserRegistry(&_EORegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.PauserRegistry(&_EORegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_EORegistryCoordinator *EORegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _EORegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_EORegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _EORegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_EORegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _EORegistryCoordinator.Contract.QuorumCount(&_EORegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _EORegistryCoordinator.Contract.QuorumCount(&_EORegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_EORegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _EORegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_EORegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _EORegistryCoordinator.Contract.Registries(&_EORegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _EORegistryCoordinator.Contract.Registries(&_EORegistryCoordinator.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) ServiceManager() (common.Address, error) {
	return _EORegistryCoordinator.Contract.ServiceManager(&_EORegistryCoordinator.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) ServiceManager() (common.Address, error) {
	return _EORegistryCoordinator.Contract.ServiceManager(&_EORegistryCoordinator.CallOpts)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) SocketRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "socketRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SocketRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.SocketRegistry(&_EORegistryCoordinator.CallOpts)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) SocketRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.SocketRegistry(&_EORegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EORegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.StakeRegistry(&_EORegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_EORegistryCoordinator *EORegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _EORegistryCoordinator.Contract.StakeRegistry(&_EORegistryCoordinator.CallOpts)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "createSlashableStakeQuorum", operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) CreateSlashableStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_EORegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) CreateSlashableStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_EORegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "createTotalDelegatedStakeQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_EORegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_EORegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "deregisterOperator", operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.DeregisterOperator(&_EORegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.DeregisterOperator(&_EORegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) DeregisterOperator0(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "deregisterOperator0", quorumNumbers)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) DeregisterOperator0(quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.DeregisterOperator0(&_EORegistryCoordinator.TransactOpts, quorumNumbers)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) DeregisterOperator0(quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.DeregisterOperator0(&_EORegistryCoordinator.TransactOpts, quorumNumbers)
}

// DisableM2QuorumRegistration is a paid mutator transaction binding the contract method 0x733b7507.
//
// Solidity: function disableM2QuorumRegistration() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) DisableM2QuorumRegistration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "disableM2QuorumRegistration")
}

// DisableM2QuorumRegistration is a paid mutator transaction binding the contract method 0x733b7507.
//
// Solidity: function disableM2QuorumRegistration() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) DisableM2QuorumRegistration() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.DisableM2QuorumRegistration(&_EORegistryCoordinator.TransactOpts)
}

// DisableM2QuorumRegistration is a paid mutator transaction binding the contract method 0x733b7507.
//
// Solidity: function disableM2QuorumRegistration() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) DisableM2QuorumRegistration() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.DisableM2QuorumRegistration(&_EORegistryCoordinator.TransactOpts)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.EjectOperator(&_EORegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.EjectOperator(&_EORegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) EnableOperatorSets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "enableOperatorSets")
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) EnableOperatorSets() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.EnableOperatorSets(&_EORegistryCoordinator.TransactOpts)
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) EnableOperatorSets() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.EnableOperatorSets(&_EORegistryCoordinator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _accountIdentifier) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _accountIdentifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _accountIdentifier) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.Initialize(&_EORegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _accountIdentifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _accountIdentifier) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.Initialize(&_EORegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _accountIdentifier)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.Pause(&_EORegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.Pause(&_EORegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.PauseAll(&_EORegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.PauseAll(&_EORegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "registerOperator", operator, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) RegisterOperator(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RegisterOperator(&_EORegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) RegisterOperator(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RegisterOperator(&_EORegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xf17c5168.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) RegisterOperator0(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryTypesPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "registerOperator0", quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xf17c5168.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) RegisterOperator0(quorumNumbers []byte, socket string, params IBLSApkRegistryTypesPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RegisterOperator0(&_EORegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xf17c5168.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) RegisterOperator0(quorumNumbers []byte, socket string, params IBLSApkRegistryTypesPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RegisterOperator0(&_EORegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x5452d7a2.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryTypesPubkeyRegistrationParams, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "registerOperatorWithChurn", quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x5452d7a2.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryTypesPubkeyRegistrationParams, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RegisterOperatorWithChurn(&_EORegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x5452d7a2.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryTypesPubkeyRegistrationParams, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RegisterOperatorWithChurn(&_EORegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RenounceOwnership(&_EORegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.RenounceOwnership(&_EORegistryCoordinator.TransactOpts)
}

// SetAccountIdentifier is a paid mutator transaction binding the contract method 0x143e5915.
//
// Solidity: function setAccountIdentifier(address _accountIdentifier) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) SetAccountIdentifier(opts *bind.TransactOpts, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "setAccountIdentifier", _accountIdentifier)
}

// SetAccountIdentifier is a paid mutator transaction binding the contract method 0x143e5915.
//
// Solidity: function setAccountIdentifier(address _accountIdentifier) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SetAccountIdentifier(_accountIdentifier common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetAccountIdentifier(&_EORegistryCoordinator.TransactOpts, _accountIdentifier)
}

// SetAccountIdentifier is a paid mutator transaction binding the contract method 0x143e5915.
//
// Solidity: function setAccountIdentifier(address _accountIdentifier) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) SetAccountIdentifier(_accountIdentifier common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetAccountIdentifier(&_EORegistryCoordinator.TransactOpts, _accountIdentifier)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address newChainManager) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) SetChainManager(opts *bind.TransactOpts, newChainManager common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "setChainManager", newChainManager)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address newChainManager) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SetChainManager(newChainManager common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetChainManager(&_EORegistryCoordinator.TransactOpts, newChainManager)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address newChainManager) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) SetChainManager(newChainManager common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetChainManager(&_EORegistryCoordinator.TransactOpts, newChainManager)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetChurnApprover(&_EORegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetChurnApprover(&_EORegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "setEjectionCooldown", _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetEjectionCooldown(&_EORegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetEjectionCooldown(&_EORegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetEjector(&_EORegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetEjector(&_EORegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetOperatorSetParams(&_EORegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.SetOperatorSetParams(&_EORegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.TransferOwnership(&_EORegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.TransferOwnership(&_EORegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.Unpause(&_EORegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.Unpause(&_EORegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.UpdateOperators(&_EORegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.UpdateOperators(&_EORegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_EORegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_EORegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _EORegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.UpdateSocket(&_EORegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_EORegistryCoordinator *EORegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _EORegistryCoordinator.Contract.UpdateSocket(&_EORegistryCoordinator.TransactOpts, socket)
}

// EORegistryCoordinatorChainManagerUpdatedIterator is returned from FilterChainManagerUpdated and is used to iterate over the raw logs and unpacked data for ChainManagerUpdated events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorChainManagerUpdatedIterator struct {
	Event *EORegistryCoordinatorChainManagerUpdated // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorChainManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorChainManagerUpdated)
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
		it.Event = new(EORegistryCoordinatorChainManagerUpdated)
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
func (it *EORegistryCoordinatorChainManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorChainManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorChainManagerUpdated represents a ChainManagerUpdated event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorChainManagerUpdated struct {
	PrevChainManager common.Address
	NewChainManager  common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChainManagerUpdated is a free log retrieval operation binding the contract event 0x82dba664578316cd747b8d55d907a9a79c86cc87e5f679cab0e94f1463e8979a.
//
// Solidity: event ChainManagerUpdated(address prevChainManager, address newChainManager)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterChainManagerUpdated(opts *bind.FilterOpts) (*EORegistryCoordinatorChainManagerUpdatedIterator, error) {

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "ChainManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorChainManagerUpdatedIterator{contract: _EORegistryCoordinator.contract, event: "ChainManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchChainManagerUpdated is a free log subscription operation binding the contract event 0x82dba664578316cd747b8d55d907a9a79c86cc87e5f679cab0e94f1463e8979a.
//
// Solidity: event ChainManagerUpdated(address prevChainManager, address newChainManager)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchChainManagerUpdated(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorChainManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "ChainManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorChainManagerUpdated)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "ChainManagerUpdated", log); err != nil {
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

// ParseChainManagerUpdated is a log parse operation binding the contract event 0x82dba664578316cd747b8d55d907a9a79c86cc87e5f679cab0e94f1463e8979a.
//
// Solidity: event ChainManagerUpdated(address prevChainManager, address newChainManager)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseChainManagerUpdated(log types.Log) (*EORegistryCoordinatorChainManagerUpdated, error) {
	event := new(EORegistryCoordinatorChainManagerUpdated)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "ChainManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *EORegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(EORegistryCoordinatorChurnApproverUpdated)
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
func (it *EORegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*EORegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorChurnApproverUpdatedIterator{contract: _EORegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorChurnApproverUpdated)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*EORegistryCoordinatorChurnApproverUpdated, error) {
	event := new(EORegistryCoordinatorChurnApproverUpdated)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorEjectorUpdatedIterator struct {
	Event *EORegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorEjectorUpdated)
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
		it.Event = new(EORegistryCoordinatorEjectorUpdated)
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
func (it *EORegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*EORegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorEjectorUpdatedIterator{contract: _EORegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorEjectorUpdated)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*EORegistryCoordinatorEjectorUpdated, error) {
	event := new(EORegistryCoordinatorEjectorUpdated)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorInitializedIterator struct {
	Event *EORegistryCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorInitialized)
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
		it.Event = new(EORegistryCoordinatorInitialized)
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
func (it *EORegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorInitialized represents a Initialized event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*EORegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorInitializedIterator{contract: _EORegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorInitialized)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*EORegistryCoordinatorInitialized, error) {
	event := new(EORegistryCoordinatorInitialized)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorM2QuorumsDisabledIterator is returned from FilterM2QuorumsDisabled and is used to iterate over the raw logs and unpacked data for M2QuorumsDisabled events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorM2QuorumsDisabledIterator struct {
	Event *EORegistryCoordinatorM2QuorumsDisabled // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorM2QuorumsDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorM2QuorumsDisabled)
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
		it.Event = new(EORegistryCoordinatorM2QuorumsDisabled)
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
func (it *EORegistryCoordinatorM2QuorumsDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorM2QuorumsDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorM2QuorumsDisabled represents a M2QuorumsDisabled event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorM2QuorumsDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterM2QuorumsDisabled is a free log retrieval operation binding the contract event 0xa4cd42920ed0d1372ba4051d4577279f236fbbe677a67f3f7d645e82425dd98d.
//
// Solidity: event M2QuorumsDisabled()
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterM2QuorumsDisabled(opts *bind.FilterOpts) (*EORegistryCoordinatorM2QuorumsDisabledIterator, error) {

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "M2QuorumsDisabled")
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorM2QuorumsDisabledIterator{contract: _EORegistryCoordinator.contract, event: "M2QuorumsDisabled", logs: logs, sub: sub}, nil
}

// WatchM2QuorumsDisabled is a free log subscription operation binding the contract event 0xa4cd42920ed0d1372ba4051d4577279f236fbbe677a67f3f7d645e82425dd98d.
//
// Solidity: event M2QuorumsDisabled()
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchM2QuorumsDisabled(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorM2QuorumsDisabled) (event.Subscription, error) {

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "M2QuorumsDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorM2QuorumsDisabled)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "M2QuorumsDisabled", log); err != nil {
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

// ParseM2QuorumsDisabled is a log parse operation binding the contract event 0xa4cd42920ed0d1372ba4051d4577279f236fbbe677a67f3f7d645e82425dd98d.
//
// Solidity: event M2QuorumsDisabled()
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseM2QuorumsDisabled(log types.Log) (*EORegistryCoordinatorM2QuorumsDisabled, error) {
	event := new(EORegistryCoordinatorM2QuorumsDisabled)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "M2QuorumsDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *EORegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(EORegistryCoordinatorOperatorDeregistered)
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
func (it *EORegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*EORegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorOperatorDeregisteredIterator{contract: _EORegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorOperatorDeregistered)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*EORegistryCoordinatorOperatorDeregistered, error) {
	event := new(EORegistryCoordinatorOperatorDeregistered)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorRegisteredIterator struct {
	Event *EORegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorOperatorRegistered)
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
		it.Event = new(EORegistryCoordinatorOperatorRegistered)
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
func (it *EORegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*EORegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorOperatorRegisteredIterator{contract: _EORegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorOperatorRegistered)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*EORegistryCoordinatorOperatorRegistered, error) {
	event := new(EORegistryCoordinatorOperatorRegistered)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *EORegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(EORegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *EORegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*EORegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _EORegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorOperatorSetParamsUpdated)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*EORegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(EORegistryCoordinatorOperatorSetParamsUpdated)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorOperatorSetsEnabledIterator is returned from FilterOperatorSetsEnabled and is used to iterate over the raw logs and unpacked data for OperatorSetsEnabled events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorSetsEnabledIterator struct {
	Event *EORegistryCoordinatorOperatorSetsEnabled // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorOperatorSetsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorOperatorSetsEnabled)
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
		it.Event = new(EORegistryCoordinatorOperatorSetsEnabled)
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
func (it *EORegistryCoordinatorOperatorSetsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorOperatorSetsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorOperatorSetsEnabled represents a OperatorSetsEnabled event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorSetsEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetsEnabled is a free log retrieval operation binding the contract event 0x0b88306ff4627121f5b3e5b1c5f88f6b1e42fd2c0478ef1c91662d49d1f07755.
//
// Solidity: event OperatorSetsEnabled()
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterOperatorSetsEnabled(opts *bind.FilterOpts) (*EORegistryCoordinatorOperatorSetsEnabledIterator, error) {

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "OperatorSetsEnabled")
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorOperatorSetsEnabledIterator{contract: _EORegistryCoordinator.contract, event: "OperatorSetsEnabled", logs: logs, sub: sub}, nil
}

// WatchOperatorSetsEnabled is a free log subscription operation binding the contract event 0x0b88306ff4627121f5b3e5b1c5f88f6b1e42fd2c0478ef1c91662d49d1f07755.
//
// Solidity: event OperatorSetsEnabled()
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchOperatorSetsEnabled(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorOperatorSetsEnabled) (event.Subscription, error) {

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "OperatorSetsEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorOperatorSetsEnabled)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorSetsEnabled", log); err != nil {
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

// ParseOperatorSetsEnabled is a log parse operation binding the contract event 0x0b88306ff4627121f5b3e5b1c5f88f6b1e42fd2c0478ef1c91662d49d1f07755.
//
// Solidity: event OperatorSetsEnabled()
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseOperatorSetsEnabled(log types.Log) (*EORegistryCoordinatorOperatorSetsEnabled, error) {
	event := new(EORegistryCoordinatorOperatorSetsEnabled)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorSetsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *EORegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorOperatorSocketUpdate)
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
		it.Event = new(EORegistryCoordinatorOperatorSocketUpdate)
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
func (it *EORegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*EORegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorOperatorSocketUpdateIterator{contract: _EORegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorOperatorSocketUpdate)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*EORegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(EORegistryCoordinatorOperatorSocketUpdate)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOwnershipTransferredIterator struct {
	Event *EORegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorOwnershipTransferred)
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
		it.Event = new(EORegistryCoordinatorOwnershipTransferred)
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
func (it *EORegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EORegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorOwnershipTransferredIterator{contract: _EORegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorOwnershipTransferred)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*EORegistryCoordinatorOwnershipTransferred, error) {
	event := new(EORegistryCoordinatorOwnershipTransferred)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorPausedIterator struct {
	Event *EORegistryCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorPaused)
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
		it.Event = new(EORegistryCoordinatorPaused)
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
func (it *EORegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorPaused represents a Paused event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*EORegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorPausedIterator{contract: _EORegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorPaused)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParsePaused(log types.Log) (*EORegistryCoordinatorPaused, error) {
	event := new(EORegistryCoordinatorPaused)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *EORegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(EORegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *EORegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*EORegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _EORegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*EORegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(EORegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EORegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorUnpausedIterator struct {
	Event *EORegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *EORegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EORegistryCoordinatorUnpaused)
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
		it.Event = new(EORegistryCoordinatorUnpaused)
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
func (it *EORegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EORegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EORegistryCoordinatorUnpaused represents a Unpaused event raised by the EORegistryCoordinator contract.
type EORegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*EORegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EORegistryCoordinatorUnpausedIterator{contract: _EORegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EORegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EORegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EORegistryCoordinatorUnpaused)
				if err := _EORegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EORegistryCoordinator *EORegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*EORegistryCoordinatorUnpaused, error) {
	event := new(EORegistryCoordinatorUnpaused)
	if err := _EORegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
