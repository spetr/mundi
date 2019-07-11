// +build !windows

package main

import (
	"flag"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (
	configFileName = flag.String("config", "mundi.yaml", "Configuration file")
)

func loadConfig() {
	// Load config file
	configFile, err := os.Open(*configFileName)
	errorChecker(err)
	byteValue, _ := ioutil.ReadAll(configFile)
	configFile.Close()
	yaml.Unmarshal(byteValue, &config)
}
