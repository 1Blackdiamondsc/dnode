package poa

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"wings-blockchain/x/poa/msgs"
	"wings-blockchain/x/poa/types"
)

// New message handler for PoA module
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgAddValidator:
			return handleMsgAddValidator(ctx, keeper, msg)

		case msgs.MsgReplaceValidator:
			return handleMsgReplaceValidator(ctx, keeper, msg)

		case msgs.MsgRemoveValidator:
			return handleMsgRemoveValidator(ctx, keeper, msg)

		default:
			return sdk.ErrUnknownRequest(fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())).Result()
		}
	}
}

// Handle MsgAddValidator for add new validator
func handleMsgAddValidator(ctx sdk.Context, keeper Keeper, msg msgs.MsgAddValidator) sdk.Result {
	if keeper.HasValidator(ctx, msg.Address) {
		return types.ErrValidatorExists(msg.Address.String()).Result()
	}

	maxValidators := keeper.GetMaxValidators(ctx)
	amount := keeper.GetValidatorAmount(ctx)

	if amount+1 > maxValidators {
		return types.ErrMaxValidatorsReached(maxValidators).Result()
	}

	keeper.AddValidator(ctx, msg.Address, msg.EthAddress)

	return sdk.Result{}
}

// Handle MsgRemoveValidator for remove validator
func handleMsgRemoveValidator(ctx sdk.Context, keeper Keeper, msg msgs.MsgRemoveValidator) sdk.Result {
	if !keeper.HasValidator(ctx, msg.Address) {
		return types.ErrValidatorDoesntExists(msg.Address.String()).Result()
	}

	minValidators := keeper.GetMinValidators(ctx)
	amount := keeper.GetValidatorAmount(ctx)

	if amount-1 < minValidators {
		return types.ErrMinValidatorsReached(minValidators).Result()
	}

	keeper.RemoveValidator(ctx, msg.Address)

	return sdk.Result{}
}

// Handle MsgReplaceValidator for replace validator
func handleMsgReplaceValidator(ctx sdk.Context, keeper Keeper, msg msgs.MsgReplaceValidator) sdk.Result {
	if !keeper.HasValidator(ctx, msg.OldValidator) {
		return types.ErrValidatorDoesntExists(msg.OldValidator.String()).Result()
	}

	if keeper.HasValidator(ctx, msg.NewValidator) {
		return types.ErrValidatorExists(msg.NewValidator.String()).Result()
	}

	keeper.ReplaceValidator(ctx, msg.OldValidator, msg.NewValidator, msg.EthAddress)

	return sdk.Result{}
}
