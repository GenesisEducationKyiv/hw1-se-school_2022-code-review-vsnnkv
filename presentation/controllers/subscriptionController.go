package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"net/http"
)

type SubscriptionController struct {
	subscriptionService services.SubscriptionServiceInterface
}

func NewSubscriptionController(s services.SubscriptionServiceInterface) *SubscriptionController {
	return &SubscriptionController{subscriptionService: s}
}

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var passedParam email
	if err := c.BindJSON(&passedParam); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}

	err := controller.subscriptionService.SaveEmail(passedParam.email)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, "Email додано")
	}

}

type email struct {
	email string `form:"email" binding:"required"`
}
