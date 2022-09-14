package controllers

type RateService interface {
	GetRate() (int64, error)
}

type RateController struct {
	service RateService
}

func New(s RateService) *RateController {
	return &RateController{service: s}
}
