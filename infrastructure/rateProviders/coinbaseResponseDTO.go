package rateProviders

type coinbaseResponseDTO struct {
	Data coinResponseDTO `json:"data"`
}

type coinResponseDTO struct {
	Uah string `json:"amount"`
}
