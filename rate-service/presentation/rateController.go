package routes

import (
	"github.com/vsnnkv/btcApplicationGo/rate-service/services"
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

	var model BtcToUahResponse

	response, err := controller.service.GetRate()

	model.Uah = response
	if err == nil {
		c.JSON(http.StatusOK, model)
	} else {
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
}
