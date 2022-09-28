package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RateServiceInterface interface {
	GetRate() (int64, error)
}

type RateService struct {
}

func NewRateService() *RateService {
	return &RateService{}
}

const (
	localRateUrl = "http://localhost:8080/api/rate"
)

func (rateService *RateService) GetRate() (int64, error) {

	resp, err := http.Get(localRateUrl)

	if err != nil {
		return 0, err
	}

	var cryptoRate BtcToUahResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if err = json.Unmarshal(body, &cryptoRate); err != nil {
		return 0, err
	}

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)

	rate := cryptoRate.Uah

	return rate, nil
}
