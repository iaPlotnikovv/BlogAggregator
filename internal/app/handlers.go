package app

import (
	"context"
	"errors"
	"fmt"
	"log"

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
		handleError(err)
	}
	if exist, err := s.db.ExistsUser(context.Background(), arg); err != nil {
		handleError(err)
	} else if !exist {
		handleError(fmt.Errorf("user doesn't exist"))
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
		ID:   uuid.New(),
		Name: arg,
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
		handleError(fmt.Errorf("invalid use of command"))
	}

	fmt.Printf("\nThe current state:\n")
	current, err := config.Read()
	if err != nil {
		handleError(err)
	}
	fmt.Println(current)
	return nil
}

func HandlerReset(s *state, cmd command) error {

	if len(cmd.arg) != 0 {
		handleError(fmt.Errorf("\n expects a single argument"))
	}
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		handleError(err)
	}
	fmt.Println("\nUsers are succssesfully reset!")
	return nil
}

func HandlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		handleError(err)
	}
	for _, v := range users {
		if v == s.cfg.Current_user_name {
			fmt.Printf("\n* %s (current)\n", v)
		} else {
			fmt.Printf("\n* %s\n", v)
		}
	}

	return nil
}

func HandlerAgg(s *state, cmd command) error {
	arg, err := ArgsChecker(cmd)
	if err != nil {
		handleError(err)
	}
	//	arg := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), arg)
	if err != nil {
		handleError(err)
	}
	fmt.Println(feed)
	return nil
}

func HandlerAddFeed(s *state, cmd command) error {
	if len(cmd.arg) != 2 {
		handleError(errors.New("expects !two! arguments"))
	}
	var username = s.cfg.Current_user_name

	user_id, err := s.db.GetCurrentUserID(context.Background(), username)
	if err != nil {
		handleError(err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		UserID: user_id,
		Name:   cmd.arg[0],
		Url:    cmd.arg[1],
	})
	if err != nil {
		handleError(err)
	}
	fmt.Println(feed)
	return nil
}

func HandlerListFeeds(s *state, cmd command) error {
	if len(cmd.arg) > 1 {
		handleError(errors.New("invalid usage of command"))
	}
	feed, err := s.db.ListFeed(context.Background())
	if err != nil {
		handleError(err)
	}

	for _, v := range feed {
		fmt.Println(v)
	}
	return nil

}
