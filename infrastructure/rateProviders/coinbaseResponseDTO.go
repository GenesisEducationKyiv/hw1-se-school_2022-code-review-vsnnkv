package rateProviders

type CoinbaseResponseDTO struct {
	Data coinResponseDTO `json:"data"`
}

type coinResponseDTO struct {
	Uah string `json:"amount"`
}
