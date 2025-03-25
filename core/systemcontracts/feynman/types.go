package feynman

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/ValidatorContract
	MainnetValidatorContract string
	//go:embed mainnet/SlashContract
	MainnetSlashContract string
	//go:embed mainnet/TokenHubContract
	MainnetTokenHubContract string
	//go:embed mainnet/GovHubContract
	MainnetGovHubContract string
	//go:embed mainnet/CrossChainContract
	MainnetCrossChainContract string
	//go:embed mainnet/StakingContract
	MainnetStakingContract string
	//go:embed mainnet/StakeHubContract
	MainnetStakeHubContract string
	//go:embed mainnet/StakeCreditContract
	MainnetStakeCreditContract string
	//go:embed mainnet/GovernorContract
	MainnetGovernorContract string
	//go:embed mainnet/GovTokenContract
	MainnetGovTokenContract string
	//go:embed mainnet/TimelockContract
	MainnetTimelockContract string
	//go:embed mainnet/TokenRecoverPortalContract
	MainnetTokenRecoverPortalContract string
)

// contract codes for Feline upgrade
var (
	//go:embed feline/ValidatorContract
	FelineValidatorContract string
	//go:embed feline/SlashContract
	FelineSlashContract string
	//go:embed feline/TokenHubContract
	FelineTokenHubContract string
	//go:embed feline/GovHubContract
	FelineGovHubContract string
	//go:embed feline/CrossChainContract
	FelineCrossChainContract string
	//go:embed feline/StakingContract
	FelineStakingContract string
	//go:embed feline/StakeHubContract
	FelineStakeHubContract string
	//go:embed feline/StakeCreditContract
	FelineStakeCreditContract string
	//go:embed feline/GovernorContract
	FelineGovernorContract string
	//go:embed feline/GovTokenContract
	FelineGovTokenContract string
	//go:embed feline/TimelockContract
	FelineTimelockContract string
	//go:embed feline/TokenRecoverPortalContract
	FelineTokenRecoverPortalContract string
)
