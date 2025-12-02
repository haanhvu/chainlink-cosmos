package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/haanhvu/chainlink-cosmos/cosmos/x/chainlink/keeper"
	"github.com/haanhvu/chainlink-cosmos/cosmos/x/chainlink/types"
)

func SimulateMsgSubmitDataFromChainlinkFunctions(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmitDataFromChainlinkFunctions{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handle the SubmitDataFromChainlinkFunctions simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "SubmitDataFromChainlinkFunctions simulation not implemented"), nil, nil
	}
}
