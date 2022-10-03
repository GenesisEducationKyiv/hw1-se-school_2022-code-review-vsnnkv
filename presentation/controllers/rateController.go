package controllers

import (
	"context"
	"github.com/vsnnkv/btcApplicationGo/services"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RateController struct {
	service services.RateServiceInterface
	cache   tools.Cache
	logger  *tools.LoggerStruct
}

func NewRateController(s services.RateServiceInterface, c *tools.Cache, l *tools.LoggerStruct) *RateController {
	return &RateController{service: s,
		cache:  *c,
		logger: l}
}

func (controller *RateController) Get(c *gin.Context) {

	response, err := controller.service.GetRate()
	controller.cache.Set("BtcToUAHrate", response, 5*time.Minute)
	if err == nil {
		msg := controller.logger.LogInfo("successfully return rate and save to cache")
		tools.Publish(context.Background(), msg)
		c.JSON(http.StatusOK, response)
	} else {
		msg := controller.logger.LogError("failed to get rate")
		tools.Publish(context.Background(), msg)
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
}
