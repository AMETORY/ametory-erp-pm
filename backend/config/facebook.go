package config

type FacebookConfig struct {
	AppID               string `mapstructure:"app_id"`
	AppSecret           string `mapstructure:"app_secret"`
	RedirectURL         string `mapstructure:"redirect_url"`
	FacebookVerifyToken string `mapstructure:"facebook_verify_token"`
	AppIGID             string `mapstructure:"app_ig_id"`
	AppIGSecret         string `mapstructure:"app_ig_secret"`
	IGRedirectURL       string `mapstructure:"ig_redirect_url"`
	BaseURL             string `mapstructure:"base_url"`
}
