package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Storage    Storage    `mapstructure:"storage"`
	HttpServer HttpServer `mapstructure:"http_server"`
	Cache      Cache      `mapstructure:"cache"`
}

type Storage struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type HttpServer struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

type Cache struct {
	Host     string        `mapstructure:"host"`
	Port     int           `mapstructure:"port"`
	User     string        `mapstructure:"user"`
	Password string        `mapstructure:"password"`
	Database int           `mapstructure:"database"`
	Ttl      time.Duration `mapstructure:"ttl"`
}

func Read(configPath string) (Config, error) {
	if configPath == `` {
		return Config{}, errors.New(`config path is empty`)
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error read config file: %s", err)
	}

	var conf Config

	if err := viper.Unmarshal(&conf); err != nil {
		return Config{}, fmt.Errorf("error unmarshal config file: %s", err)
	}

	return conf, nil
}
