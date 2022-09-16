package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/models"
)

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var email models.Email
	if err := c.BindJSON(&email); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}
	code, message := controller.subscriptionService.SaveEmail(email.Email)

	c.String(code, message)
}
