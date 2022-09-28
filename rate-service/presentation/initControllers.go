package routes

import (
	"github.com/vsnnkv/btcApplicationGo/rate-service/infrastructure/rateProviders"
	"github.com/vsnnkv/btcApplicationGo/rate-service/services"
)

func InitHandler() {
	rateProvider := rateProviders.RateProvider{}

	rateService := services.NewRateService(&rateProvider)

	rateController := NewRateController(rateService)

	handler := New(rateController)
	handler.CreateRoute()
}
