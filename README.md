## Chainlink Functions - Cosmos Chain Integration

This repo is a fork from [Chainlink Functions Starter Kit](https://github.com/smartcontractkit/functions-hardhat-starter-kit), which is an example showing how to fetch off-chain data with Chainlink Functions.

To make data from Chainlink Functions arrive at a Cosmos chain, this forked repo does three things:

- Emit a data event in the Functions consumer contract
- Add a new `relayer` component that filters data event from the Ethereum address of the consumer contract, and submit it to a Cosmos chain
- Add a `cosmos` chain that has a Module to store that data in the Module Keeper