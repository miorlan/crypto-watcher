package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     "db",
		DBPort:     "5432",
		DBUser:     "my_user",
		DBPassword: "password",
		DBName:     "postgres",
	}
}
