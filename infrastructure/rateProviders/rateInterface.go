package rateProviders

type RateInterface interface {
	GetRateFromProvider() (int64, error)
}
