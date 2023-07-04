package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
)

func TestUserDevicesMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	owner := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateUserDevices(ctx, &types.MsgCreateUserDevices{Owner: owner})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestUserDevicesMsgServerUpdate(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateUserDevices
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateUserDevices{Owner: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserDevices{Owner: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserDevices{Owner: owner, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateUserDevices(ctx, &types.MsgCreateUserDevices{Owner: owner})
			require.NoError(t, err)

			_, err = srv.UpdateUserDevices(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserDevicesMsgServerDelete(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteUserDevices
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteUserDevices{Owner: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteUserDevices{Owner: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteUserDevices{Owner: owner, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateUserDevices(ctx, &types.MsgCreateUserDevices{Owner: owner})
			require.NoError(t, err)
			_, err = srv.DeleteUserDevices(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
