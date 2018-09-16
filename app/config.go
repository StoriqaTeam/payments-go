package app

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config of the app
type Config struct {
	Db Db
}

// Db is a database parameters in config
type Db struct {
	URL string
}

// ReadConfig reads config from config/base.toml.
// It returns error if either file doesn't exist or file structure doesn't match
// expected config struct
func ReadConfig() (Config, error) {
	v := viper.New()
	v.SetConfigName("base")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		return Config{}, errors.Wrap(err, "Reading config file")
	}
	var config Config
	err = v.UnmarshalExact(&config)
	if err != nil {
		return Config{}, errors.Wrap(err, "Parsing config file")
	}
	return config, nil
}
