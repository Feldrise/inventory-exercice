package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg Config

type Config struct {
	Settings struct {
		AuthKey string `yaml:"authKey"`
	} `yaml:"settings"`
	Database struct {
		Name             string `yaml:"name"`
		ConnectionString string `yaml:"connectionString"`
	} `yaml:"database"`
	Collections struct {
		Inventories    string `yaml:"inventories"`
		InventoryItems string `yaml:"inventoryItems"`
		Users          string `yaml:"users"`
	} `yaml:"collections"`
}

func Init(configPath string) {
	file, err := os.Open(configPath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Cfg)

	if err != nil {
		log.Fatal(err)
	}

}
