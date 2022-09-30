package controllers

import (
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RateController struct {
	service services.RateServiceInterface
	cache   tools.Cache
}

func NewRateController(s services.RateServiceInterface, c *tools.Cache) *RateController {
	return &RateController{service: s,
		cache: *c}
}

func (controller *RateController) Get(c *gin.Context) {

	response, err := controller.service.GetRate()
	controller.cache.Set("BtcToUAHrate", response, 5*time.Minute)
	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
}
