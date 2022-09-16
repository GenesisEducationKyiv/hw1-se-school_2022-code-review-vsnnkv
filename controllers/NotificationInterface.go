package controllers

type NotificationService interface {
	SendEmails() (int, string)
}

type NotificationController struct {
	notificationService NotificationService
}

func NewNotificationController(n NotificationService) *NotificationController {
	return &NotificationController{n}
}
