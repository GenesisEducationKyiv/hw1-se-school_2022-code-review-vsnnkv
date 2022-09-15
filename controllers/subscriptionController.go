package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *SubscriptionController) SaveEmail(email string, c *gin.Context) {
	controller.subscriptionService.SaveEmail(email)

	c.String(http.StatusOK, "Email додано")

}
