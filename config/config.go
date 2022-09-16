package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerURL     string
	CoinGekoURL   string
	EmailAddress  string
	EmailPassword string
	SMTPHost      string
	SMTPPort      string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		loadEnv()
		cfg = Config{
			ServerURL:     os.Getenv(ServerURL),
			CoinGekoURL:   os.Getenv(CoinGekoURL),
			EmailAddress:  os.Getenv(EmailAddress),
			EmailPassword: os.Getenv(EmailPassword),
			SMTPHost:      os.Getenv(SMTPHost),
			SMTPPort:      os.Getenv(SMTPPort),
		}
	})
	return &cfg
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("./../.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
