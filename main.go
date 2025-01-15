package main

import (
	"log"

	cfg "github.com/iaPlotnikovv/BlogAggregator/internal/config"
)

func main() {

	data, err := cfg.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}
