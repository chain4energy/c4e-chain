package keeper_test

import (
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
)

func TestAirdropEntryMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	owner := acountsAddresses[0].String()
	_, err := srv.AddAirdropEntries(ctx, &types.MsgAddAirdropEntries{Owner: owner})
	require.NoError(t, err)
}

//
//func TestAirdropEntryMsgServerUpdate(t *testing.T) {
//	creator := "A"
//
//	for _, tc := range []struct {
//		desc    string
//		request *types.MsgUpdateAirdropEntry
//		err     error
//	}{
//		{
//			desc:    "Completed",
//			request: &types.MsgUpdateAirdropEntry{Creator: creator},
//		},
//		{
//			desc:    "Unauthorized",
//			request: &types.MsgUpdateAirdropEntry{Creator: "B"},
//			err:     sdkerrors.ErrUnauthorized,
//		},
//		{
//			desc:    "Unauthorized",
//			request: &types.MsgUpdateAirdropEntry{Creator: creator, Id: 10},
//			err:     sdkerrors.ErrKeyNotFound,
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			srv, ctx := setupMsgServer(t)
//			_, err := srv.CreateAirdropEntry(ctx, &types.MsgCreateAirdropEntry{Creator: creator})
//			require.NoError(t, err)
//
//			_, err = srv.UpdateAirdropEntry(ctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestAirdropEntryMsgServerDelete(t *testing.T) {
//	creator := "A"
//
//	for _, tc := range []struct {
//		desc    string
//		request *types.MsgDeleteAirdropEntry
//		err     error
//	}{
//		{
//			desc:    "Completed",
//			request: &types.MsgDeleteAirdropEntry{Creator: creator},
//		},
//		{
//			desc:    "Unauthorized",
//			request: &types.MsgDeleteAirdropEntry{Creator: "B"},
//			err:     sdkerrors.ErrUnauthorized,
//		},
//		{
//			desc:    "KeyNotFound",
//			request: &types.MsgDeleteAirdropEntry{Creator: creator, Id: 10},
//			err:     sdkerrors.ErrKeyNotFound,
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			srv, ctx := setupMsgServer(t)
//
//			_, err := srv.CreateAirdropEntry(ctx, &types.MsgCreateAirdropEntry{Creator: creator})
//			require.NoError(t, err)
//			_, err = srv.DeleteAirdropEntry(ctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
