package app

import (
	"github.com/iaPlotnikovv/BlogAggregator/internal/config"
	"github.com/iaPlotnikovv/BlogAggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func StateInit(db *database.Queries) *state {
	return &state{
		db:  db,
		cfg: config.ConfigInit(),
	}
}
