package config

import (
	"github.com/hoopla/hoopla-api/pkg/databasego"
	"github.com/itv-go/yamlreader"
)

func NewDatabaseConfig() *databasego.Config {
	read, err := yamlreader.Read("config/database.yaml", &databasego.Config{})
	if err != nil {
		return nil
	}

	return read
}
