package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"net/http"
)

type DTMController struct {
	dtmService services.DTMServiceInterface
}

func NewDTMController(d services.DTMServiceInterface) *DTMController {
	return &DTMController{dtmService: d}
}

func (controller *DTMController) RegisterCustomerStatus(c *gin.Context) {
	registerCustomerRequest := struct {
		IdCustomer uint   `json:"idCustomer"`
		Currency   string `json:"currency"`
		Amount     uint   `json:"amount"`
	}{}
	transactionId := c.Query("gid")

	err := c.BindJSON(&registerCustomerRequest)
	if err != nil {
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}

	err = controller.dtmService.ChangeCustomerStatus(transactionId, registerCustomerRequest.IdCustomer,
		registerCustomerRequest.Currency, registerCustomerRequest.Amount)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, err.Error())
	}

}

func (controller *DTMController) RegisterCustomerStatusCompensate(c *gin.Context) {
	transactionId := c.Query("gid")

	err := controller.dtmService.ChangeCustomerStatusCompensation(transactionId)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, err.Error())
}
