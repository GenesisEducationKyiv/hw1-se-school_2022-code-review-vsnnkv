package controllers

import "github.com/gin-gonic/gin"

func (controller *NotificationController) SendEmails(c *gin.Context) {
	code, message := controller.notificationService.SendEmails()

	c.String(code, message)
}
