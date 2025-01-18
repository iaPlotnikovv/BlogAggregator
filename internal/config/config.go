package config

import (
	"log"
	"os"
)

const (
	file_name = "/.gatorconfig.json"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func ConfigInit() *Config {
	cfg, err := Read()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func (c *Config) SetUser(username string) error {
	c.Current_user_name = username
	err := Write(c)
	if err != nil {
		return err
	}
	return nil
}

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + file_name, nil

}
