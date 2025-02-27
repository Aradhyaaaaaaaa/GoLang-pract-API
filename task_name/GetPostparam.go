package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong for gin",
			"new_var": "hello",
		})
	})

	router.GET("/me/:id/:newId", func(c *gin.Context){//:this indicate rout where : ad id is variable where param is set
		var id=c.Param("id")
		var newId=c.Param("newId")
		 
		c.JSON(http.StatusOK, gin.H{
			"user_id":id,
			"user_newId":newId,
		})
	})
	router.POST("/me", func(c *gin.Context) {
		type MeRequest struct{
			Email string `json:"EMAILID"`//this is to show how it will look
			Password string `json:"PASSWORD"`//key is password and to map we use json
		}
		var meRequest MeRequest
		c.BindJSON(&meRequest)
		c.JSON(http.StatusOK, gin.H{
			"EMAILID":meRequest.Email,
			"PASSWORD":meRequest.Password,
		})
	})

	router.Run(":8080")
}
