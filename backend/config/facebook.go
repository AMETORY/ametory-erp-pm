package config

type FacebookConfig struct {
	AppID       string `mapstructure:"app_id"`
	AppSecret   string `mapstructure:"app_secret"`
	RedirectURL string `mapstructure:"redirect_url"`
}
