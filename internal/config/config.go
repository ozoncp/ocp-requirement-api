package config

import (
	"gopkg.in/yaml.v2"
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

type jaeger struct {
	Address string `yaml:"address"`
}

type prometheus struct {
	Address  string `yaml:"address"`
	Endpoint string `yaml:"endpoint"`
}

type kafka struct {
	Addresses []string `yaml:"addresses"`
	Topic     string   `yaml:"topic"`
}
type Config struct {
	ServiceName string     `yaml:"service-name"`
	Grpc        grpc       `yaml:"grpc"`
	Rest        rest       `yaml:"rest"`
	Database    db         `yaml:"db"`
	Jaeger      jaeger     `yaml:"jaeger"`
	Prometheus  prometheus `yaml:"prometheus"`
	Kafka       kafka      `yaml:"kafka"`
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
