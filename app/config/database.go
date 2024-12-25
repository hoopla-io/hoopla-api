package config

import (
	"github.com/itv-go/yamlreader"
	"github.com/qahvazor/qahvazor/pkg/databasego"
)

func NewDatabaseConfig() *databasego.Config {
	read, err := yamlreader.Read("config/database.yaml", &databasego.Config{})
	if err != nil {
		return nil
	}

	return read
}
