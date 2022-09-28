package controllers

import (
	"github.com/vsnnkv/btcApplicationGo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RateController struct {
	service services.RateServiceInterface
}

func NewRateController(s services.RateServiceInterface) *RateController {
	return &RateController{service: s}
}

func (controller *RateController) Get(c *gin.Context) {

	response, err := controller.service.GetRate()
	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
}
