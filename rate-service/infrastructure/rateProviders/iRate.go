package rateProviders

type IRate interface {
	GetRateFromProvider() (int64, error)
}
