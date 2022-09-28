package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	CoinGekoURL string
	BinanceUrl  string
	CoinbaseUrl string
	RateFlag    string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		loadEnv()
		cfg = Config{
			ServerPort:  os.Getenv(ServerPort),
			CoinGekoURL: os.Getenv(CoinGekoURL),
			BinanceUrl:  os.Getenv(BinanceUrl),
			CoinbaseUrl: os.Getenv(CoinbaseUrl),
			RateFlag:    os.Getenv(RateFlag),
		}
	})
	return &cfg
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("./rate-service/.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
