package chainlink

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	chainlinksimulation "github.com/haanhvu/chainlink-cosmos/cosmos/x/chainlink/simulation"
	"github.com/haanhvu/chainlink-cosmos/cosmos/x/chainlink/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	chainlinkGenesis := types.GenesisState{
		Params: types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&chainlinkGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgSubmitDataFromChainlinkFunctions          = "op_weight_msg_chainlink"
		defaultWeightMsgSubmitDataFromChainlinkFunctions int = 100
	)

	var weightMsgSubmitDataFromChainlinkFunctions int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitDataFromChainlinkFunctions, &weightMsgSubmitDataFromChainlinkFunctions, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitDataFromChainlinkFunctions = defaultWeightMsgSubmitDataFromChainlinkFunctions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitDataFromChainlinkFunctions,
		chainlinksimulation.SimulateMsgSubmitDataFromChainlinkFunctions(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
