package config

import (
	"github.com/davidcolman89/config"
)

type Config struct {
	HttpServer string
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbBbdd string
	CsvDelimiter rune
}

func NewConfig(prefix string) Config {
	return initVipper(prefix)
}

func initVipper(prefix string) Config {
	c := Config{}
	config.AddConfigPath("/etc/config")
	config.SetConfig(prefix, &c, "config")
	return c
}


