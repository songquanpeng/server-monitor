package main

import (
	"github.com/gin-gonic/gin"
)

func setRouter(router *gin.Engine) {
	router.GET("/", getIndex)
	router.POST("/gpu", postGpu)
}
