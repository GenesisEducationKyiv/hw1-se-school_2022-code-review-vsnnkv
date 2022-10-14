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
	dtmDB := repository.DtmDb{}

	fileService := services.NewEmailService(&emailsFile)
	subscriptionService := services.NewSubscriptionService(*fileService, &dtmDB, logger)
	notificationService := services.NewNotificationService(services.RateService{}, *fileService, logger)
	rateService := services.NewRateService(&rateProvider)
	dtmService := services.NewDtmService(&dtmDB)

	dtmController := controllers.NewDTMController(dtmService)
	rateController := controllers.NewRateController(rateService, cache, logger)
	rateControllerProxy := controllers.NewRateControllerProxy(rateController, logger)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService, logger)
	notificationController := controllers.NewNotificationController(notificationService, logger)

	handler := New(rateControllerProxy, subscriptionController, notificationController, dtmController)
	handler.CreateRoute()
}
