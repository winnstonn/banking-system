package config

type Config struct {
	DBConfig      *DatabaseConfig `mapstructure:"db"`
	ServerAddress string          `mapstructure:"serverAddress"`
}

type DatabaseConfig struct {
	DBHost string `mapstructure:"dbHost"`
	DBPort string `mapstructure:"dbPort"`
	DBUser string `mapstructure:"dbUser"`
	DBPass string `mapstructure:"dbPass"`
	DBName string `mapstructure:"dbName"`
}
