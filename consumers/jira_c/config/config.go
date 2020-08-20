package config

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	config "consumers/jira_c/config/types"
)

// New reads the configuration from the file/Reader and parses it into a Config object
func New(r io.Reader) (config.Config, error) {
	configBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return config.Config{}, err
	}

	var newConfig config.Config
	err = yaml.Unmarshal(configBytes, &newConfig)
	if err != nil {
		return config.Config{}, err
	}
	return newConfig, nil
}
