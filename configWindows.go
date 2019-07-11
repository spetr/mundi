// +build windows

package main

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

func loadConfig() {
	var i uint64

	// Load config defaults
	/*
		configFile, err := os.Open(*configFileName)
		errorChecker(err)
		byteValue, _ := ioutil.ReadAll(configFile)
		configFile.Close()
		yaml.Unmarshal(byteValue, &config)
	*/

	regKeyHTTP, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\DigitalData\Mundi\HTTP`, registry.QUERY_VALUE)
	if err != nil {
		log.Println(err)
	}
	defer regKeyHTTP.Close()

	regKeyHTTPS, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\DigitalData\Mundi\HTTPS`, registry.QUERY_VALUE)
	if err != nil {
		log.Println(err)
	}
	defer regKeyHTTPS.Close()

	// HTTP
	i, _, err = regKeyHTTP.GetIntegerValue("Enabled")
	if err != nil {
		log.Println(err)
	} else {
		config.HTTP.Enabled = (i != 0)
	}
	config.HTTP.Address, _, err = regKeyHTTP.GetStringValue("Address")
	if err != nil {
		log.Println(err)
	}
	config.HTTP.Port, _, err = regKeyHTTP.GetStringValue("Port")
	if err != nil {
		log.Println(err)
	}

	// HTTPS
	i, _, err = regKeyHTTPS.GetIntegerValue("Enabled")
	if err != nil {
		log.Println(err)
	} else {
		config.HTTPS.Enabled = (i != 0)
	}
	config.HTTPS.Address, _, err = regKeyHTTPS.GetStringValue("Address")
	if err != nil {
		log.Println(err)
	}
	config.HTTPS.Port, _, err = regKeyHTTPS.GetStringValue("Port")
	if err != nil {
		log.Println(err)
	}
	i, _, err = regKeyHTTPS.GetIntegerValue("HTST")
	if err != nil {
		log.Println(err)
	} else {
		config.HTTPS.Htst = (i != 0)
	}
	config.HTTPS.Cert, _, err = regKeyHTTPS.GetStringValue("Cert")
	if err != nil {
		log.Println(err)
	}
	config.HTTPS.Key, _, err = regKeyHTTPS.GetStringValue("Key")
	if err != nil {
		log.Println(err)
	}
}
