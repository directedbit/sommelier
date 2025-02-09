package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "cork"

	// StoreKey is the store key string for oracle
	StoreKey = ModuleName

	// RouterKey is the message route for oracle
	RouterKey = ModuleName

	// QuerierRoute is the querier route for oracle
	QuerierRoute = ModuleName
)

// Keys for cork store, with <prefix><key> -> <value>
const (
	_ = byte(iota)

	// CorkForAddressKeyPrefix - <prefix><val_address><address> -> <cork>
	CorkForAddressKeyPrefix // key for corks

	// CommitPeriodStartKey - <prefix> -> int64(height)
	CommitPeriodStartKey // key for commit period height start

	// LatestInvalidationNonceKey - <prefix> -> uint64(latestNonce)
	LatestInvalidationNonceKey

	// CellarIDsKey - <prefix> -> []string
	CellarIDsKey

	// ScheduledCorkKeyPrefix - <prefix><block_height><val_address><address> -> <cork>
	ScheduledCorkKeyPrefix

	// LatestCorkIDKey - <key> -> uint64(latestCorkID)
	LatestCorkIDKey

	// CorkResultPrefix - <prefix><id> -> CorkResult
	CorkResultPrefix
)

// GetCorkForValidatorAddressKey returns the key for a validators vote for a given address
func GetCorkForValidatorAddressKey(val sdk.ValAddress, contract common.Address) []byte {
	return append(GetCorkValidatorKeyPrefix(val), contract.Bytes()...)
}

// GetCorkValidatorKeyPrefix returns the key prefix for cork commits for a validator
func GetCorkValidatorKeyPrefix(val sdk.ValAddress) []byte {
	return append([]byte{CorkForAddressKeyPrefix}, val.Bytes()...)
}

func MakeCellarIDsKey() []byte {
	return []byte{CellarIDsKey}
}

func GetScheduledCorkKeyPrefix() []byte {
	return []byte{ScheduledCorkKeyPrefix}
}

func GetScheduledCorkKeyByBlockHeightPrefix(blockHeight uint64) []byte {
	return append(GetScheduledCorkKeyPrefix(), sdk.Uint64ToBigEndian(blockHeight)...)
}

func GetScheduledCorkKey(blockHeight uint64, id []byte, val sdk.ValAddress, contract common.Address) []byte {
	blockHeightBytes := sdk.Uint64ToBigEndian(blockHeight)
	return bytes.Join([][]byte{GetScheduledCorkKeyPrefix(), blockHeightBytes, id, val.Bytes(), contract.Bytes()}, []byte{})
}

func GetCorkResultPrefix() []byte {
	return []byte{CorkResultPrefix}
}

func GetCorkResultKey(id []byte) []byte {
	return append(GetCorkResultPrefix(), id...)
}
