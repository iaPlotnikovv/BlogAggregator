package config

import (
	"encoding/json"
	"os"
)

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

	//fmt.Println(cfg)

	return &cfg, nil
}
