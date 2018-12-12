package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/OCEChain/OCEChain/client/context"
	"github.com/OCEChain/OCEChain/client/utils"
	"github.com/OCEChain/OCEChain/codec"
	"github.com/OCEChain/OCEChain/crypto/keys/keyerror"
	"github.com/OCEChain/OCEChain/x/auth"
	authtxb "github.com/OCEChain/OCEChain/x/auth/client/txbuilder"
)

// SignBody defines the properties of a sign request's body.
type SignBody struct {
	Tx               auth.StdTx `json:"tx"`
	LocalAccountName string     `json:"name"`
	Password         string     `json:"password"`
	ChainID          string     `json:"chain_id"`
	Sequence         int64      `json:"sequence"`
	AppendSig        bool       `json:"append_sig"`
}

// nolint: unparam
// sign tx REST handler
func SignTxRequestHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var m SignBody

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = cdc.UnmarshalJSON(body, &m)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txBldr := authtxb.TxBuilder{
			ChainID:  m.ChainID,
			Sequence: m.Sequence,
		}

		signedTx, err := txBldr.SignStdTx(m.LocalAccountName, m.Password, m.Tx, m.AppendSig)
		if keyerror.IsErrKeyNotFound(err) {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		} else if keyerror.IsErrWrongPassword(err) {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		} else if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, signedTx, cliCtx.Indent)
	}
}
