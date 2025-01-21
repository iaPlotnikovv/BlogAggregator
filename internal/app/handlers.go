package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iaPlotnikovv/BlogAggregator/internal/config"
	"github.com/iaPlotnikovv/BlogAggregator/internal/database"
)

func ArgsChecker(cmd command) (string, error) {

	fmt.Println("Checking args...")

	switch {
	case len(cmd.arg) == 1:
		//fmt.Println("here.")
		return cmd.arg[0], nil

	case len(cmd.arg) == 0:
		//fmt.Println("here.!")
		return "", fmt.Errorf("\n`%s` expects a single argument", cmd.name)

	case len(cmd.arg) > 1:
		return "", fmt.Errorf("\n`%s` expects a !single! argument", cmd.name)

	default:
		return "", errors.New("\nSomething went wrong! try again")

	}

}

func HandlerLogin(s *state, cmd command) error {

	arg, err := ArgsChecker(cmd)
	if err != nil {
		log.Fatal(err)
	}
	if exist, err := s.db.ExistsUser(context.Background(), arg); err != nil {
		log.Fatal(err)
	} else if !exist {
		log.Fatal("User doesn't exist!")
		//	os.Exit(1)
	}

	s.cfg.SetUser(arg)

	fmt.Printf("\nUser successfully set!\n")
	return nil
}

func HandlerRegister(s *state, cmd command) error {
	arg, err := ArgsChecker(cmd)
	if err != nil {
		log.Fatal(err)
	}

	if exist, err := s.db.ExistsUser(context.Background(), arg); err != nil {
		log.Fatal(err)
	} else if exist {
		log.Fatal("User is already exists!")
		//os.Exit(1)
	}

	data := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      arg,
	}

	user, err := s.db.CreateUser(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("\nUser %s was created!\n", user.Name)
	fmt.Printf("\nUser Data:\n ID=%s Name=%s CreatedAt=%s\n", user.ID, user.Name, user.CreatedAt)

	return nil
}

func ConfigChecker(s *state, cmd command) error {
	if len(cmd.arg) != 0 {
		log.Fatal("Invalid use of command!")
	}

	fmt.Printf("\nThe current state:\n")
	current, err := config.Read()
	if err != nil {
		return err
	}
	fmt.Println(current)
	return nil
}
