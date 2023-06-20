package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"library/x/library/types"
)

func (k msgServer) CreateAuthor(goCtx context.Context, msg *types.MsgCreateAuthor) (*types.MsgCreateAuthorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateAuthorResponse{}, nil
}
