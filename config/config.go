package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Config type to hold the values in the config.json
type Config struct {
	ProjectEnv string `json:"projectEnv"`
	AppName    string `json:"appName"`
	AppURI     string `json:"appURI"`
	Version    string `json:"version"`
	Server     struct {
		Address string `json:"address"`
	} `json:"server"`
	Data struct {
		URI    string `json:"uri"`
		DbName string `json:"dbName"`
	} `json:"data"`
	Discogs struct {
		URI    string `json:"uri"`
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"discogs"`
	Musicbrainz struct {
		URI string `json:"uri"`
	} `json:"musicbrainz"`
}

// Read takes the path to the config.json and Unmarshels it into the Config Type value
func (c *Config) Read(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "Failed to read file @ "+path)
	}
	err = json.Unmarshal(file, c)
	if err != nil {
		return errors.Wrap(err, "Failed to Unmarshal Json")
	}
	return nil
}
