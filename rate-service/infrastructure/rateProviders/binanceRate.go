package rateProviders

import (
	"encoding/json"
	"fmt"
	"github.com/vsnnkv/btcApplicationGo/rate-service/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type binanceRate struct {
	Rate
}

func newBinanceRate() IRate {
	rate, err := getBinanceRateBtcToUah()
	return &Rate{rateBtcToUah: rate, err: err}

}

func getBinanceRateBtcToUah() (int64, error) {
	cfg := config.Get()

	resp, err := http.Get(cfg.BinanceUrl)

	if err != nil {
		return 0, err
	}

	var cryptoRate BinanceResponse
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
