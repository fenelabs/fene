package planck

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/SlashContract
	MainnetSlashContract string
	//go:embed mainnet/TokenHubContract
	MainnetTokenHubContract string
	//go:embed mainnet/CrossChainContract
	MainnetCrossChainContract string
)

// contract codes for feline upgrade
var (
	//go:embed feline/SlashContract
	FelineSlashContract string
	//go:embed feline/TokenHubContract
	FelineTokenHubContract string
	//go:embed feline/CrossChainContract
	FelineCrossChainContract string
	//go:embed feline/StakingContract
	FelineStakingContract string
)
