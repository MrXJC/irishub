package gov

import (
	"fmt"
	protocolKeeper "github.com/irisnet/irishub/app/protocol/keeper"
	sdk "github.com/irisnet/irishub/types"
	"github.com/irisnet/irishub/types/common"
)

var _ Proposal = (*SoftwareUpgradeProposal)(nil)

type SoftwareUpgradeProposal struct {
	TextProposal
	Version      uint64
	Software     string
	SwitchHeight uint64
}

func (sp *SoftwareUpgradeProposal) Execute(ctx sdk.Context, k Keeper) error {

	logger := ctx.Logger().With("module", "x/gov")

	if _, ok := k.pk.GetUpgradeConfig(ctx); ok {
		logger.Info("Execute SoftwareProposal Failure", "info",
			fmt.Sprintf("Software Upgrade Switch Period is in process. current height:%d", ctx.BlockHeight()))
		return nil
	}
	if k.pk.IsValidProtocolVersion(ctx, sp.Version) {
		logger.Info("Execute SoftwareProposal Failure", "info",
			fmt.Sprintf("version [%s] in SoftwareUpgradeProposal isn't valid ", sp.ProposalID))
		return nil
	}
	if uint64(ctx.BlockHeight())+1 >= sp.SwitchHeight {
		logger.Info("Execute SoftwareProposal Failure", "info",
			fmt.Sprintf("switch height must be more than blockHeight + 1"))
		return nil
	}

	k.pk.SetUpgradeConfig(ctx,
		protocolKeeper.UpgradeConfig{sp.ProposalID,
			common.ProtocolDefinition{sp.Version, sp.Software, sp.SwitchHeight}})

	logger.Info("Execute SoftwareProposal Success", "info",
		fmt.Sprintf("current height:%d", ctx.BlockHeight()))

	return nil
}
