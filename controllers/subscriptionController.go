package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *SubscriptionController) SaveEmail(email string, c *gin.Context) {
	code := controller.subscriptionService.SaveEmail(email)

	switch code {

	case 200:
		c.String(http.StatusOK, "Email додано")
	case 400:
		c.String(http.StatusBadRequest, "Email вже було додано")
	case 409:
		c.String(http.StatusConflict, "email not valid")
	default:
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
	//if code == 500 {
	//	c.String(http.StatusBadRequest, "Помилка сервера")
	//} else if code == 200 {
	//	c.String(http.StatusOK, "Email додано")
	//} else {
	//	c.String(http.StatusBadRequest, "Email вже було додано")
	//}
}
