package config

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// New reads the configuration from the file/Reader and parses it into a Config object
func New(r io.Reader) (Config, error) {
	configBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return Config{}, err
	}

	var newConfig Config
	err = yaml.Unmarshal(configBytes, &newConfig)
	if err != nil {
		return Config{}, err
	}
	return newConfig, nil
}
