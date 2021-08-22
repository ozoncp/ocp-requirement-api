package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config Config

type db struct {
	Dsn string `yaml:"dsn"`
}

type grpc struct {
	Address string `yaml:"address"`
}

type rest struct {
	Address string `yaml:"address"`
}

type Config struct {
	Grpc     grpc `yaml:"grpc"`
	Rest     rest `yaml:"rest"`
	Database db   `yaml:"db"`
}

func ReadConfigYAML(path string) (Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return Config{}, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Failed to close the file.")
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
