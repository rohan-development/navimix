package config

import "github.com/BurntSushi/toml"

var config Config

func Load() *Config {
	var config Config
	_, err := toml.DecodeFile("navimix.toml", &config)
	if err != nil {
		panic(err)
	}
	return &config
}
