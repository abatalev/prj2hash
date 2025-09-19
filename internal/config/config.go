package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Excludes []string `yaml:"excludes"` // DEPRECATED! Remove in next version
	Rules    []string `yaml:"rules"`
}

func readConfig(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func LoadConfig(fileName string) *Config {
	if _, err := os.Stat(fileName); err != nil {
		return &Config{Excludes: []string{}}
	}
	cfg, _ := readConfig(fileName)
	// if err != nil {
	// 	os.(1)
	// }
	return cfg
}
