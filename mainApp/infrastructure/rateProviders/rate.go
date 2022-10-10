package rateProviders

type Rate struct {
	rateBtcToUah int64
	err          error
}

func (r *Rate) GetRateFromProvider() (int64, error) {
	return r.rateBtcToUah, r.err
}
