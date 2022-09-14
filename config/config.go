package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerURL   string
	CoinGekoURL string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		loadEnv()
		cfg = Config{
			ServerURL:   os.Getenv(ServerURL),
			CoinGekoURL: os.Getenv(CoinGekoURL),
		}
	})
	return &cfg
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		// in case we test from inner directories;
		// sequence to go to the upper one
		err = godotenv.Load("./../.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
