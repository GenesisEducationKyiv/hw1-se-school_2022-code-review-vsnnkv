package controllers

type RateService interface {
	GetRate() (int64, error)
}

type RateController struct {
	service RateService
}

func NewRateController(s RateService) *RateController {
	return &RateController{service: s}
}
