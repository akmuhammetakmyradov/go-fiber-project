package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configs struct {
	Listen   Listen     `mapstructure:"listen"`
	Postgres DbPostgres `mapstructure:"postgres"`
	Redis    Redis      `mapstructure:"redis"`
	JWT      JWT        `mapstructure:"jwt"`
}

type (
	Listen struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	DbPostgres struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DbName   string `mapstructure:"db_name"`
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Sslmode  string `mapstructure:"ssl_mode"`
	}

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	}

	JWT struct {
		AccessSecret   string `mapstructure:"access_secret"`
		AccessTokenExp int64  `mapstructure:"access_exp"`
	}
)

func LoadConfiguration() (*Configs, error) {

	// TODO path config
	pathConfig := "./config.yml"
	// for docker
	// pathConfig := "/app/config.yml"

	var config Configs
	viper.SetConfigFile(pathConfig)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err in load config: %v", err)
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		fmt.Printf("err in load config Unmarshal: %v", err)
		return nil, err
	}

	return &config, nil
}
