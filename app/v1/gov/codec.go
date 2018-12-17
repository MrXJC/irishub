package gov

import (
	"github.com/irisnet/irishub/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgSubmitProposal{}, "cosmos-sdk/MsgSubmitProposal1", nil)
	cdc.RegisterConcrete(MsgSubmitTxTaxUsageProposal{}, "gov/MsgSubmitTxTaxUsageProposal1", nil)
	cdc.RegisterConcrete(MsgSubmitSoftwareUpgradeProposal{}, "gov/MsgSubmitSoftwareUpgradeProposal1", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "cosmos-sdk/MsgDeposit1", nil)
	cdc.RegisterConcrete(MsgVote{}, "cosmos-sdk/MsgVote1", nil)

	cdc.RegisterInterface((*Proposal)(nil), nil)
	cdc.RegisterConcrete(&TextProposal{}, "gov/TextProposal1", nil)

	////////////////////  iris begin  ///////////////////////////
	cdc.RegisterConcrete(&ParameterProposal{}, "gov/ParameterProposal1", nil)
	cdc.RegisterConcrete(&SoftwareUpgradeProposal{}, "gov/SoftwareUpgradeProposal1", nil)
	cdc.RegisterConcrete(&HaltProposal{}, "gov/TerminatorProposal1", nil)
	cdc.RegisterConcrete(&TaxUsageProposal{}, "gov/TaxUsageProposal1", nil)
	////////////////////  iris end  ///////////////////////////
}

var msgCdc = codec.New()