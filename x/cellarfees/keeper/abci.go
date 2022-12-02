package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	appParams "github.com/peggyjv/sommelier/v4/app/params"
)

// BeginBlocker emits rewards each block they are available by sending them to the distribution module's fee collector
// account. Emissions are a constant value based on the last peak supply of distributable fees so that the reward supply
// will decrease linearly until exhausted.
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	// Handle reward emissions
	moduleAccount := k.GetFeesAccount(ctx)
	remainingRewardsSupply := k.bankKeeper.GetBalance(ctx, moduleAccount.GetAddress(), appParams.BaseCoinUnit).Amount

	if remainingRewardsSupply.IsZero() {
		return
	}

	previousSupplyPeak := k.GetLastRewardSupplyPeak(ctx)
	params := k.GetParams(ctx)

	var emissionAmount sdk.Int
	if remainingRewardsSupply.GT(previousSupplyPeak) {
		k.SetLastRewardSupplyPeak(ctx, remainingRewardsSupply)
		emissionAmount = remainingRewardsSupply.Quo(sdk.NewInt(int64(params.RewardEmissionPeriod)))
	} else {
		emissionAmount = previousSupplyPeak.Quo(sdk.NewInt(int64(params.RewardEmissionPeriod)))
	}

	// Emission should be at least 1usomm and at most the remaining reward supply
	if emissionAmount.IsZero() {
		emissionAmount = sdk.OneInt()
	} else if emissionAmount.GTE(remainingRewardsSupply) {
		// We zero out the previous peak value here to avoid doing it every block. We set the final emission
		// to the remaining supply here even though it's potentially redundant because it's less code than
		// having another check where we would also have to zero out the prevoius peak supply.
		k.SetLastRewardSupplyPeak(ctx, sdk.ZeroInt())
		emissionAmount = remainingRewardsSupply
	}

	coin := sdk.NewCoin(appParams.BaseCoinUnit, emissionAmount)
	emission := sdk.NewCoins(coin)

	// Send to fee collector for distribution
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, moduleAccount.GetName(), authtypes.FeeCollectorName, emission)
	if err != nil {
		panic(err)
	}
}

// EndBlocker is called at the end of every block
func (k Keeper) EndBlocker(ctx sdk.Context) {}
