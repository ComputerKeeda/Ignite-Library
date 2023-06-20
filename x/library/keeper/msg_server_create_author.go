package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"library/x/library/types"
)

func (k msgServer) CreateAuthor(goCtx context.Context, msg *types.MsgCreateAuthor) (*types.MsgCreateAuthorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var author = types.Author{
		Name:     msg.Name,
		Bookname: msg.Bookname,
	}

	id := k.AppendAuthor(ctx, author)

	return &types.MsgCreateAuthorResponse{
		Id: id,
	}, nil
}
