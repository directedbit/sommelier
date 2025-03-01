package cli

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestQueryScheduledCorksByBlockHeightCmd(t *testing.T) {
	testCases := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "Block height overflow",
			args: []string{
				"18446744073709551616",
			},
			err: sdkerrors.New("", uint32(1), "strconv.Atoi: parsing \"18446744073709551616\": value out of range"),
		},
	}

	for _, tc := range testCases {
		cmd := *queryScheduledCorksByBlockHeight()
		cmd.SetArgs(tc.args)
		err := cmd.Execute()

		require.Equal(t, tc.err.Error(), err.Error())
	}
}

func TestQueryScheduledCorksByIDCmd(t *testing.T) {
	testCases := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "Invalid ID",
			args: []string{
				"bad",
			},
			err: sdkerrors.New("", uint32(1), "invalid ID length, must be a keccak256 hash"),
		},
	}

	for _, tc := range testCases {
		cmd := *queryScheduledCorksByID()
		cmd.SetArgs(tc.args)
		err := cmd.Execute()

		require.Equal(t, tc.err.Error(), err.Error())
	}
}

func TestQueryCorkResultCmd(t *testing.T) {
	testCases := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "Invalid ID",
			args: []string{
				"bad",
			},
			err: sdkerrors.New("", uint32(1), "invalid ID length, must be a keccak256 hash"),
		},
	}

	for _, tc := range testCases {
		cmd := *queryCorkResult()
		cmd.SetArgs(tc.args)
		err := cmd.Execute()

		require.Equal(t, tc.err.Error(), err.Error())
	}
}
