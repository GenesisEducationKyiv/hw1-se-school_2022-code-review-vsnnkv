package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *RateController) Get(c *gin.Context) {

	response, err := controller.service.GetRate()
	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}
}
