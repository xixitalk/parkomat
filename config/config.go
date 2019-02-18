package config

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Web struct {
	IP        string `json:"ip" toml:"ip"`
	Port      int    `json:"port" toml:"port"`
	Path      string `json:"path" toml:"path"`
	AccessLog string `json:"access_log" toml:"access_log"`
	ErrorLog  string `json:"error_log" toml:"error_log"`
}

type WebDav struct {
	Enabled  bool   `json:"enabled" toml:"enabled"`
	Username string `json:"username" toml:"username"`
	Password string `json:"password" toml:"password"`
	Mount    string `json:"mount" toml:"mount"`
}

type Config struct {
	Web    Web    `json:"web" toml:"web"`
	WebDav WebDav `json:"webdav" toml:"webdav"`
}

func (config *Config) Refresh() {
	// TODO: not implemented
	// This function should refresh configuration
	// Without the need of restarting the service
}

// NewConfigFromFile reads configuration from specified filename
func NewConfigFromFile(filename string) (config *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if _, err := toml.Decode(string(file), &config); err != nil {
		return nil, err
	}

	config.normalizeData()
	return
}

// normalizeData makes sure all domains are in the same format
func (config *Config) normalizeData() {
}

// NewConfigFromJSONFile reads configuration from JSON formatted file
func NewConfigFromJSONFile(filename string) (config *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	config.normalizeData()
	return
}

