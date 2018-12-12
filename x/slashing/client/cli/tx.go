package cli

import (
	"github.com/OCEChain/OCEChain/client/context"
	"github.com/OCEChain/OCEChain/client/utils"
	"github.com/OCEChain/OCEChain/codec"
	sdk "github.com/OCEChain/OCEChain/types"
	authcmd "github.com/OCEChain/OCEChain/x/auth/client/cli"
	authtxb "github.com/OCEChain/OCEChain/x/auth/client/txbuilder"
	"github.com/OCEChain/OCEChain/x/slashing"

	"github.com/spf13/cobra"
)

// GetCmdUnjail implements the create unjail validator command.
func GetCmdUnjail(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unjail",
		Args:  cobra.NoArgs,
		Short: "unjail validator previously jailed for downtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

			valAddr, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg := slashing.NewMsgUnjail(sdk.ValAddress(valAddr))
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg}, false)
			}
			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	return cmd
}
