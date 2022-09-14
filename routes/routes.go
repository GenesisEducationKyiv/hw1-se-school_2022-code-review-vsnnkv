package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vsnnkv/btcApplicationGo/controllers"
	"github.com/vsnnkv/btcApplicationGo/services"
)

func getRate(c *gin.Context) {
	// rate, err := services.GetRate()

	cont := controllers.New(&services.RateService{})

	cont.Get(c)
}
