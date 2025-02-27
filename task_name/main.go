package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main1(){
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong for gin",
			"new_var": "hello",
		})
	})
	router.Run(":8000")
	}