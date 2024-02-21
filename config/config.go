package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
 
)

 

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB      string
}

func Load() Config{
	cfg := Config{}

	err := godotenv.Load()
	if err != nil{
		fmt.Println("Error while loading from godotenv!", err.Error())
		return Config{}
	}

	cfg.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	cfg.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "user"))
	cfg.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "password"))
	cfg.PostgresDB = cast.ToString(GetOrReturnDefault("POSTGRES_DB", "database"))

	return cfg
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{}{
	value := os.Getenv(key)
	if value != ""{
		return value
	}
	return defaultValue
}