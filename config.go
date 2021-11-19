package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

var configPath = getAppFolder() + "config.yml"

func readConfig() map[string]interface{} {
	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Got error while reading %s config file: %v\n", configPath, err)
	}

	var confObj map[string]interface{}
	err = yaml.Unmarshal(f, &confObj)
	if err != nil {
		log.Fatalf("Got error while unmarshalling %s config file: %v\n", configPath, err)
	}
	return confObj
}

func GetPort() int {
	conf := readConfig()
	return conf["port"].(int)
}

func GetKeys() []string {
	conf := readConfig()
	var keys []string
	for _, key := range conf["keys"].([]interface{}) {
		keys = append(keys, key.(string))
	}
	return keys
}

func getAppFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	sep := string(os.PathSeparator)
	return fmt.Sprint(usr.HomeDir, sep, ".arma_drm_server", sep)
}
