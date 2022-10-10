package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"net/http"
)

type NotificationController struct {
	notificationService services.NotificationServiceInterface
}

func NewNotificationController(n services.NotificationServiceInterface) *NotificationController {
	return &NotificationController{n}
}

func (controller *NotificationController) SendEmails(c *gin.Context) {
	err := controller.notificationService.SendEmails()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, "email відправлено")
	}
}
