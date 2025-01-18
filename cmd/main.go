package main

import (
	"log"
	"os"

	"github.com/iaPlotnikovv/BlogAggregator/internal/app"
)

func main() {

	state := app.StateInit()

	cmd := app.CmdListInit()

	cmd.Register("login", app.HandlerLogin)
	cmd.Register("help", cmd.Help)

	input := os.Args
	if len(input) < 2 {
		log.Fatal("Error! Unknown command. Use <gator help>")
	}

	cmd.Run(state, app.CommandInit(input[1], input[2:]))

}
