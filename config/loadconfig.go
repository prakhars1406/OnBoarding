package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Configuration struct {
	PORT                string
	APP_MODE            string
	DATASTORE                                  string
	MONGO_SERVER                               string
	LOG_STRATEGY        int64
}

var (
	configuration     *Configuration = nil
	configFile        *string        = nil
	logStrategy       int            = 0
)

//defined all the required flags
func init() {
	configFile = flag.String(FILE, EMPTY_STRING, A_STRING)
}

func LoadAppConfiguration() {

	if configuration == nil {

		flag.Parse()
		if len(*configFile) == 0 {
			fmt.Println("Mandatory arguments are missing for App Execution")
			StopService("Mandatory arguments not provided for executing the App")
		}
		configuration = loadConfiguration(*configFile)
	}
}

func loadConfiguration(filename string) *Configuration {
	if configuration == nil {
		configFile, err := os.Open(filename)
		defer configFile.Close()
		if err != nil {
			StopService(err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		err1 := jsonParser.Decode(&configuration)
		if err1 != nil {
			fmt.Println("Failed to parse configuration file")
			StopService(err1.Error())
		}
	}
	return configuration
}

func GetAppConfiguration() *Configuration {
	if configuration == nil {
		LoadAppConfiguration()
	}
	return configuration
}

func StopService(log string) {
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(os.Kill)
}
