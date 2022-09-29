package rateProviders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type coinGekoRate struct {
	Rate
}

func newCoinGekoRate() RateInterface {
	rate, err := getCoinGekoRateBtcToUah()
	return &Rate{rateBtcToUah: rate,
		err: err}
}

func getCoinGekoRateBtcToUah() (int64, error) {
	var cfg rateConfig
	cfg.getConf()

	resp, err := http.Get(cfg.CoingekoUrl)

	if err != nil {
		return 0, err
	}

	var cryptoRate coinGekoResponseDTO

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
