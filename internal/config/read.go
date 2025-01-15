package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	file_name = "/.gatorconfig.json"
)

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + file_name, nil

}

func Read() (*Config, error) {

	filePath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println(cfg)

	return &cfg, nil
}
