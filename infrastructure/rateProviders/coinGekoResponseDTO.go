package rateProviders

type CoinGekoResponseDTO struct {
	Bitkoin bitcoinResponseDTO `json:"bitcoin"`
}

type bitcoinResponseDTO struct {
	Uah int64 `json:"uah"`
}
