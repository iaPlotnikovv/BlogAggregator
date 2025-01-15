package config

import (
	"encoding/json"
	"os"
)

func Write(cfg *Config) error {

	cfg.SetUser()
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)

	return err

}
