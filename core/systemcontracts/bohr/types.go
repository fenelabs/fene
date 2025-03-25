package bohr

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/ValidatorContract
	MainnetValidatorContract string
	//go:embed mainnet/StakeHubContract
	MainnetStakeHubContract string
)

// contract codes for Feline upgrade
var (
	//go:embed feline/ValidatorContract
	FelineValidatorContract string
	//go:embed feline/StakeHubContract
	FelineStakeHubContract string
)
