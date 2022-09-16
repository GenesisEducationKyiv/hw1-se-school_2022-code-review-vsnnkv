package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/controllers"
	"github.com/vsnnkv/btcApplicationGo/models"
	"github.com/vsnnkv/btcApplicationGo/repository"
	"github.com/vsnnkv/btcApplicationGo/services"
)

func getRate(c *gin.Context) {

	rateController := controllers.NewRateController(&services.RateService{})
	rateController.Get(c)
}

func subscribe(c *gin.Context) {
	emailsFile := repository.EmailFile{}
	fileService := services.NewFileService(&emailsFile)
	subscriptionService := services.NewSubscriptionService(*fileService)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)

	var email models.Email
	if err := c.BindJSON(&email); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}
	subscriptionController.SaveEmail(email.Email, c)

}

func sendEmails(c *gin.Context) {
	emailsFile := repository.EmailFile{}
	fileService := services.NewFileService(&emailsFile)
	notificationService := services.NewNotificationService(services.RateService{}, *fileService)
	notificationController := controllers.NewNotificationController(notificationService)

	notificationController.SendEmails(c)
}
