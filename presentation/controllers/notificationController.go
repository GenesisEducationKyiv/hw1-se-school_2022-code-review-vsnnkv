package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
)

type NotificationController struct {
	notificationService services.NotificationServiceInterface
	logger              *tools.Logger
}

func NewNotificationController(s services.NotificationServiceInterface, l *tools.Logger) *NotificationController {
	return &NotificationController{notificationService: s, logger: l}
}

func (controller *NotificationController) SendEmails(c *gin.Context) {
	err := controller.notificationService.SendEmails()
	if err != nil {
		controller.logger.LogError("ouldn`t send messages")
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		controller.logger.LogInfo("successfully send emails")
		c.String(http.StatusOK, "email відправлено")
	}
}
