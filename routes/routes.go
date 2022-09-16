package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/controllers"
)

type Handler struct {
	rateController         *controllers.RateController
	subscriptionController *controllers.SubscriptionController
	notificationController *controllers.NotificationController
}

func New(r *controllers.RateController, s *controllers.SubscriptionController, n *controllers.NotificationController) *Handler {
	return &Handler{
		rateController:         r,
		subscriptionController: s,
		notificationController: n,
	}
}

func (h *Handler) CreateRoute() {
	router := gin.Default()

	router.GET("/api/rate", h.rateController.Get)
	router.POST("/api/subscribe", h.subscriptionController.SaveEmail)
	router.GET("/api/sendEmails", h.notificationController.SendEmails)

	router.Run()
}
