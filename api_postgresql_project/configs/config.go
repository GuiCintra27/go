package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

func init(){
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error{
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = &config{}
	cfg.API.Port = viper.GetString("api.port")
	cfg.DB.Host = viper.GetString("database.host")
	cfg.DB.User = viper.GetString("database.user")
	cfg.DB.Password = viper.GetString("database.password")
	cfg.DB.Database = viper.GetString("database.database")
	cfg.DB.Port = viper.GetString("database.port")

	return nil
}

func GetDB() DBConfig{
	return cfg.DB
}

func GetAPI() APIConfig{
	return cfg.API
}