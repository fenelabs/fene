package moran

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/RelayerHubContract
	MainnetRelayerHubContract string
	//go:embed mainnet/LightClientContract
	MainnetLightClientContract string
	//go:embed mainnet/CrossChainContract
	MainnetCrossChainContract string
)

// contract codes for Feline upgrade
var (
	//go:embed feline/RelayerHubContract
	FelineRelayerHubContract string
	//go:embed feline/LightClientContract
	FelineLightClientContract string
	//go:embed feline/CrossChainContract
	FelineCrossChainContract string
)
