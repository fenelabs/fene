package haber_fix

import _ "embed"

// contract codes for Feline upgrade
var (
	//go:embed feline/ValidatorContract
	FelineValidatorContract string
	//go:embed feline/SlashContract
	FelineSlashContract string
)

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/ValidatorContract
	MainnetValidatorContract string
	//go:embed mainnet/SlashContract
	MainnetSlashContract string
)
