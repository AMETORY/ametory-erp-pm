package config

type TiktokConfig struct {
	AppKey    string `mapstructure:"app_key"`
	AppSecret string `mapstructure:"app_secret"`
	ServiceID string `mapstructure:"service_id"`
}
