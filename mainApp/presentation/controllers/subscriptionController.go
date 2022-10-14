package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
)

type SubscriptionController struct {
	subscriptionService services.SubscriptionServiceInterface
	logger              tools.LoggerInterface
}

func NewSubscriptionController(s services.SubscriptionServiceInterface, l tools.LoggerInterface) *SubscriptionController {
	return &SubscriptionController{subscriptionService: s, logger: l}
}

type Email struct {
	Email string `form:"email" binding:"required"`
}

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var passedParam Email
	if err := c.BindJSON(&passedParam); err != nil {
		controller.logger.LogError(err.Error())
		return
	}

	err := controller.subscriptionService.SaveEmail(passedParam.Email)
	if err != nil {
		controller.logger.LogError("failed to add email")
		c.String(http.StatusInternalServerError, err.Error())
	} else {

		err = controller.subscriptionService.CreateCustomer(passedParam.Email)
		controller.logger.LogInfo("successfully added email")
		c.String(http.StatusOK, "Email додано")
	}
}
