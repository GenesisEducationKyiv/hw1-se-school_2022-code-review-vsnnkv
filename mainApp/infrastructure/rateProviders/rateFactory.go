package rateProviders

import "errors"

type RateProviderInterface interface {
	CreateRateMethod(flag string) (RateInterface, error)
}

type RateProvider struct {
}

func (*RateProvider) CreateRateMethod(flag string) (RateInterface, error) {
	switch flag {
	case "coinGeko":
		return newCoinGekoRate(), nil

	case "binance":
		return newBinanceRate(), nil

	case "coinbase":
		return newCoinbaseRate(), nil

	default:
		return nil, errors.New("Передано некоректний флаг")
	}
}
