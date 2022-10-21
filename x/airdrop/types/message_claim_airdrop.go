package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimAirdrop = "claim_airdrop"

var _ sdk.Msg = &MsgClaimAirdrop{}

func NewMsgClaimAirdrop(creator string, amount string) *MsgClaimAirdrop {
	return &MsgClaimAirdrop{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgClaimAirdrop) Route() string {
	return RouterKey
}

func (msg *MsgClaimAirdrop) Type() string {
	return TypeMsgClaimAirdrop
}

func (msg *MsgClaimAirdrop) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimAirdrop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimAirdrop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
