package types

//go:generate mockgen --source=x/axelarcork/types/expected_keepers.go --destination=x/axelarcork/tests/mocks/expected_keepers_mocks.go --package=mocks

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
)

// AccountKeeper defines the expected account keeper.
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI

	SetModuleAccount(sdk.Context, authtypes.ModuleAccountI)
}

type ICS4Wrapper interface {
	WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI, acknowledgement ibcexported.Acknowledgement) error
	SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error
}

// ChannelKeeper defines the channel contract that must be fulfilled when
// creating a x/ratelimit keeper.
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, portID string, channelID string) (channeltypes.Channel, bool)
	GetChannelClientState(ctx sdk.Context, portID string, channelID string) (string, ibcexported.ClientState, error)
}

// StakingKeeper defines the expected staking keeper methods
type StakingKeeper interface {
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator
	GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) int64
	GetLastTotalPower(ctx sdk.Context) (power sdk.Int)
	IterateValidators(sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool))
	IterateBondedValidatorsByPower(sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool))
	IterateLastValidators(sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool))
	Validator(sdk.Context, sdk.ValAddress) stakingtypes.ValidatorI
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI
	Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec)
	Jail(sdk.Context, sdk.ConsAddress)
	PowerReduction(ctx sdk.Context) sdk.Int
}

type TransferKeeper interface {
	Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error)
}

type DistributionKeeper interface {
	GetFeePool(ctx sdk.Context) (feePool distributiontypes.FeePool)
	SetFeePool(ctx sdk.Context, feePool distributiontypes.FeePool)
}

// GravityKeeper defines the expected gravity keeper methods
type GravityKeeper interface {
	GetOrchestratorValidatorAddress(ctx sdk.Context, orchAddr sdk.AccAddress) sdk.ValAddress
}
