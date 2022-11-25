package slashing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spike-engine/juno/modules"

	"github.com/spike-engine/bdjuno/v3/database"
	slashingsource "github.com/spike-engine/bdjuno/v3/modules/slashing/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.BlockModule   = &Module{}
)

// Module represent x/slashing module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source slashingsource.Source
}

// NewModule returns a new Module instance
func NewModule(source slashingsource.Source, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "slashing"
}
