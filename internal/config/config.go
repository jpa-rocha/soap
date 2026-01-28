// Package config defines and populates the config struct
package config

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	HTTPTimeout = 5
)

type Config struct {
	Logger   Logger
	Settings Settings
}

type Logger struct {
	Level string
}

type Settings struct {
	MTU     int
	Address string
}

//nolint:gochecknoglobals
var Opt Config

const value = "value"

func InitConfig(parentContext context.Context) {
	viper.SetConfigName("soap")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/soap/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("SOAP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		slog.ErrorContext(parentContext, "fatal error config file", slog.String(value, err.Error()))
		os.Exit(1)
	}

	Opt = Config{
		Logger{
			Level: viper.GetString("logger.level"),
		},
		Settings{
			MTU:     viper.GetInt("settings.mtu"),
			Address: viper.GetString("settings.address"),
		},
	}

	defaults()
}

func defaults() {
	viper.SetDefault("logger.level", "INFO")
}
