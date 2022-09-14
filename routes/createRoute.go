package routes

import "github.com/gin-gonic/gin"

func CreateRoute() {
	router := gin.Default()

	router.GET("/api/rate", getRate)

	router.Run()
}
