package keeper

import (
	"context"
	"crypto/sha256"

	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/util"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) CreateReferencePayloadLink(goCtx context.Context, msg *types.MsgCreateReferencePayloadLink) (*types.MsgCreateReferencePayloadLinkResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create reference payloadLink")
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Debug("create payload link for a given payload: ", msg.PayloadHash)

	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// get transaction bytes
	inputBytes := ctx.TxBytes()

	// create a TxHash commonly used as a transaction ID
	hash := sha256.Sum256(inputBytes)
	txHash := hex.EncodeToString(hash[:])

	// create referenceId
	referenceId, err := createReferenceID(64, txHash)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "failed to generate referenceID")
	}
	ctx.Logger().Debug("calculated referenceID = %s", referenceId)

	// create reference payload link
	referenceKey := util.CalculateHash(referenceId)
	referenceValue := util.CalculateHash(util.HashConcat(referenceId, msg.PayloadHash))

	ctx.Logger().Debug("calculated referenceKey = %s / referenceValue = %s", referenceKey, referenceValue)

	// publish reference payload link

	// Check if a Payload Link was already stored at the given key
	if !(k.CheckIfPayloadLinkExists(ctx, referenceKey)) {
		return nil, sdkerrors.Wrap(types.ErrAlreadyExists, "data was found at the given key, cannot overwrite present payloadlinks")
	}

	// store payload link
	k.AppendPayloadLink(ctx, referenceKey, referenceValue)

	/*
		As long as referenceId cannot be correlated with a specific account address it can be included in the emitted event.
	*/

	// emit related newPayloadLinkEvent
	newPayloadLinkEvent := &types.NewPayloadLink{
		ReferenceId:    referenceId,
		ReferenceKey:   referenceKey,
		ReferenceValue: referenceKey,
	}
	err = ctx.EventManager().EmitTypedEvent(newPayloadLinkEvent)

	return &types.MsgCreateReferencePayloadLinkResponse{ReferenceId: referenceId}, nil
}

func createReferenceID(length int, txHash string) (string, error) {

	// convert to bytes
	dataBytes := []byte(txHash)
	seed := binary.BigEndian.Uint64(dataBytes)

	rand.Seed(int64(seed))
	b := make([]byte, length+2)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b)[2 : length+2], nil
}
