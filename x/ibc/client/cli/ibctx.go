package cli

import (
	"encoding/hex"

	"github.com/OCEChain/OCEChain/client"
	"github.com/OCEChain/OCEChain/client/context"
	"github.com/OCEChain/OCEChain/client/utils"
	codec "github.com/OCEChain/OCEChain/codec"
	sdk "github.com/OCEChain/OCEChain/types"
	authcmd "github.com/OCEChain/OCEChain/x/auth/client/cli"
	authtxb "github.com/OCEChain/OCEChain/x/auth/client/txbuilder"
	"github.com/OCEChain/OCEChain/x/ibc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagTo     = "to"
	flagAmount = "amount"
	flagChain  = "chain"
)

// IBCTransferCmd implements the IBC transfer command.
func IBCTransferCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

			from, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg, err := buildMsg(from)
			if err != nil {
				return err
			}
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTo, "", "Address to send coins")
	cmd.Flags().String(flagAmount, "", "Amount of coins to send")
	cmd.Flags().String(flagChain, "", "Destination chain to send coins")

	return cmd
}

func buildMsg(from sdk.AccAddress) (sdk.Msg, error) {
	amount := viper.GetString(flagAmount)
	coins, err := sdk.ParseCoins(amount)
	if err != nil {
		return nil, err
	}

	dest := viper.GetString(flagTo)
	bz, err := hex.DecodeString(dest)
	if err != nil {
		return nil, err
	}
	to := sdk.AccAddress(bz)

	packet := ibc.NewIBCPacket(from, to, coins, viper.GetString(client.FlagChainID),
		viper.GetString(flagChain))

	msg := ibc.IBCTransferMsg{
		IBCPacket: packet,
	}

	return msg, nil
}
