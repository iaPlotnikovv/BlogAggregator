package app

import (
	"fmt"
	"log"
)

type cmdHandler func(*state, command) error

type commands struct {
	CmdList map[string]cmdHandler
}

type command struct {
	name string
	arg  []string
}

func CommandInit(name string, args []string) command {
	return command{
		name: name,
		arg:  args,
	}
}

func CmdListInit() *commands {
	cmd := &commands{
		CmdList: make(map[string]cmdHandler),
	}
	return cmd

}
func (c *commands) Register(name string, f cmdHandler) error {
	c.CmdList[name] = f
	return nil
}
func (c *commands) Run(s *state, cmd command) error {
	if handler, ok := c.CmdList[cmd.name]; ok {
		err := handler(s, cmd)
		if err != nil {
			return fmt.Errorf("\nerror in run!: %v", err)
		}
	} else {
		log.Fatal("\nUnknown command!\n\n")

	}
	return nil
}

func (c *commands) Help(s *state, cmd command) error {

	if len(cmd.arg) != 0 {
		log.Fatal("Invalid use of command!")
	} else {
		fmt.Println("\nAvailable commands:")
		for command := range c.CmdList {
			fmt.Println(command)
		}

	}
	return nil

}
