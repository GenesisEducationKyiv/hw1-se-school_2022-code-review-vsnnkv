package controllers

import (
	"context"
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
	logger         *tools.LoggerStruct
}

func NewRateControllerProxy(r *RateController, l *tools.LoggerStruct) *RateControllerProxy {
	return &RateControllerProxy{rateController: r, logger: l}
}

func (r *RateControllerProxy) Get(c *gin.Context) {
	rate, isExist := r.rateController.cache.Get("BtcToUAHrate")
	tools.Log.Info("test info")
	tools.Publish(context.Background(), "message")

	rateString := fmt.Sprint(rate)
	if isExist {
		msg := r.logger.LogInfo("find and return rate from cache")
		tools.Publish(context.Background(), msg)
		c.String(http.StatusOK, rateString)
	} else {
		msg := r.logger.LogInfo("called rate Service")
		tools.Publish(context.Background(), msg)
		r.rateController.Get(c)
	}
}
