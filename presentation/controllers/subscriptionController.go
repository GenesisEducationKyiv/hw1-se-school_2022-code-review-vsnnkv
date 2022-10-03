package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
)

type SubscriptionController struct {
	subscriptionService services.SubscriptionServiceInterface
	logger              *tools.LoggerStruct
}

func NewSubscriptionController(s services.SubscriptionServiceInterface, l *tools.LoggerStruct) *SubscriptionController {
	return &SubscriptionController{subscriptionService: s, logger: l}
}

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var passedParam email
	if err := c.BindJSON(&passedParam); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}

	err := controller.subscriptionService.SaveEmail(passedParam.email)
	if err != nil {
		msg := controller.logger.LogError("failed to add email")
		tools.Publish(context.Background(), msg)

		c.String(http.StatusInternalServerError, err.Error())
	} else {
		msg := controller.logger.LogInfo("successfully added email")
		tools.Publish(context.Background(), msg)
		c.String(http.StatusOK, "Email додано")
	}

}

type email struct {
	email string `form:"email" binding:"required"`
}
