package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT string `mapstructure:"port"`
	// MongoUri    string `mapstructure:"MONGO_URI"`
	// MongoDBName string `mapstructure:"MONGO_DB_NAME"`
	GeminiApi string `mapstructure:"GEMINI_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return

}
