package config

type EmailConfig struct {
	Server   string `mapstructure:"server"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
	Tls      bool   `mapstructure:"tls"`
	UseAPI   bool   `mapstructure:"use_api"`
}

type EmailApiConfig struct {
	ApiKey    string `mapstructure:"api_key"`
	ApiSecret string `mapstructure:"api_secret"`
	Domain    string `mapstructure:"domain"`
	From      string `mapstructure:"from"`
}
