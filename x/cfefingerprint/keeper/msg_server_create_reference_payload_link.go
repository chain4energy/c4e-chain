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
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) CreateReferencePayloadLink(goCtx context.Context, msg *types.MsgCreateReferencePayloadLink) (*types.MsgCreateReferencePayloadLinkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// get transaction bytes
	inputBytes := ctx.TxBytes()

	// create a TxHash commonly used as a transaction ID
	hash := sha256.Sum256(inputBytes)
	txHash := hex.EncodeToString(hash[:])

	// create referenceID
	referenceID, err := createReferenceID(64, txHash)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "failed to generate referenceID")
	}

	// create reference payload link
	referenceKey := util.CalculateHash(referenceID)
	referenceValue := util.CalculateHash(util.HashConcat(referenceID, msg.PayloadHash))

	ctx.Logger().Debug("referenceKey   = %s", referenceKey)

	// publish reference payload link

	// Check if a Payload Link was already stored at the given key
	if !(k.CheckIfPayloadLinkExists(ctx, referenceKey)) {
		return nil, sdkerrors.Wrap(types.ErrAlreadyExists, "data was found at the given key, cannot overwrite present payloadlinks")
	}

	// store payload link
	k.AppendPayloadLink(ctx, referenceKey, referenceValue)

	return &types.MsgCreateReferencePayloadLinkResponse{ReferenceId: referenceID}, nil
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
