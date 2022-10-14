package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/presentation/controllers"
	"github.com/vsnnkv/btcApplicationGo/tools"
)

type Handler struct {
	rateControllerProxy    *controllers.RateControllerProxy
	subscriptionController *controllers.SubscriptionController
	notificationController *controllers.NotificationController
	dtmController          *controllers.DTMController
}

func New(r *controllers.RateControllerProxy, s *controllers.SubscriptionController, n *controllers.NotificationController,
	d *controllers.DTMController) *Handler {
	return &Handler{
		rateControllerProxy:    r,
		subscriptionController: s,
		notificationController: n,
		dtmController:          d,
	}
}

func (h *Handler) CreateRoute() {
	router := gin.Default()

	router.GET("/api/rate", h.rateControllerProxy.Get)
	router.POST("/api/subscribe", h.subscriptionController.SaveEmail)
	router.GET("/api/sendEmails", h.notificationController.SendEmails)

	router.POST("/register-order", h.dtmController.RegisterCustomerStatus)
	router.POST("/register-order-compensate", h.dtmController.RegisterCustomerStatusCompensate)

	ctx, _ := context.WithCancel(context.Background())
	tools.Start(ctx)

	router.Run()
}
