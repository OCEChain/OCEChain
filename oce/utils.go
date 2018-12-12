package main

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
)

// ExportGenesisFile creates and writes the genesis configuration to disk. An
// error is returned if building or writing the configuration to file fails.
func ExportGenesisFile(
	genFile, chainID string, validators []types.GenesisValidator, appState json.RawMessage,
) error {

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		Validators: validators,
		AppState:   appState,
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

// ReadOrCreatePrivValidator read of create the private key file for this config
func ReadOrCreatePrivValidator(privValFile string) crypto.PubKey {
	var privValidator *privval.FilePV

	if common.FileExists(privValFile) {
		privValidator = privval.LoadFilePV(privValFile)
	} else {
		privValidator = privval.GenFilePV(privValFile)
		privValidator.Save()
	}

	return privValidator.GetPubKey()
}
