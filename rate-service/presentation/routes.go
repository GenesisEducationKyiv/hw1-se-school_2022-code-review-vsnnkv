package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/rate-service/config"
)

type Handler struct {
	rateController *RateController
}

func New(r *RateController) *Handler {
	return &Handler{
		rateController: r,
	}
}

func (h *Handler) CreateRoute() {
	router := gin.Default()

	cfg := config.Get()

	router.GET("/api/rate", h.rateController.Get)

	router.Run(cfg.ServerPort)
}
