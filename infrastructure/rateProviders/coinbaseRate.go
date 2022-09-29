package rateProviders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type coinbaseRate struct {
	Rate
}

func newCoinbaseRate() RateInterface {
	rate, err := getCoinbaseRateBtcToUah()
	return &Rate{rateBtcToUah: rate,
		err: err}
}

func getCoinbaseRateBtcToUah() (int64, error) {
	var cfg rateConfig
	cfg.getConf()

	resp, err := http.Get(cfg.CoinbaseUrl)

	if err != nil {
		return 0, err
	}

	var cryptoRate coinbaseResponseDTO

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
