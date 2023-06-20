package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"library/x/library/types"
)

func (k Keeper) ListAuthor(goCtx context.Context, req *types.QueryListAuthorRequest) (*types.QueryListAuthorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var authors []types.Author
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.AuthorKey))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var author types.Author
		if err := k.cdc.Unmarshal(value, &author); err != nil {
			return err
		}

		authors = append(authors, author)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &types.QueryListAuthorResponse{
		Author:  authors,
		Pagination: pageRes,
	}, nil
}
