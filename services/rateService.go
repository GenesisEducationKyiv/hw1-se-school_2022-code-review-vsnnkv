package services

import (
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/infrastructure/rateProviders"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"time"
)

type RateServiceInterface interface {
	GetRate() (int64, error)
}

type RateService struct {
	rateProviders rateProviders.RateProviderInterface
}

func NewRateService(r rateProviders.RateProviderInterface) *RateService {
	return &RateService{rateProviders: r}
}

const (
	backupFlag = "coinbase"
)

func (rateService *RateService) GetRate() (int64, error) {

	cfg := config.Get()
	flag := cfg.RateFlag

	method, err := rateService.rateProviders.CreateRateMethod(flag)

	if err != nil {
		return 0, err
	}

	rate, err := method.GetRateFromProvider()

	if err != nil {
		return createAndStartChain()
	}

	cache := tools.NewCache(5*time.Minute, 6*time.Minute)
	cache.Set("BtcToUAHrate", rate, 5*time.Minute)
	return rate, err

}

func createAndStartChain() (int64, error) {
	coingeko := &rateProviders.CoinGekoRate{}

	coinbase := &rateProviders.CoinbaseRate{}
	coinbase.SetNext(coingeko)

	binance := &rateProviders.BinanceRate{}
	binance.SetNext(coinbase)

	return binance.GetRateInChain()
}
