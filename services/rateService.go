package services

import (
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/services/rateFactory"
)

type RateService struct {
}

func (*RateService) GetRate() (int64, error) {

	cfg := config.Get()
	flag := cfg.RateFlag

	method, err := rateFactory.GetSomeRate(flag)

	if err != nil {
		return 0, err
	}

	rate, err := method.GetRateFromProvider()

	return rate, err

}
