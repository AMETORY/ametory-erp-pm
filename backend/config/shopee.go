package config

type ShopeeConfig struct {
	APISecret   string `mapstructure:"api_secret"`
	PartnerID   string `mapstructure:"partner_id"`
	Host        string `mapstructure:"host"`
	RedireclURL string `mapstructure:"redirect_url"`
}
