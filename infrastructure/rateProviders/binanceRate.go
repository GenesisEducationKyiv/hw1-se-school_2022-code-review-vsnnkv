package rateProviders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type BinanceRate struct {
	Rate
	next ChainInterface
}

func (r *BinanceRate) GetRateInChain() (int64, error) {

	rate, err := getBinanceRateBtcToUah()

	if err != nil {
		return r.next.GetRateInChain()
	} else {
		return rate, err
	}
}

func (r *BinanceRate) SetNext(next ChainInterface) {
	r.next = next
}

func newBinanceRate() RateInterface {
	rate, err := getBinanceRateBtcToUah()
	return &Rate{rateBtcToUah: rate, err: err}

}

func getBinanceRateBtcToUah() (int64, error) {
	var cfg rateConfig
	cfg.getConf()

	resp, err := http.Get(cfg.BinanceUrl)

	if err != nil {
		return 0, err
	}

	var cryptoRate binanceResponse

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
