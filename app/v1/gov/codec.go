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
}

var msgCdc = codec.New()
