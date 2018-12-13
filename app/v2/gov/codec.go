package gov

import (
	"github.com/irisnet/irishub/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgSubmitProposal{}, "cosmos-sdk/MsgSubmitProposal/2", nil)
	cdc.RegisterConcrete(MsgSubmitTxTaxUsageProposal{}, "gov/MsgSubmitTxTaxUsageProposal/2", nil)
	cdc.RegisterConcrete(MsgSubmitSoftwareUpgradeProposal{}, "gov/MsgSubmitSoftwareUpgradeProposal/2", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "cosmos-sdk/MsgDeposit/2", nil)
	cdc.RegisterConcrete(MsgVote{}, "cosmos-sdk/MsgVote/2", nil)

	cdc.RegisterInterface((*Proposal)(nil), nil)
	cdc.RegisterConcrete(&TextProposal{}, "gov/TextProposal/2", nil)

	////////////////////  iris begin  ///////////////////////////
	cdc.RegisterConcrete(&ParameterProposal{}, "gov/ParameterProposal/2", nil)
	cdc.RegisterConcrete(&SoftwareUpgradeProposal{}, "gov/SoftwareUpgradeProposal/2", nil)
	cdc.RegisterConcrete(&HaltProposal{}, "gov/TerminatorProposal/2", nil)
	cdc.RegisterConcrete(&TaxUsageProposal{}, "gov/TaxUsageProposal/2", nil)
	////////////////////  iris end  ///////////////////////////
}

var msgCdc = codec.New()
