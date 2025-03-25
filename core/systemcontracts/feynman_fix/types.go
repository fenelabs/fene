package feynman_fix

import _ "embed"

// contract codes for Feline upgrade
var (
	//go:embed feline/ValidatorContract
	FelineValidatorContract string
	//go:embed feline/StakeHubContract
	FelineStakeHubContract string
)
