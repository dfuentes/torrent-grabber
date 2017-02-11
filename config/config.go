package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Feeds []Feed `yaml:"feeds"`
}

type Feed struct {
	URL       string   `yaml:"url"`
	Filters   []string `yaml:"filters"`
	OutputDir string   `yaml:"output-dir"`
}

func Load(path string) (Config, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(contents, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
