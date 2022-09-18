package rateFactory

import (
	"encoding/json"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/models"
	"net/http"
)

type coinGekoRate struct {
	Rate
}

func NewCoinGekoRate() IRate {
	rate, err := getCoinGekoRateBtcToUah()
	return &Rate{rateBtcToUah: rate,
		err: err}
}

func getCoinGekoRateBtcToUah() (int64, error) {
	cfg := config.Get()

	resp, err := http.Get(cfg.CoinGekoURL)

	if err != nil {
		return 0, err
	}

	var cryptoRate models.CoinGekoResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&cryptoRate); err != nil {
		return 0, err
	}

	return cryptoRate.Bitkoin.Uah, nil
}
