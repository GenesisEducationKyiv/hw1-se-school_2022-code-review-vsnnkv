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

	emailsFile := repository.EmailRepository{}
	rateProvider := rateProviders.RateProvider{}
	dtmDB := repository.DtmDb{}

	fileService := services.NewEmailService(&emailsFile)
	subscriptionService := services.NewSubscriptionService(*fileService, &dtmDB)
	notificationService := services.NewNotificationService(services.RateService{}, *fileService)
	rateService := services.NewRateService(&rateProvider)
	dtmService := services.NewDtmService(&dtmDB)

	dtmController := controllers.NewDTMController(dtmService)
	rateController := controllers.NewRateController(rateService, cache)
	rateControllerProxy := controllers.NewRateControllerProxy(rateController)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)
	notificationController := controllers.NewNotificationController(notificationService)

	handler := New(rateControllerProxy, subscriptionController, notificationController, dtmController)
	handler.CreateRoute()
}