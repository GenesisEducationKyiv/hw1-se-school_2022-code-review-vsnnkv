package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	EmailAddress  string
	EmailPassword string
	SMTPHost      string
	SMTPPort      string
	EmailFile     string
	RateFlag      string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		loadEnv()
		cfg = Config{
			ServerPort:    os.Getenv(ServerPort),
			EmailAddress:  os.Getenv(EmailAddress),
			EmailPassword: os.Getenv(EmailPassword),
			SMTPHost:      os.Getenv(SMTPHost),
			SMTPPort:      os.Getenv(SMTPPort),
			EmailFile:     os.Getenv(EmailFile),
			RateFlag:      os.Getenv(RateFlag),
		}
	})
	return &cfg
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("./subscription-service/.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
