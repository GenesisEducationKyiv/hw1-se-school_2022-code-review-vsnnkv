package rateFactory

import (
	"encoding/json"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/models"
	"net/http"
	"strconv"
	"strings"
)

type binanceRate struct {
	Rate
}

func NewBinanceRate() IRate {
	rate, err := getBinanceRateBtcToUah()
	return &Rate{rateBtcToUah: rate, err: err}

}

func getBinanceRateBtcToUah() (int64, error) {
	cfg := config.Get()

	resp, err := http.Get(cfg.BinanceUrl)

	if err != nil {
		return 0, err
	}

	var cryptoRate models.BinanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&cryptoRate); err != nil {
		return 0, err
	}

	rate := trimStringFromDot(cryptoRate.Uah)

	i, err := strconv.ParseInt(rate, 10, 64)

	if err != nil {
		return 0, err
	}

	return i, nil

}

func trimStringFromDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
