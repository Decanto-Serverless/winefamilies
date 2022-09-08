package main

import (
	"net/http"

	"winefamilies/env"
	"winefamilies/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET("/", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.GET(baseUrl+"/winefamilies", handlers.GetWinefamilies)
	r.GET(baseUrl+"/winefamilyById", handlers.GetWinefamily)

	r.Run(env.GetInstance().Port)
}
