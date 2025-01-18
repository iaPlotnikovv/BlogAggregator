package app

import (
	"fmt"
	"log"

	"github.com/iaPlotnikovv/BlogAggregator/internal/config"
)

func HandlerLogin(s *state, cmd command) error {

	switch {
	case len(cmd.arg) == 1:

		s.cfg.SetUser(cmd.arg[0])

		fmt.Printf("\nUser successfully set!\n")
		fmt.Printf("\nThe current state:\n")
		current, err := config.Read()
		if err != nil {
			return err
		}
		fmt.Println(current)

	case len(cmd.arg) == 0:
		log.Fatal("\n`login` expects a single argument, the username.\n")

	case len(cmd.arg) > 1:
		log.Fatal("\n`login` expects a !single! argument, the username.\n")

	default:
		log.Fatal("\nSomething went wrong! try again...\n")

	}

	return nil
}

func HandlerHelp(s *state, cmd command) error {
	var list *commands

	if len(cmd.arg) != 0 {
		log.Fatal("Invalid use of command!")
	} else {
		fmt.Println("\nAvailable commands:")
		for command, _ := range list.CmdList {
			fmt.Println(command)
		}

	}
	return nil

}
