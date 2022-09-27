package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/models"
	"github.com/vsnnkv/btcApplicationGo/services"
)

type SubscriptionController struct {
	subscriptionService services.SubscriptionServiceInterface
}

func NewSubscriptionController(s services.SubscriptionServiceInterface) *SubscriptionController {
	return &SubscriptionController{subscriptionService: s}
}

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var email models.Email
	if err := c.BindJSON(&email); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}
	code, message := controller.subscriptionService.SaveEmail(email.Email)

	c.String(code, message)
}
