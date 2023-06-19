package keeper

import (
	"cosmossdk.io/errors"
	"crypto/sha256"
	"encoding/hex"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CreatePayloadLink(ctx sdk.Context, payloadHash string) (*string, error) {
	k.Logger(ctx).Debug("create payload link", "payloadHash", payloadHash)
	if payloadHash == "" {
		return nil, errors.Wrapf(c4eerrors.ErrParam, "payloadHash cannot be empty")
	}

	referenceId, err := createReferenceIdFromTxHash(ctx)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrLogic, "failed to generate referenceId (%s)", err.Error())
	}
	referenceKey := types.CalculateHash(referenceId)
	referenceValue := types.CalculatePayloadLink(referenceId, payloadHash)
	k.Logger(ctx).Debug("calculated payload link values", "referenceId", referenceId, "referenceKey", referenceKey, "referenceValue", referenceValue)

	if err = k.ValidatePayloadLinkDoesntExist(ctx, referenceKey); err != nil {
		return nil, err
	}

	k.AppendPayloadLink(ctx, referenceKey, referenceValue)

	// TODO: reconsider hidding referenceKey and referenceValue from the emitted event.
	// Anyone can see the transaction and the payloadHash contained in that transaction. Having access to the code how the
	// referenceId and referenceKey are created, you can retrieve the referenceValue. Consider how to hide these values.
	newPayloadLinkEvent := &types.EventNewPayloadLink{
		ReferenceId:    referenceId,
		ReferenceKey:   referenceKey,
		ReferenceValue: referenceValue,
	}
	if err = ctx.EventManager().EmitTypedEvent(newPayloadLinkEvent); err != nil {
		k.Logger(ctx).Error("new payload link emit event error", "error", err.Error())
	}

	return &referenceId, nil
}

func createReferenceIdFromTxHash(ctx sdk.Context) (string, error) {
	// get transaction bytes
	inputBytes := ctx.TxBytes()
	// create a TxHash commonly used as a transaction ID
	hash := sha256.Sum256(inputBytes)
	txHash := hex.EncodeToString(hash[:])

	return types.CreateReferenceID(64, txHash)
}

func (k Keeper) ValidatePayloadLinkDoesntExist(ctx sdk.Context, referenceKey string) error {
	_, found := k.GetPayloadLink(ctx, referenceKey)
	if found {
		return errors.Wrapf(c4eerrors.ErrAlreadyExists, "payload link with reference key %s already exists", referenceKey)
	}
	return nil
}

func (k Keeper) VerifyPayloadLink(ctx sdk.Context, referenceId, payloadHash string) error {
	payloadLink, err := k.MustGetPayloadLinkByReferenceId(ctx, referenceId)
	if err != nil {
		return err
	}

	expectedPayloadLink := types.CalculatePayloadLink(referenceId, payloadHash)

	if expectedPayloadLink != payloadLink {
		return errors.Wrapf(c4eerrors.ErrParam, "expected payload link %s for reference id %s doesn't match payload link %s", expectedPayloadLink, referenceId, payloadLink)
	}

	return nil
}
