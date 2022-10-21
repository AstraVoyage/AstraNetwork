package simulation

import (
	"math/rand"

	"AstraNetwork/x/airdrop/keeper"
	"AstraNetwork/x/airdrop/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgClaimAirdrop(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgClaimAirdrop{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ClaimAirdrop simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ClaimAirdrop simulation not implemented"), nil, nil
	}
}
