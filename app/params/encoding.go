package params

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakeTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	gravitytypes "github.com/peggyjv/gravity-bridge/module/v3/x/gravity/types"
	corkTypes "github.com/peggyjv/sommelier/v6/x/cork/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeCodec() EncodingConfig {
	modBasic := module.NewBasicManager()
	encodingConfig := MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	modBasic.RegisterLegacyAminoCodec(encodingConfig.Amino)
	modBasic.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	//now all the extra types

	distTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	bankTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	stakeTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	transfertypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	authz.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	govTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	upgradeTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	corkTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	gravitytypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	vestingTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() EncodingConfig {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          authTx.NewTxConfig(marshaler, authTx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}
}
