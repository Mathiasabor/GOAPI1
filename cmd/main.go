package main

import (
	"fmt" 
	"github.com/gin-gonic/gin"

)

func main(){
	fmt.Println("le nouveau api commence")
	router:=gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"hello world",
		})
	})
	router.Run(":5000")
}