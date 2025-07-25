package config

type ServerConfig struct {
	AppName                string `mapstructure:"app_name"`
	Port                   string `mapstructure:"port"`
	SecretKey              string `mapstructure:"secret_key"`
	FrontendURL            string `mapstructure:"frontend_url"`
	BaseURL                string `mapstructure:"base_url"`
	StorageProvider        string `mapstructure:"storage_provider"`
	Debug                  bool   `mapstructure:"debug"`
	TokenExpiredDay        int    `mapstructure:"token_expired_day"`
	RefreshTokenExpiredDay int    `mapstructure:"refresh_token_expired_day"`
	SimulateBroadcast      bool   `mapstructure:"simulate_broadcast"`
	SkipVerify             bool   `mapstructure:"skip_verify"`
}
