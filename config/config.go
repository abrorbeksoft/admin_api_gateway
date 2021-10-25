package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	AuthHost string
	AuthPort string
}

func Load() *Config {
	if err:=godotenv.Load(); err!= nil {
		fmt.Println(err.Error())
	}
	con:=Config{}
	con.AuthHost=cast.ToString(getOrReturnDefault("AUTH_HOST","localhost"))
	con.AuthPort=cast.ToString(getOrReturnDefault("AUTH_PORT","8000"))
	return &con
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_,exists:=os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
