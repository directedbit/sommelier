package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProposalTypeAddManagedCellarIDs      = "AddManagedCellarIDs"
	ProposalTypeRemoveManagedCellarIDs   = "RemoveManagedCellarIDs"
	ProposalTypeScheduledCork            = "ScheduledCork"
	ProposalTypeCommunitySpend           = "CommunitySpend"
	ProposalTypeAddChainConfiguration    = "AddChainConfiguration"
	ProposalTypeRemoveChainConfiguration = "RemoveChainConfiguration"
)

var _ govtypes.Content = &AddManagedCellarIDsProposal{}
var _ govtypes.Content = &RemoveManagedCellarIDsProposal{}
var _ govtypes.Content = &ScheduledCorkProposal{}
var _ govtypes.Content = &CommunityPoolSpendProposal{}
var _ govtypes.Content = &AddChainConfigurationProposal{}
var _ govtypes.Content = &RemoveChainConfigurationProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddManagedCellarIDs)
	govtypes.RegisterProposalTypeCodec(&AddManagedCellarIDsProposal{}, "sommelier/AddManagedCellarIDsProposal")

	govtypes.RegisterProposalType(ProposalTypeRemoveManagedCellarIDs)
	govtypes.RegisterProposalTypeCodec(&RemoveManagedCellarIDsProposal{}, "sommelier/RemoveManagedCellarIDsProposal")

	govtypes.RegisterProposalType(ProposalTypeScheduledCork)
	govtypes.RegisterProposalTypeCodec(&ScheduledCorkProposal{}, "sommelier/ScheduledCorkProposal")

	govtypes.RegisterProposalType(ProposalTypeAddChainConfiguration)
	govtypes.RegisterProposalTypeCodec(&AddChainConfigurationProposal{}, "sommelier/AddChainConfigurationProposal")

	govtypes.RegisterProposalType(ProposalTypeRemoveChainConfiguration)
	govtypes.RegisterProposalTypeCodec(&RemoveChainConfigurationProposal{}, "sommelier/RemoveChainConfigurationProposal")

}

func NewAddManagedCellarIDsProposal(title string, description string, cellarIds *CellarIDSet) *AddManagedCellarIDsProposal {
	return &AddManagedCellarIDsProposal{
		Title:       title,
		Description: description,
		CellarIds:   cellarIds,
	}
}

func (m *AddManagedCellarIDsProposal) ProposalRoute() string {
	return RouterKey
}

func (m *AddManagedCellarIDsProposal) ProposalType() string {
	return ProposalTypeAddManagedCellarIDs
}

func (m *AddManagedCellarIDsProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}

	if len(m.CellarIds.Ids) == 0 {
		return fmt.Errorf("can't have an add prosoposal with no cellars")
	}

	return nil
}

func NewRemoveManagedCellarIDsProposal(title string, description string, cellarIds *CellarIDSet) *RemoveManagedCellarIDsProposal {
	return &RemoveManagedCellarIDsProposal{
		Title:       title,
		Description: description,
		CellarIds:   cellarIds,
	}
}

func (m *RemoveManagedCellarIDsProposal) ProposalRoute() string {
	return RouterKey
}

func (m *RemoveManagedCellarIDsProposal) ProposalType() string {
	return ProposalTypeRemoveManagedCellarIDs
}

func (m *RemoveManagedCellarIDsProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}

	if len(m.CellarIds.Ids) == 0 {
		return fmt.Errorf("can't have a remove prosoposal with no cellars")
	}

	return nil
}

func NewScheduledCorkProposal(title string, description string, blockHeight uint64, chainName string, chainID uint64, targetContractAddress string, contractCallProtoJSON string) *ScheduledCorkProposal {
	return &ScheduledCorkProposal{
		Title:                 title,
		Description:           description,
		BlockHeight:           blockHeight,
		ChainName:             chainName,
		ChainId:               chainID,
		TargetContractAddress: targetContractAddress,
		ContractCallProtoJson: contractCallProtoJSON,
	}
}

func (m *ScheduledCorkProposal) ProposalRoute() string {
	return RouterKey
}

func (m *ScheduledCorkProposal) ProposalType() string {
	return ProposalTypeScheduledCork
}

func (m *ScheduledCorkProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}

	if len(m.ContractCallProtoJson) == 0 {
		return sdkerrors.Wrapf(ErrInvalidJSON, "cannot have empty contract call")
	}

	if !json.Valid([]byte(m.ContractCallProtoJson)) {
		return sdkerrors.Wrapf(ErrInvalidJSON, "%s", m.ContractCallProtoJson)
	}

	if !common.IsHexAddress(m.TargetContractAddress) {
		return sdkerrors.Wrapf(ErrInvalidEVMAddress, "%s", m.TargetContractAddress)
	}

	return nil
}

func NewCommunitySpendProposal(title string, description string, recipient string, chainID uint64, chainName string, amount sdk.Coin) *CommunityPoolSpendProposal {
	return &CommunityPoolSpendProposal{
		Title:       title,
		Description: description,
		Recipient:   recipient,
		ChainId:     chainID,
		ChainName:   chainName,
		Amount:      amount,
	}
}

func (m *CommunityPoolSpendProposal) ProposalRoute() string {
	return RouterKey
}

func (m *CommunityPoolSpendProposal) ProposalType() string {
	return ProposalTypeCommunitySpend
}

func (m *CommunityPoolSpendProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}

	if m.Amount.Amount.IsZero() {
		return ErrValuelessSend
	}

	if m.Recipient == "" {
		return sdkerrors.Wrapf(ErrInvalidEVMAddress, "empty recipient")
	}

	if !common.IsHexAddress(m.Recipient) {
		return sdkerrors.Wrapf(ErrInvalidEVMAddress, "%s", m.Recipient)
	}

	return nil
}

func NewAddChainConfigurationProposal(title string, description string, configuration ChainConfiguration) *AddChainConfigurationProposal {
	return &AddChainConfigurationProposal{
		Title:              title,
		Description:        description,
		ChainConfiguration: &configuration,
	}
}

func (m *AddChainConfigurationProposal) ProposalRoute() string {
	return RouterKey
}

func (m *AddChainConfigurationProposal) ProposalType() string {
	return ProposalTypeAddChainConfiguration
}

func (m *AddChainConfigurationProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}

	if err := m.ChainConfiguration.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

func NewRemoveChainConfigurationProposal(title string, description string, chainID uint64) *RemoveChainConfigurationProposal {
	return &RemoveChainConfigurationProposal{
		Title:       title,
		Description: description,
		ChainId:     chainID,
	}
}

func (m *RemoveChainConfigurationProposal) ProposalRoute() string {
	return RouterKey
}

func (m *RemoveChainConfigurationProposal) ProposalType() string {
	return ProposalTypeRemoveChainConfiguration
}

func (m *RemoveChainConfigurationProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}
