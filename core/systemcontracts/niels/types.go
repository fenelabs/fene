package niels

import _ "embed"

// contract codes for Feline upgrade
var (
	//go:embed feline/ValidatorContract
	FelineValidatorContract string
	//go:embed feline/SlashContract
	FelineSlashContract string
	//go:embed feline/SystemRewardContract
	FelineSystemRewardContract string
	//go:embed feline/LightClientContract
	FelineLightClientContract string
	//go:embed feline/TokenHubContract
	FelineTokenHubContract string
	//go:embed feline/RelayerIncentivizeContract
	FelineRelayerIncentivizeContract string
	//go:embed feline/RelayerHubContract
	FelineRelayerHubContract string
	//go:embed feline/GovHubContract
	FelineGovHubContract string
	//go:embed feline/TokenManagerContract
	FelineTokenManagerContract string
	//go:embed feline/CrossChainContract
	FelineCrossChainContract string
)
