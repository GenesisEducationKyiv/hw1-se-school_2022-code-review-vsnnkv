package rateFactory

type IRate interface {
	GetRateFromProvider() (int64, error)
}
