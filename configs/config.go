package configs

import "github.com/spf13/viper"

type Conf struct {
	RedisAddr      string `mapstructure:"REDIS_ADDR"`
	RedisPassword  string `mapstructure:"REDIS_PASSWORD"`
	RedisDB        int    `mapstructure:"REDIS_DB"`
	LimitReqIp     int    `mapstructure:"LIMIT_REQ_IP"`
	LimitReqToken  int    `mapstructure:"LIMIT_REQ_TOKEN"`
	TimeLimitIp    int    `mapstructure:"TTL_IP"`
	TimeLimitToken int    `mapstructure:"TTL_TOKEN"`
	Port           string `mapstructure:"PORT"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	var err error
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg, err
}
