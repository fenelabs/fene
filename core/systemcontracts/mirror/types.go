package mirror

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/TokenManagerContract
	MainnetTokenManagerContract string
	//go:embed mainnet/TokenHubContract
	MainnetTokenHubContract string
	//go:embed mainnet/RelayerIncentivizeContract
	MainnetRelayerIncentivizeContract string
)

// contract codes for Chapel upgrade
var (
	//go:embed feline/TokenManagerContract
	FelineTokenManagerContract string
	//go:embed feline/TokenHubContract
	FelineTokenHubContract string
	//go:embed feline/RelayerIncentivizeContract
	FelineRelayerIncentivizeContract string
)
