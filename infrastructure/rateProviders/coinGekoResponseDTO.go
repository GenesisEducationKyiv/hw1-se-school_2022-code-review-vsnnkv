package rateProviders

type coinGekoResponseDTO struct {
	Bitkoin bitcoinResponseDTO `json:"bitcoin"`
}

type bitcoinResponseDTO struct {
	Uah int64 `json:"uah"`
}
