package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/controllers"
	"github.com/vsnnkv/btcApplicationGo/models"
	"github.com/vsnnkv/btcApplicationGo/services"
)

func getRate(c *gin.Context) {

	rateController := controllers.NewRateController(&services.RateService{})
	rateController.Get(c)
}

func subscribe(c *gin.Context) {
	emailsFile := services.FileService{}
	subscriptionController := controllers.NewSubscriptionController(services.NewSubscriptionService(emailsFile))

	var email models.Email
	if err := c.BindJSON(&email); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}
	subscriptionController.SaveEmail(email.Email, c)

}
