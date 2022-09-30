package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/presentation/controllers"
)

type Handler struct {
	rateControllerProxy    *controllers.RateControllerProxy
	subscriptionController *controllers.SubscriptionController
	notificationController *controllers.NotificationController
}

func New(r *controllers.RateControllerProxy, s *controllers.SubscriptionController, n *controllers.NotificationController) *Handler {
	return &Handler{
		rateControllerProxy:    r,
		subscriptionController: s,
		notificationController: n,
	}
}

func (h *Handler) CreateRoute() {
	router := gin.Default()

	router.GET("/api/rate", h.rateControllerProxy.Get)
	router.POST("/api/subscribe", h.subscriptionController.SaveEmail)
	router.GET("/api/sendEmails", h.notificationController.SendEmails)

	router.Run()
}
