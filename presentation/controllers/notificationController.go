package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
)

type NotificationController struct {
	notificationService services.NotificationServiceInterface
	logger              *tools.LoggerStruct
}

func NewNotificationController(s services.NotificationServiceInterface, l *tools.LoggerStruct) *NotificationController {
	return &NotificationController{notificationService: s, logger: l}
}

func (controller *NotificationController) SendEmails(c *gin.Context) {
	err := controller.notificationService.SendEmails()
	if err != nil {

		msg := controller.logger.LogError("couldn`t send messages")
		tools.Publish(context.Background(), msg)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		msg := controller.logger.LogInfo("successfully send emails")
		tools.Publish(context.Background(), msg)
		c.String(http.StatusOK, "email відправлено")
	}
}
