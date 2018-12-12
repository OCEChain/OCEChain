package rest

import (
	"github.com/OCEChain/OCEChain/client/context"
	"github.com/OCEChain/OCEChain/codec"
	"github.com/OCEChain/OCEChain/crypto/keys"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	registerQueryRoutes(cliCtx, r, cdc)
	registerTxRoutes(cliCtx, r, cdc, kb)
}
