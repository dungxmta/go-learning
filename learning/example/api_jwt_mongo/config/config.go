package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	// name of config file (without extension)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// call multiple times to add many search paths
	// optionally look for config in the working directory
	// viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Find and read the config file
	err := viper.ReadInConfig()
	// Handle errors reading the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func GetStr(key string) string {
	return viper.GetString(key)
}
