package types

func NewMsgSubmitDataFromChainlinkFunctions(creator string, data string) *MsgSubmitDataFromChainlinkFunctions {
	return &MsgSubmitDataFromChainlinkFunctions{
		Creator: creator,
		Data:    data,
	}
}
