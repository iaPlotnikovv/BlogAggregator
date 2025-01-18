package app

import "github.com/iaPlotnikovv/BlogAggregator/internal/config"

type state struct {
	cfg *config.Config
}

func StateInit() *state {
	return &state{
		cfg: config.ConfigInit(),
	}
}
