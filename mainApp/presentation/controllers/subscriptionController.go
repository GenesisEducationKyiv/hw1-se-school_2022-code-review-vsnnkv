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

type Email struct {
	Email string `form:"email" binding:"required"`
}

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var passedParam Email
	if err := c.BindJSON(&passedParam); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}

	err := controller.subscriptionService.SaveEmail(passedParam.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {

		err = controller.subscriptionService.CreateCustomer(passedParam.Email)

		c.String(http.StatusOK, "Email додано")
	}
}
