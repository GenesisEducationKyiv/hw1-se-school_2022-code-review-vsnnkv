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
	logger  tools.LoggerInterface
}

func NewRateController(s services.RateServiceInterface, c *tools.Cache, l tools.LoggerInterface) *RateController {
	return &RateController{service: s,
		cache:  *c,
		logger: l}
}

func (controller *RateController) Get(c *gin.Context) {

	response, err := controller.service.GetRate()
	controller.cache.Set("BtcToUAHrate", response, 5*time.Minute)
	if err == nil {
		controller.logger.LogInfo("successfully return rate and save to cache")
		c.JSON(http.StatusOK, response)
	} else {
		controller.logger.LogError("failed to get rate")
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
}
