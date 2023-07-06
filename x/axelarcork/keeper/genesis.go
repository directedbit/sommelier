package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/peggyjv/sommelier/v6/x/axelarcork/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, gs types.GenesisState) {
	k.SetParams(ctx, *gs.Params)
	// Set the vote period at initialization

	for i, config := range gs.ChainConfigurations.Configurations {
		k.SetChainConfigurationByID(ctx, config.Id, *config)
		k.SetCellarIDs(ctx, config.Id, *gs.CellarIds[i])
		k.SetLatestInvalidationNonce(ctx, config.Id, gs.InvalidationNonces[i])

		for _, corkResult := range gs.CorkResults[i].CorkResults {
			k.SetCorkResult(ctx, config.Id, corkResult.Cork.IDHash(corkResult.BlockHeight), *corkResult)
		}

		for _, scheduledCork := range gs.ScheduledCorks[i].ScheduledCorks {
			valAddr, err := sdk.ValAddressFromHex(scheduledCork.Validator)
			if err != nil {
				panic(err)
			}

			k.SetScheduledCork(ctx, config.Id, scheduledCork.BlockHeight, valAddr, *scheduledCork.Cork)
		}
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	var gs types.GenesisState

	ps := k.GetParamSet(ctx)
	gs.Params = &ps

	k.IterateChainConfigurations(ctx, func(config types.ChainConfiguration) (stop bool) {
		gs.ChainConfigurations.Configurations = append(gs.ChainConfigurations.Configurations, &config)

		cellarIDs := k.GetCellarIDs(ctx, config.Id)
		var cellarIDSet types.CellarIDSet
		for _, id := range cellarIDs {
			cellarIDSet.Ids = append(cellarIDSet.Ids, id.String())
		}
		gs.CellarIds = append(gs.CellarIds, &cellarIDSet)

		gs.InvalidationNonces = append(gs.InvalidationNonces, k.GetLatestInvalidationNonce(ctx, config.Id))

		var scheduledCorks types.ScheduledCorks
		scheduledCorks.ScheduledCorks = append(scheduledCorks.ScheduledCorks, k.GetScheduledCorks(ctx, config.Id)...)
		gs.ScheduledCorks = append(gs.ScheduledCorks, &scheduledCorks)

		var corkResults types.CorkResults
		corkResults.CorkResults = append(corkResults.CorkResults, k.GetCorkResults(ctx, config.Id)...)
		gs.CorkResults = append(gs.CorkResults, &corkResults)

		return false
	})

	return gs
}
