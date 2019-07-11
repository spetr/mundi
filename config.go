package main

//Config file
type (
	tConfig struct {
		Socket struct {
			Enabled bool   `yaml:"enabled"`
			Path    string `yaml:"path"`
		} `yaml:"socket"`
		HTTP struct {
			Enabled bool   `yaml:"enabled"`
			Address string `yaml:"address"`
			Port    string `yaml:"port"`
		} `yaml:"http"`
		HTTPS struct {
			Enabled bool   `yaml:"enabled"`
			Address string `yaml:"address"`
			Port    string `yaml:"port"`
			Htst    bool   `yaml:"htst"`
			Cert    string `yaml:"cert"`
			Key     string `yaml:"key"`
		} `yaml:"https"`
	}
)

var config tConfig
