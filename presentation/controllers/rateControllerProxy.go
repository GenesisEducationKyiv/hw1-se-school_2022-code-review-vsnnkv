package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/http"
)

type RateControllerInterface interface {
	Get(c *gin.Context)
}

type RateControllerProxy struct {
	rateController *RateController
	logger         *tools.Logger
}

func NewRateControllerProxy(r *RateController, l *tools.Logger) *RateControllerProxy {
	return &RateControllerProxy{rateController: r, logger: l}
}

func (r *RateControllerProxy) Get(c *gin.Context) {
	rate, isExist := r.rateController.cache.Get("BtcToUAHrate")
	r.logger.LogInfo("checking is rate in cache")
	rateString := fmt.Sprint(rate)
	if isExist {
		r.logger.LogInfo("find and return rate from cache")
		c.String(http.StatusOK, rateString)
	} else {
		r.logger.LogInfo("called rate Service")
		r.rateController.Get(c)
	}
}
