package config

import "github.com/spf13/viper"

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
	KEY   string `mapstructure:"KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath(".pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigFile("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
