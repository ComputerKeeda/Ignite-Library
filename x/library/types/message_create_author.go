package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAuthor = "create_author"

var _ sdk.Msg = &MsgCreateAuthor{}

func NewMsgCreateAuthor(creator string, name string, bookname string) *MsgCreateAuthor {
	return &MsgCreateAuthor{
		Creator:  creator,
		Name:     name,
		Bookname: bookname,
	}
}

func (msg *MsgCreateAuthor) Route() string {
	return RouterKey
}

func (msg *MsgCreateAuthor) Type() string {
	return TypeMsgCreateAuthor
}

func (msg *MsgCreateAuthor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAuthor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAuthor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
