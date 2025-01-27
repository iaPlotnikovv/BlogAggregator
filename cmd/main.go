package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/iaPlotnikovv/BlogAggregator/internal/app"
	"github.com/iaPlotnikovv/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

const (
	dbURL = "postgres://iplotnikow:@localhost:5432/gator?sslmode=disable"
)

func main() {

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error in db connection")
	}
	dbQueries := database.New(db)

	state := app.StateInit(dbQueries)

	cmd := app.CmdListInit()

	app.AddCommands(cmd)

	input := os.Args

	if len(input) < 2 {
		log.Fatal("Error! Unknown command. Use <gator help>")
	}

	cmd.Run(state, app.CommandInit(input[1], input[2:]))

}
