package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
)

type NotificationController struct {
	notificationService services.NotificationServiceInterface
}

func NewNotificationController(n services.NotificationServiceInterface) *NotificationController {
	return &NotificationController{n}
}

func (controller *NotificationController) SendEmails(c *gin.Context) {
	code, message := controller.notificationService.SendEmails()

	c.String(code, message)
}
