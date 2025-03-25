package luban

import _ "embed"

// contract codes for Feline upgrade
var (
	//go:embed feline/ValidatorContract
	FelineValidatorContract string
	//go:embed feline/SlashContract
	FelineSlashContract string
	//go:embed feline/SystemRewardContract
	FelineSystemRewardContract string
	//go:embed feline/RelayerHubContract
	FelineRelayerHubContract string
	//go:embed feline/CrossChainContract
	FelineCrossChainContract string
)

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/ValidatorContract
	MainnetValidatorContract string
	//go:embed mainnet/SlashContract
	MainnetSlashContract string
	//go:embed mainnet/SystemRewardContract
	MainnetSystemRewardContract string
	//go:embed mainnet/RelayerHubContract
	MainnetRelayerHubContract string
	//go:embed mainnet/CrossChainContract
	MainnetCrossChainContract string
)
