package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Mode                     string `mapstructure:"MODE"`
	AccessSecret             string `mapstructure:"ACCESS_SECRET"`
	ElasticHost              string `mapstructure:"ELASTIC_HOST"`
	ElasticPassword          string `mapstructure:"ELASTIC_PASSWORD"`
	ElasticCrtPath           string `mapstructure:"ELASTIC_CRT_PATH"`
	ProjectServiceRpcAddress string `mapstructure:"PROJECT_SERVICE_RPC_ADDR" required:"true"`
}

func NewConfig() *Config {
	var cfg Config
	viper.AutomaticEnv()
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDir := filepath.Dir(executablePath)
	viper.SetConfigFile(filepath.Join(executableDir, ".env"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
