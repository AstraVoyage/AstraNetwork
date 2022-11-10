package keeper

import (
	"context"

	"AstraNetwork/x/airdrop/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimAirdrop(goCtx context.Context, msg *types.MsgClaimAirdrop) (*types.MsgClaimAirdropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Genesis address sending tokens to airdrop
	airdropSender, _ := sdk.AccAddressFromBech32("astra1srhapxe5ze83gesmcfv9hnzlt9s0wz7zweuly7")

	// Address from Keplr sign on user
	airdropRecipient, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Airdrop value to be sent
	airdropValue := sdk.Coins{sdk.NewInt64Coin("astra", 7777)}

	// Send Airdrop value to Keplr sign on user
	err := k.bankKeeper.SendCoins(ctx, airdropSender, airdropRecipient, airdropValue)
	if err != nil {
		return nil, err
	}

	// Create variable of type Post
	var claimed = types.Claimed{
		Creator: msg.Creator,
		Amount:  msg.Amount,
	}

	// Add a claimed to the store and get back the ID
	id := k.AppendClaimed(ctx, claimed)

	return &types.MsgClaimAirdropResponse{Id: id}, nil
}
