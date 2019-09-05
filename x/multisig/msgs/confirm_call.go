package msgs

import (
	"wings-blockchain/x/multisig/types"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Message to confirm call
type MsgConfirmCall struct {
	MsgId  uint64		  `json:"msg_id"`
	Sender sdk.AccAddress `json:"sender"`
}

func NewMsgConfirmCall(msgId uint64, sender sdk.AccAddress) MsgConfirmCall {
	return MsgConfirmCall{
		MsgId:  msgId,
		Sender: sender,
	}
}

func (msg MsgConfirmCall) Route() string {
	return types.DefaultRoute
}

func (msg MsgConfirmCall) Type() string {
	return "confirm_call"
}

func (msg MsgConfirmCall) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}

	return nil
}

func (msg MsgConfirmCall) GetSignBytes() []byte {
	bc, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(bc)
}

func (msg MsgConfirmCall) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}