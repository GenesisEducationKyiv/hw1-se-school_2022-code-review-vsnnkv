package routes

import (
	"github.com/gin-gonic/gin"
	controllers2 "github.com/vsnnkv/btcApplicationGo/presentation/controllers"
)

type Handler struct {
	rateController         *controllers2.RateController
	subscriptionController *controllers2.SubscriptionController
	notificationController *controllers2.NotificationController
}

func New(r *controllers2.RateController, s *controllers2.SubscriptionController, n *controllers2.NotificationController) *Handler {
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
