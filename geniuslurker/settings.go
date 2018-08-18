package geniuslurker

import (
	"os"

	"github.com/spf13/viper"
)

//InitSettings initializes settings from config file
func InitSettings() {
	viper.SetConfigType("yaml")
	configReader, err := os.Open("/etc/geniuslurker.yaml")
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	viper.ReadConfig(configReader)
}
