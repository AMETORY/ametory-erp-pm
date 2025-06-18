package config

type LazadaConfig struct {
	APIKey      string `mapstructure:"api_key"`
	APISecret   string `mapstructure:"api_secret"`
	Region      string `mapstructure:"region"`
	CallbackURL string `mapstructure:"callback_url"` // Optional, if you need to set a callback URL
}
