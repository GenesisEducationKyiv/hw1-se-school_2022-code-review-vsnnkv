package rateFactory

import (
	"encoding/json"
	"fmt"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

type coinbaseRate struct {
	Rate
}

func newCoinbaseRate() IRate {
	rate, err := getCoinGekoRateBtcToUah()
	return &Rate{rateBtcToUah: rate,
		err: err}
}

func getCoinbaseRateBtcToUah() (int64, error) {
	cfg := config.Get()

	resp, err := http.Get(cfg.CoinGekoURL)

	if err != nil {
		return 0, err
	}

	var cryptoRate models.CoinbaseResponseDTO

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if err = json.Unmarshal(body, &cryptoRate); err != nil {
		return 0, err
	}

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)

	rate := trimStringFromDot(cryptoRate.Data.Uah)

	i, err := strconv.ParseInt(rate, 10, 64)

	if err != nil {
		return 0, err
	}

	return i, nil
}
