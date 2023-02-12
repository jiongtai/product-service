package config

import "github.com/spf13/viper"

type Config struct {
	Server   `yaml:"Server"`
	DataBase `yaml:"Database"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	vp := viper.New()
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = vp.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
