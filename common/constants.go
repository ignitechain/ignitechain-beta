package common

const (
	FoundationSmartContract        = "0x0000000000000000000000000000000000000021"
	SystemRewardSmartContract      = "0x0000000000000000000000000000000000000051"
	IgniteDeployerContract         = "0x0000000000000000000000000000000000000101"
	MultiSigTimeLockContract       = "0x0000000000000000000000000000000000000151"
	MultiSigWalletContract         = "0x0000000000000000000000000000000000000201"
	ValidatorsContractAddress      = "0x0000000000000000000000000000000000000251"
	DevRewardContractAddress       = "0x0000000000000000000000000000000000000501"
	DelegatorRewardContractAddress = "0x0000000000000000000000000000000000001001"
	IgniteBodyContractAddress      = "0x0000000000000000000000000000000000002001"

	BlockRewardToValidator = 50
	BlockRewardToSystem    = 10
	BlockRewardToDevs      = 7

	LowGasPrice    = 5000000000
	MediumGasPrice = 9000000000
)

var BlackList = map[Address]bool{}
