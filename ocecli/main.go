package main

import (
	"github.com/OCEChain/OCEChain/app"
	"github.com/OCEChain/OCEChain/client"
	"github.com/OCEChain/OCEChain/client/keys"
	"github.com/OCEChain/OCEChain/client/lcd"
	_ "github.com/OCEChain/OCEChain/client/lcd/statik"
	"github.com/OCEChain/OCEChain/client/rpc"
	"github.com/OCEChain/OCEChain/client/tx"
	"github.com/OCEChain/OCEChain/version"
	authcmd "github.com/OCEChain/OCEChain/x/auth/client/cli"
	bankcmd "github.com/OCEChain/OCEChain/x/bank/client/cli"
	ibccmd "github.com/OCEChain/OCEChain/x/ibc/client/cli"
	slashingcmd "github.com/OCEChain/OCEChain/x/slashing/client/cli"
	stakecmd "github.com/OCEChain/OCEChain/x/stake/client/cli"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "ocecli",
		Short: "OCEChain light-client",
	}
)

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := app.MakeCodec()

	// TODO: Setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc.

	// add standard rpc, and tx commands
	rpc.AddCommands(rootCmd)
	rootCmd.AddCommand(client.LineBreak)
	tx.AddCommands(rootCmd, cdc)
	rootCmd.AddCommand(client.LineBreak)

	// add query/post commands (custom to binary)
	rootCmd.AddCommand(
		client.GetCommands(
			stakecmd.GetCmdQueryValidator("stake", cdc),
			stakecmd.GetCmdQueryValidators("stake", cdc),
			stakecmd.GetCmdQueryValidatorUnbondingDelegations("stake", cdc),
			stakecmd.GetCmdQueryValidatorRedelegations("stake", cdc),
			stakecmd.GetCmdQueryDelegation("stake", cdc),
			stakecmd.GetCmdQueryDelegations("stake", cdc),
			stakecmd.GetCmdQueryPool("stake", cdc),
			stakecmd.GetCmdQueryParams("stake", cdc),
			stakecmd.GetCmdQueryUnbondingDelegation("stake", cdc),
			stakecmd.GetCmdQueryUnbondingDelegations("stake", cdc),
			stakecmd.GetCmdQueryRedelegation("stake", cdc),
			stakecmd.GetCmdQueryRedelegations("stake", cdc),
			slashingcmd.GetCmdQuerySigningInfo("slashing", cdc),
			authcmd.GetAccountCmd("acc", cdc, app.GetAccountDecoder(cdc)),
		)...)

	rootCmd.AddCommand(
		client.PostCommands(
			bankcmd.SendTxCmd(cdc),
			ibccmd.IBCTransferCmd(cdc),
			ibccmd.IBCRelayCmd(cdc),
			stakecmd.GetCmdCreateValidator(cdc),
			stakecmd.GetCmdEditValidator(cdc),
			stakecmd.GetCmdDelegate(cdc),
			stakecmd.GetCmdUnbond("stake", cdc),
			stakecmd.GetCmdRedelegate("stake", cdc),
			slashingcmd.GetCmdUnjail(cdc),
		)...)

	// add proxy, version and key info
	rootCmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		// Note: Handle with #870
		panic(err)
	}
}
