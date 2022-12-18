package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

func LoadEnvConfig() *Specification {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	a := viper.GetViper().AllKeys()
	for _, env := range a {
		if viper.GetString(env) != "" {
			os.Setenv(strings.ToUpper(env), viper.GetString(env))
		}
	}
	var s Specification
	err = viper.Unmarshal(&s)
	err = envconfig.Process("", &s)
	if err != nil {
		panic(err)
	}
	return &s
}
