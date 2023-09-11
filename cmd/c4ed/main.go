package main

import (
	"github.com/chain4energy/c4e-chain/v2/cmd/c4ed/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	"os"

	"github.com/chain4energy/c4e-chain/v2/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
