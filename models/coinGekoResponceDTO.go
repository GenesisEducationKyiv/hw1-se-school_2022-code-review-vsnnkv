package models

type CoinGekoResponceDTO struct {
	Bitkoin bitcoinResponceDTO `json:"bitcoin"`
}

type bitcoinResponceDTO struct {
	Uah int64 `json:"uah"`
}
