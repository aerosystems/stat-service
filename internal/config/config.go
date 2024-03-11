package config

type Config struct {
	Mode            string `mapstructure:"MODE"`
	AccessSecret    string `mapstructure:"ACCESS_SECRET"`
	ElasticHost     string `mapstructure:"ELASTIC_HOST"`
	ElasticPassword string `mapstructure:"ELASTIC_PASSWORD"`
}
