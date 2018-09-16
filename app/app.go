package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Db db
}

type db struct {
	URL string
}

func readConfig() config {
	viper.SetConfigName("base")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	var config config
	viper.Unmarshal(&config)
	return config
}

// PrintConfig prints current config
func PrintConfig() {
	readConfig()
	fmt.Printf("%+v\n", readConfig())
}
