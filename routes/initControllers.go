package routes

import (
	"github.com/vsnnkv/btcApplicationGo/controllers"
	"github.com/vsnnkv/btcApplicationGo/repository"
	"github.com/vsnnkv/btcApplicationGo/services"
)

func InitHandler() {
	emailsFile := repository.EmailFile{}

	fileService := services.NewFileService(&emailsFile)
	subscriptionService := services.NewSubscriptionService(*fileService)
	notificationService := services.NewNotificationService(services.RateService{}, *fileService)

	rateController := controllers.NewRateController(&services.RateService{})
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)
	notificationController := controllers.NewNotificationController(notificationService)

	handler := New(rateController, subscriptionController, notificationController)
	handler.CreateRoute()
}
