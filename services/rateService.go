package services

import (
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/services/rateFactory"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"time"
)

type RateService struct {
}

const (
	backupFlag = "coinbase"
)

func (*RateService) GetRate() (int64, error) {

	cfg := config.Get()
	flag := cfg.RateFlag

	method, err := rateFactory.GetSomeRate(flag)

	if err != nil {
		return 0, err
	}

	rate, err := method.GetRateFromProvider()
	if err != nil {
		rate, err = callBackup(backupFlag)
	}

	cache := tools.NewCache(5*time.Minute, 6*time.Minute)
	cache.Set("BtcToUAHrate", rate, 5*time.Minute)
	return rate, err

}

func callBackup(newFlag string) (int64, error) {
	method, err := rateFactory.GetSomeRate(newFlag)
	if err != nil {
		return 0, err
	}
	rate, err := method.GetRateFromProvider()
	if err != nil {
		return 0, err
	}

	return rate, err

}
