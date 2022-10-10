package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/services"
	"net/http"
)

type SubscriptionController struct {
	subscriptionService services.SubscriptionServiceInterface
	dtm                 DTMController
}

func NewSubscriptionController(s services.SubscriptionServiceInterface, d *DTMController) *SubscriptionController {
	return &SubscriptionController{subscriptionService: s,
		dtm: *d}
}

type Email struct {
	Email string `form:"email" binding:"required"`
}

func (controller *SubscriptionController) SaveEmail(c *gin.Context) {
	var passedParam Email
	if err := c.BindJSON(&passedParam); err != nil {
		fmt.Printf("failed  %s\n", err.Error())
		return
	}

	err := controller.subscriptionService.SaveEmail(passedParam.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {

		createdOrderRequest := orderRequest{}
		c.BindJSON(&createdOrderRequest)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		err := controller.dtm.createOrder(&createdOrderRequest)
		fmt.Println(err)
		//here should pass to DTS
		c.String(http.StatusOK, "Email додано")
	}

}

type orderRequest struct {
	IdCustomer uint   `json:"idCustomer"`
	Email      string `json:"email"`
}
