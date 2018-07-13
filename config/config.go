package config

import (
	"github.com/spf13/viper"
	"fmt"
)

type config struct {
	Username string
	Password string
}

var Config config

func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.scc")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}
