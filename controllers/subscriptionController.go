package controllers

import (
	"github.com/gin-gonic/gin"
)

func (controller *SubscriptionController) SaveEmail(email string, c *gin.Context) {
	code, message := controller.subscriptionService.SaveEmail(email)

	c.String(code, message)
}
