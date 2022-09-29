package rateProviders

type ChainInterface interface {
	GetRateInChain() (int64, error)
	SetNext(ChainInterface)
}
