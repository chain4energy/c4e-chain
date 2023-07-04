package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
)

func TestUserCertificatesMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	owner := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateUserCertificates(ctx, &types.MsgCreateUserCertificates{Owner: owner})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestUserCertificatesMsgServerUpdate(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateUserCertificates
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateUserCertificates{Owner: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserCertificates{Owner: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserCertificates{Owner: owner, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateUserCertificates(ctx, &types.MsgCreateUserCertificates{Owner: owner})
			require.NoError(t, err)

			_, err = srv.UpdateUserCertificates(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserCertificatesMsgServerDelete(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteUserCertificates
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteUserCertificates{Owner: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteUserCertificates{Owner: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteUserCertificates{Owner: owner, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateUserCertificates(ctx, &types.MsgCreateUserCertificates{Owner: owner})
			require.NoError(t, err)
			_, err = srv.DeleteUserCertificates(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
