package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/controllers"
	"github.com/vsnnkv/btcApplicationGo/models"
	"github.com/vsnnkv/btcApplicationGo/repository"
	"github.com/vsnnkv/btcApplicationGo/services"
)

//
//type routes struct {
//	rateController *controllers.RateController
//	subscriptionController *controllers.SubscriptionController
//}
//
//func Newroutes(r *controllers.RateController, s *controllers.SubscriptionController) *routes{
//	return &routes{
//		rateController: r,
//		subscriptionController: s,
//	}
//}

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
