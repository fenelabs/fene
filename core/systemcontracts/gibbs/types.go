package gibbs

import _ "embed"

// contract codes for Feline upgrade
var (
	//go:embed feline/TokenHubContract
	FelineTokenHubContract string
	//go:embed feline/StakingContract
	FelineStakingContract string
)

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/TokenHubContract
	MainnetTokenHubContract string
	//go:embed mainnet/StakingContract
	MainnetStakingContract string
)
