package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Email    EmailConfig    `mapstructure:"email"`
	EmailApi EmailApiConfig `mapstructure:"email_api"`
	Google   GoogleConfig   `mapstructure:"google"`
	Whatsapp WhatsappConfig `mapstructure:"whatsapp"`
	Facebook FacebookConfig `mapstructure:"facebook"`
	Tiktok   TiktokConfig   `mapstructure:"tiktok"`
	Shopee   ShopeeConfig   `mapstructure:"shopee"`
	Lazada   LazadaConfig   `mapstructure:"lazada"`
}

var App = &Config{}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config") // Nama file config (tanpa ekstensi)
	viper.SetConfigType("yaml")   // Format file config (yaml, json, toml, dll.)
	viper.AddConfigPath(".")      // Path ke file config (direktori saat ini)
	viper.AutomaticEnv()          // Baca environment variables

	// Baca file config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("gagal membaca file config: %v", err)
	}

	// Unmarshal config ke struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("gagal unmarshal config: %v", err)
	}

	App = &cfg
	return &cfg, nil
}
