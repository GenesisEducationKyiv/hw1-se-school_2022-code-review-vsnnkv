package services

import (
	"encoding/json"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/models"
	"net/http"
)

type RateService struct {
}

func (i *RateService) GetRate() (int64, error) {
	cfg := config.Get()

	resp, err := http.Get(cfg.CoinGekoURL)

	if err != nil {
		return 0, err
	}

	var cryptoRate models.CoinGekoResponceDTO
	if err := json.NewDecoder(resp.Body).Decode(&cryptoRate); err != nil {
		return 0, err
	}

	return cryptoRate.Bitkoin.Uah, nil
}
