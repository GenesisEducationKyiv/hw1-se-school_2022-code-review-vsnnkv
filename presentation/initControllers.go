package routes

import (
	"github.com/vsnnkv/btcApplicationGo/infrastructure/rateProviders"
	"github.com/vsnnkv/btcApplicationGo/infrastructure/repository"
	"github.com/vsnnkv/btcApplicationGo/presentation/controllers"
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"time"
)

func InitHandler() {
	cache := tools.NewCache(5*time.Minute, 6*time.Minute)
	logger := tools.NewLogger()
	emailsFile := repository.EmailRepository{}
	rateProvider := rateProviders.RateProvider{}

	fileService := services.NewEmailService(&emailsFile)
	subscriptionService := services.NewSubscriptionService(*fileService, logger)
	rateService := services.NewRateService(&rateProvider)
	notificationService := services.NewNotificationService(*rateService, *fileService, logger)

	rateController := controllers.NewRateController(rateService, cache, logger)
	rateControllerProxy := controllers.NewRateControllerProxy(rateController, logger)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService, logger)
	notificationController := controllers.NewNotificationController(notificationService, logger)

	handler := New(rateControllerProxy, subscriptionController, notificationController)
	handler.CreateRoute()
}
