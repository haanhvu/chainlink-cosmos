package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/haanhvu/chainlink-cosmos/cosmos/x/chainlink/types"
)

func (k msgServer) SubmitDataFromChainlinkFunctions(ctx context.Context, msg *types.MsgSubmitDataFromChainlinkFunctions) (*types.MsgSubmitDataFromChainlinkFunctionsResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgSubmitDataFromChainlinkFunctionsResponse{}, nil
}
