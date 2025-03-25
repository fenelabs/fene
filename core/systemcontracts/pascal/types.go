package pascal

import _ "embed"

// contract codes for Mainnet upgrade
var (

	//go:embed mainnet/ValidatorContract
	MainnetValidatorContract string

	//go:embed mainnet/SlashContract
	MainnetSlashContract string

	//go:embed mainnet/SystemRewardContract
	MainnetSystemRewardContract string

	//go:embed mainnet/LightClientContract
	MainnetLightClientContract string

	//go:embed mainnet/TokenHubContract
	MainnetTokenHubContract string

	//go:embed mainnet/RelayerIncentivizeContract
	MainnetRelayerIncentivizeContract string

	//go:embed mainnet/RelayerHubContract
	MainnetRelayerHubContract string

	//go:embed mainnet/GovHubContract
	MainnetGovHubContract string

	//go:embed mainnet/TokenManagerContract
	MainnetTokenManagerContract string

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

// contract codes for Rialto upgrade
var (

	//go:embed rialto/ValidatorContract
	RialtoValidatorContract string

	//go:embed rialto/SlashContract
	RialtoSlashContract string

	//go:embed rialto/SystemRewardContract
	RialtoSystemRewardContract string

	//go:embed rialto/LightClientContract
	RialtoLightClientContract string

	//go:embed rialto/TokenHubContract
	RialtoTokenHubContract string

	//go:embed rialto/RelayerIncentivizeContract
	RialtoRelayerIncentivizeContract string

	//go:embed rialto/RelayerHubContract
	RialtoRelayerHubContract string

	//go:embed rialto/GovHubContract
	RialtoGovHubContract string

	//go:embed rialto/TokenManagerContract
	RialtoTokenManagerContract string

	//go:embed rialto/CrossChainContract
	RialtoCrossChainContract string

	//go:embed rialto/StakingContract
	RialtoStakingContract string

	//go:embed rialto/StakeHubContract
	RialtoStakeHubContract string

	//go:embed rialto/StakeCreditContract
	RialtoStakeCreditContract string

	//go:embed rialto/GovernorContract
	RialtoGovernorContract string

	//go:embed rialto/GovTokenContract
	RialtoGovTokenContract string

	//go:embed rialto/TimelockContract
	RialtoTimelockContract string

	//go:embed rialto/TokenRecoverPortalContract
	RialtoTokenRecoverPortalContract string
)
