package routes

import (
	"github.com/vsnnkv/btcApplicationGo/infrastructure/rateProviders"
	"github.com/vsnnkv/btcApplicationGo/infrastructure/repository"
	"github.com/vsnnkv/btcApplicationGo/presentation/controllers"
	"github.com/vsnnkv/btcApplicationGo/services"
)

func InitHandler() {
	emailsFile := repository.EmailFile{}
	rateProvider := rateProviders.RateProvider{}

	fileService := services.NewFileService(&emailsFile)
	subscriptionService := services.NewSubscriptionService(*fileService)
	notificationService := services.NewNotificationService(services.RateService{}, *fileService)
	rateService := services.NewRateService(&rateProvider)

	rateController := controllers.NewRateController(rateService)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)
	notificationController := controllers.NewNotificationController(notificationService)

	handler := New(rateController, subscriptionController, notificationController)
	handler.CreateRoute()
}
