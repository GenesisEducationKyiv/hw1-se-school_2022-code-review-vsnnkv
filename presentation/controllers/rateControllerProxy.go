package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RateControllerInterface interface {
	Get(c *gin.Context)
}

type RateControllerProxy struct {
	rateController *RateController
}

func NewRateControllerProxy(r *RateController) *RateControllerProxy {
	return &RateControllerProxy{rateController: r}
}

func (r *RateControllerProxy) Get(c *gin.Context) {
	rate, isExist := r.rateController.cache.Get("BtcToUAHrate")

	rateString := fmt.Sprint(rate)
	if isExist {
		c.String(http.StatusOK, rateString)
	} else {
		r.rateController.Get(c)
	}
}
