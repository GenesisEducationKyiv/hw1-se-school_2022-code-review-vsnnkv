package rateFactory

import "errors"

func GetSomeRate(flag string) (IRate, error) {
	switch flag {
	case "coinGeko":
		return newCoinGekoRate(), nil

	case "binance":
		return newBinanceRate(), nil

	case "coinbase":
		return newCoinbaseRate(), nil

	default:
		err := errors.New("Передано некоректний флаг")
		return nil, err
	}
}
