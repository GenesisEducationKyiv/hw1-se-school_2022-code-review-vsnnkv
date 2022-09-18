package rateFactory

import "errors"

func GetSomeRate(flag string) (IRate, error) {
	switch flag {
	case "coinGeko":
		return NewCoinGekoRate(), nil

	case "binance":
		return NewBinanceRate(), nil

	default:
		err := errors.New("Передано некоректний флаг")
		return nil, err
	}
}
