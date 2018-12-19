package gov

import (
	"github.com/irisnet/irishub/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgSubmitProposal{}, "cosmos-sdk/MsgSubmitProposal3", nil)
	cdc.RegisterConcrete(MsgSubmitTxTaxUsageProposal{}, "gov/MsgSubmitTxTaxUsageProposal3", nil)
	cdc.RegisterConcrete(MsgSubmitSoftwareUpgradeProposal{}, "gov/MsgSubmitSoftwareUpgradeProposal3", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "cosmos-sdk/MsgDeposit3", nil)
	cdc.RegisterConcrete(MsgVote{}, "cosmos-sdk/MsgVote3", nil)

	cdc.RegisterInterface((*Proposal)(nil), nil)
	cdc.RegisterConcrete(&TextProposal{}, "gov/TextProposal3", nil)

	////////////////////  iris begin  ///////////////////////////
	cdc.RegisterConcrete(&ParameterProposal{}, "gov/ParameterProposal3", nil)
	cdc.RegisterConcrete(&SoftwareUpgradeProposal{}, "gov/SoftwareUpgradeProposal3", nil)
	cdc.RegisterConcrete(&HaltProposal{}, "gov/TerminatorProposal3", nil)
	cdc.RegisterConcrete(&TaxUsageProposal{}, "gov/TaxUsageProposal3", nil)
	////////////////////  iris end  ///////////////////////////
}

var msgCdc = codec.New()
