package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	auctionTypes "github.com/peggyjv/sommelier/v4/x/auction/types"
)

type runsBeforeWrapper func()

// Happy path test for submitting a bid both fully and partially
func (suite *KeeperTestSuite) TestHappyPathSubmitBidAndFulfillFully() {
	ctx, auctionKeeper := suite.ctx, suite.auctionKeeper
	require := suite.Require()

	// -----> Create an auction we can bid on first
	params := auctionTypes.Params{PriceMaxBlockAge: 10}
	auctionKeeper.setParams(ctx, params)

	sommPrice := auctionTypes.TokenPrice{Denom: auctionTypes.UsommDenom, UsdPrice: sdk.MustNewDecFromStr("0.01"), LastUpdatedBlock: 5}

	/* #nosec */
	saleToken := "gravity0xdac17f958d2ee523a2206206994597c13d831ec7"
	saleTokenPrice := auctionTypes.TokenPrice{Denom: saleToken, UsdPrice: sdk.MustNewDecFromStr("0.02"), LastUpdatedBlock: 5}
	auctionedSaleTokens := sdk.NewCoin(saleToken, sdk.NewInt(10000))

	auctionKeeper.setTokenPrice(ctx, sommPrice)
	auctionKeeper.setTokenPrice(ctx, saleTokenPrice)

	// Mock bank keeper fund transfer
	suite.mockSendCoinsFromModuleToModule(ctx, permissionedFunder.GetName(), auctionTypes.ModuleName, sdk.NewCoins(auctionedSaleTokens))

	// Start auction
	decreaseRate := sdk.MustNewDecFromStr("0.05")
	blockDecreaseInterval := uint64(5)
	err := auctionKeeper.BeginAuction(ctx, auctionedSaleTokens, decreaseRate, blockDecreaseInterval, permissionedFunder.GetName(), permissionedReciever.GetName())
	require.Nil(err)

	// Submit a bid
	auctionID := uint32(1)
	bidder := cosmos_address_1
	require.Nil(err)

	bid := sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(5000))
	minAmount := sdk.NewCoin(saleToken, sdk.NewInt(1))

	fulfilledBid := sdk.NewCoin(saleToken, sdk.NewInt(2500))

	// Mock out bank keeper calls
	bidderAcc, _ := sdk.AccAddressFromBech32(bidder)
	suite.mockGetModuleAccount(ctx)
	suite.mockGetBalance(ctx, authtypes.NewEmptyModuleAccount("mock").GetAddress(), saleToken, auctionedSaleTokens)
	suite.mockSendCoinsFromAccountToModule(ctx, bidderAcc, auctionTypes.ModuleName, sdk.NewCoins(bid))
	suite.mockSendCoinsFromModuleToAccount(ctx, auctionTypes.ModuleName, bidderAcc, sdk.NewCoins(fulfilledBid))

	// ~Actually~ submit the bid now
	response, err := auctionKeeper.SubmitBid(sdk.WrapSDKContext(ctx), &auctionTypes.MsgSubmitBidRequest{
		AuctionId:              auctionID,
		Bidder:                 bidder,
		MaxBidInUsomm:          bid,
		SaleTokenMinimumAmount: minAmount,
	})
	require.Nil(err)

	// Assert bid store and response bids are both equal to expected bid
	expectedBid := auctionTypes.Bid{
		Id:                        uint64(1),
		AuctionId:                 uint32(1),
		Bidder:                    bidder,
		MaxBidInUsomm:             bid,
		SaleTokenMinimumAmount:    minAmount,
		TotalFulfilledSaleTokens:  fulfilledBid,
		SaleTokenUnitPriceInUsomm: sdk.NewDec(2),
		TotalUsommPaid:            bid,
	}
	require.Equal(&expectedBid, response.Bid)

	storedBid, found := auctionKeeper.GetBid(ctx, uint32(1), uint64(1))
	require.True(found)
	require.Equal(expectedBid, storedBid)

	// Assert auction token amounts are updated
	expectedUpdatedAuction := auctionTypes.Auction{
		Id:                         auctionID,
		StartingTokensForSale:      auctionedSaleTokens,
		StartBlock:                 uint64(ctx.BlockHeight()),
		EndBlock:                   0,
		InitialPriceDecreaseRate:   decreaseRate,
		CurrentPriceDecreaseRate:   decreaseRate,
		PriceDecreaseBlockInterval: blockDecreaseInterval,
		InitialUnitPriceInUsomm:    sdk.NewDec(2),
		CurrentUnitPriceInUsomm:    sdk.NewDec(2),
		RemainingTokensForSale:     sdk.NewCoin(saleToken, sdk.NewInt(7500)), // this is the important part, need to make sure it decremented
		FundingModuleAccount:       permissionedFunder.GetName(),
		ProceedsModuleAccount:      permissionedReciever.GetName(),
	}

	activeAuction, found := auctionKeeper.GetActiveAuctionByID(ctx, auctionID)
	require.True(found)
	require.Equal(expectedUpdatedAuction, activeAuction)

	// Now check flow of a bid that can only be partially fulfilled, and verify it finishes the auction
	newBidder := cosmos_address_2
	newBid := sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(50000))
	newFulfilledAmt := sdk.NewCoin(saleToken, sdk.NewInt(7500))
	paidAmt := sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(15000))

	// Mock out necessary bank keeper calls for bid completion
	newBidderAcc, _ := sdk.AccAddressFromBech32(newBidder)
	suite.mockGetModuleAccount(ctx)
	suite.mockGetBalance(ctx, authtypes.NewEmptyModuleAccount("mock").GetAddress(), saleToken, expectedUpdatedAuction.RemainingTokensForSale)
	suite.mockSendCoinsFromAccountToModule(ctx, newBidderAcc, auctionTypes.ModuleName, sdk.NewCoins(paidAmt))
	suite.mockSendCoinsFromModuleToAccount(ctx, auctionTypes.ModuleName, newBidderAcc, sdk.NewCoins(newFulfilledAmt))

	// Mock out final keeper calls necessary to finish the auction due to bid draining the availible supply
	suite.mockGetBalance(ctx, authtypes.NewModuleAddress(auctionTypes.ModuleName), saleToken, sdk.NewCoin(saleToken, sdk.NewInt(0)))
	totalUsommExpected := sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(20000))
	suite.mockSendCoinsFromModuleToModule(ctx, auctionTypes.ModuleName, permissionedReciever.GetName(), sdk.NewCoins(totalUsommExpected))

	// Submit the partially fulfillable bid now
	response, err = auctionKeeper.SubmitBid(sdk.WrapSDKContext(ctx), &auctionTypes.MsgSubmitBidRequest{
		AuctionId:              auctionID,
		Bidder:                 newBidder,
		MaxBidInUsomm:          newBid,
		SaleTokenMinimumAmount: minAmount,
	})
	require.Nil(err)

	// Assert bid store and response bids are both equal to the new expected bid
	newExpectedBid := auctionTypes.Bid{
		Id:                        uint64(2),
		AuctionId:                 uint32(1),
		Bidder:                    newBidder,
		MaxBidInUsomm:             newBid,
		SaleTokenMinimumAmount:    minAmount,
		TotalFulfilledSaleTokens:  newFulfilledAmt,
		SaleTokenUnitPriceInUsomm: sdk.NewDec(2),
		TotalUsommPaid:            paidAmt,
	}
	require.Equal(&newExpectedBid, response.Bid)

	storedBid, found = auctionKeeper.GetBid(ctx, uint32(1), uint64(2))
	require.True(found)
	require.Equal(newExpectedBid, storedBid)

	// Verify bid caused auction to finish
	expectedUpdatedAuction.RemainingTokensForSale.Amount = sdk.NewInt(0)
	expectedUpdatedAuction.EndBlock = uint64(ctx.BlockHeight())

	_, found = auctionKeeper.GetActiveAuctionByID(ctx, auctionID)
	require.False(found)

	endedAuction, found := auctionKeeper.GetEndedAuctionByID(ctx, auctionID)
	require.True(found)
	require.Equal(expectedUpdatedAuction, endedAuction)
}

// Unhappy path tests for all failure modes of SubmitBid
func (suite *KeeperTestSuite) TestUnhappyPathsForSubmitBid() {
	ctx, auctionKeeper := suite.ctx, suite.auctionKeeper
	require := suite.Require()

	// Create an active auction for bids to test against
	params := auctionTypes.Params{PriceMaxBlockAge: 10}
	auctionKeeper.setParams(ctx, params)

	sommPrice := auctionTypes.TokenPrice{Denom: auctionTypes.UsommDenom, UsdPrice: sdk.MustNewDecFromStr("0.01"), LastUpdatedBlock: 5}

	/* #nosec */
	saleToken := "gravity0x853d955acef822db058eb8505911ed77f175b99e"
	saleTokenPrice := auctionTypes.TokenPrice{Denom: saleToken, UsdPrice: sdk.MustNewDecFromStr("0.02"), LastUpdatedBlock: 5}
	auctionedSaleTokens := sdk.NewCoin(saleToken, sdk.NewInt(10000))

	auctionKeeper.setTokenPrice(ctx, sommPrice)
	auctionKeeper.setTokenPrice(ctx, saleTokenPrice)

	// Mock bank keeper fund transfer
	suite.mockSendCoinsFromModuleToModule(ctx, permissionedFunder.GetName(), auctionTypes.ModuleName, sdk.NewCoins(auctionedSaleTokens))

	// Start auction
	decreaseRate := sdk.MustNewDecFromStr("0.05")
	blockDecreaseInterval := uint64(5)
	err := auctionKeeper.BeginAuction(ctx, auctionedSaleTokens, decreaseRate, blockDecreaseInterval, permissionedFunder.GetName(), permissionedReciever.GetName())
	require.Nil(err)

	// Verify auction got added to active auction store
	auctionID := uint32(1)
	originalAuction, found := auctionKeeper.GetActiveAuctionByID(ctx, auctionID)
	require.True(found)

	tests := []struct {
		name              string
		bid               auctionTypes.MsgSubmitBidRequest
		expectedError     error
		runsBefore  runsBeforeWrapper
		submitBidResponse *auctionTypes.MsgSubmitBidResponse
	}{
		{
			name: "Auction ID not found",
			bid: auctionTypes.MsgSubmitBidRequest{
				AuctionId:              uint32(420),
				Bidder:                 cosmos_address_1,
				MaxBidInUsomm:          sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(100)),
				SaleTokenMinimumAmount: sdk.NewCoin(saleToken, sdk.NewInt(1)),
			},
			expectedError:     sdkerrors.Wrapf(auctionTypes.ErrAuctionNotFound, "Auction id: %d", uint32(420)),
			runsBefore:  func() {},
			submitBidResponse: &auctionTypes.MsgSubmitBidResponse{},
		},
		{
			name: "Denom mismatch",
			bid: auctionTypes.MsgSubmitBidRequest{
				AuctionId:              auctionID,
				Bidder:                 cosmos_address_1,
				MaxBidInUsomm:          sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(100)),
				SaleTokenMinimumAmount: sdk.NewCoin("blemflarcks", sdk.NewInt(1)),
			},
			expectedError:     sdkerrors.Wrapf(auctionTypes.ErrBidAuctionDenomMismatch, "Bid denom: blemflarcks, Auction denom: %s", saleToken),
			runsBefore:  func() {},
			submitBidResponse: &auctionTypes.MsgSubmitBidResponse{},
		},
		{
			name: "Minimum amount to purchase larger than bid can obtain",
			bid: auctionTypes.MsgSubmitBidRequest{
				AuctionId:              auctionID,
				Bidder:                 cosmos_address_1,
				MaxBidInUsomm:          sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(1)),
				SaleTokenMinimumAmount: sdk.NewCoin(saleToken, sdk.NewInt(1)),
			},
			expectedError: sdkerrors.Wrapf(auctionTypes.ErrInsufficientBid, "minimum purchase price: 2, max bid: 1"),
			runsBefore: func() {
				suite.mockGetModuleAccount(ctx)
				suite.mockGetBalance(ctx, authtypes.NewEmptyModuleAccount("mock").GetAddress(), saleToken, originalAuction.RemainingTokensForSale)
			},
			submitBidResponse: &auctionTypes.MsgSubmitBidResponse{},
		},
		{
			name: "Minimum amount larger than remaining tokens in auction",
			bid: auctionTypes.MsgSubmitBidRequest{
				AuctionId:              auctionID,
				Bidder:                 cosmos_address_1,
				MaxBidInUsomm:          sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(40000)),
				SaleTokenMinimumAmount: sdk.NewCoin(saleToken, sdk.NewInt(10002)),
			},
			expectedError: sdkerrors.Wrapf(auctionTypes.ErrMinimumPurchaseAmountLargerThanTokensRemaining, "Minimum purchase: %s, amount remaining: %s", sdk.NewInt(10002), originalAuction.RemainingTokensForSale.String()),
			runsBefore: func() {
				suite.mockGetModuleAccount(ctx)
				suite.mockGetBalance(ctx, authtypes.NewEmptyModuleAccount("mock").GetAddress(), saleToken, originalAuction.RemainingTokensForSale)
			},
			submitBidResponse: &auctionTypes.MsgSubmitBidResponse{},
		},
		{
			name: "Validate Basic canary 1 -- bid denom must be in usomm",
			bid: auctionTypes.MsgSubmitBidRequest{
				AuctionId:              auctionID,
				Bidder:                 cosmos_address_1,
				MaxBidInUsomm:          sdk.NewCoin("cinnamonRollCoin", sdk.NewInt(200)),
				SaleTokenMinimumAmount: sdk.NewCoin(saleToken, sdk.NewInt(100)),
			},
			expectedError: sdkerrors.Wrapf(auctionTypes.ErrBidMustBeInUsomm, "bid: %s", sdk.NewCoin("cinnamonRollCoin", sdk.NewInt(200)).String()),
			runsBefore: func() {
				suite.mockGetModuleAccount(ctx)
				suite.mockGetBalance(ctx, authtypes.NewEmptyModuleAccount("mock").GetAddress(), saleToken, originalAuction.RemainingTokensForSale)
			},
			submitBidResponse: &auctionTypes.MsgSubmitBidResponse{},
		},
		{
			name: "Validate Basic canary 2 -- minimum amount of sale tokens cannot be 0",
			bid: auctionTypes.MsgSubmitBidRequest{
				AuctionId:              auctionID,
				Bidder:                 cosmos_address_1,
				MaxBidInUsomm:          sdk.NewCoin(auctionTypes.UsommDenom, sdk.NewInt(200)),
				SaleTokenMinimumAmount: sdk.NewCoin(saleToken, sdk.NewInt(0)),
			},
			expectedError: sdkerrors.Wrapf(auctionTypes.ErrMinimumAmountMustBePositive, "sale token amount: %s", sdk.NewCoin(saleToken, sdk.NewInt(0)).String()),
			runsBefore: func() {
				suite.mockGetModuleAccount(ctx)
				suite.mockGetBalance(ctx, authtypes.NewEmptyModuleAccount("mock").GetAddress(), saleToken, originalAuction.RemainingTokensForSale)
			},
			submitBidResponse: &auctionTypes.MsgSubmitBidResponse{},
		},
	}

	for _, tc := range tests {
		tc := tc // Redefine variable here due to passing it to function literal below (scopelint)
		suite.T().Run(fmt.Sprint(tc.name), func(t *testing.T) {
			// Run expected bank keeper functions, if any
			tc.runsBefore()
			response, err := auctionKeeper.SubmitBid(sdk.WrapSDKContext(ctx), &tc.bid)

			// Verify bid errors are as expected
			require.Equal(tc.expectedError.Error(), err.Error())
			require.Equal(tc.submitBidResponse, response)

			// Verify original auction not changed
			foundAuction, found := auctionKeeper.GetActiveAuctionByID(ctx, auctionID)
			require.True(found)
			require.Equal(originalAuction, foundAuction)
		})
	}
}
