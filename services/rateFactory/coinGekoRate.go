package rateFactory

import (
	"encoding/json"
	"fmt"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/models"
	"io/ioutil"
	"net/http"
)

type coinGekoRate struct {
	Rate
}

func newCoinGekoRate() IRate {
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
	//if err := json.NewDecoder(resp.Body).Decode(&cryptoRate); err != nil {
	//	return 0, err
	//}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if err = json.Unmarshal(body, &cryptoRate); err != nil {
		return 0, err
	}

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)

	return cryptoRate.Bitkoin.Uah, nil
}
