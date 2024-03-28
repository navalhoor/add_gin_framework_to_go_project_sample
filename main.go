package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//init gin
	InitGin()
}
func InitGin() {
	//Set gin engin mode (debug-release-test)
	gin.SetMode("debug")
	ginEngine := gin.Default()
	//Maximum request size
	ginEngine.MaxMultipartMemory = 100 << 20 // 100 MB
	
	ginEngine.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	//Run Gin on especial port and host
	ginEngine.Run(":8282")
}
