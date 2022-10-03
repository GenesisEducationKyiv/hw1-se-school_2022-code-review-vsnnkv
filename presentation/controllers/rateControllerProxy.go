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
}

func NewRateControllerProxy(r *RateController) *RateControllerProxy {
	return &RateControllerProxy{rateController: r}
}

func (r *RateControllerProxy) Get(c *gin.Context) {
	rate, isExist := r.rateController.cache.Get("BtcToUAHrate")
	tools.Log.Info("test info")
	tools.PublishTest(context.Background(), "message")

	rateString := fmt.Sprint(rate)
	if isExist {
		c.String(http.StatusOK, rateString)
	} else {
		tools.Log.Error("test as error")
		r.rateController.Get(c)
	}
}
