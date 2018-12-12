package context

import (
	sdk "github.com/OCEChain/OCEChain/types"
	"github.com/OCEChain/OCEChain/x/auth"
)

// StdSignMsg is a convenience structure for passing along
// a Msg with the other requirements for a StdSignDoc before
// it is signed. For use in the CLI.
type StdSignMsg struct {
	ChainID  string      `json:"chain_id"`
	Sequence int64       `json:"sequence"`
	Fee      auth.StdFee `json:"fee"`
	Msgs     []sdk.Msg   `json:"msgs"`
	Memo     string      `json:"memo"`
}

// get message bytes
func (msg StdSignMsg) Bytes() []byte {
	return auth.StdSignBytes(msg.ChainID, msg.Sequence, msg.Fee, msg.Msgs, msg.Memo)
}
