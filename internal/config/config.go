package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort   int    `mapstructure:"SERVER_PORT"`
	DbName       string `mapstructure:"DATABASE_NAME"`
	DbHost       string `mapstructure:"DATABASE_HOST"`
	DbPort       int    `mapstructure:"DATABASE_PORT"`
	DbUser       string `mapstructure:"DATABASE_USER"`
	DbPassword   string `mapstructure:"DATABASE_PASSWORD"`
	LdapHost     string `mapstructure:"LDAP_HOST"`
	LdapBaseDn   string `mapstructure:"LDAP_BASE_DN"`
	LdapUser     string `mapstructure:"LDAP_USER"`
	LdapPassword string `mapstructure:"LDAP_PASSWORD"`
}

func InitConfig() (*Config, error) {
	var cfg = Config{}
	// TODO Перенести пароли в env файл
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../../configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("can't load configuration file: %s", err)
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &cfg, err
}
