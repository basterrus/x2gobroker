package configs

import (
	"github.com/spf13/viper"
	"os"
)

func Init() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	env := os.Getenv("POSTGRES_PASSWORD")
	viper.SetConfigName(env)
	return viper.MergeInConfig()
}
