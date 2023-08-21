package main

import (
	"fmt" 
	"github.com/gin-gonic/gin"
	"goapi1/internal/app/handlers"

)
//func IndexHandler(c *gin.Context){
//	name :=c.Params.ByName("name")
//	c.JSON(200, gin.H{
//		"Accueil" : "Bonjour"+name,
//	})
//}

func main(){
	fmt.Println("le nouveau api commence")
	router:=gin.Default()
	router.GET("/", handlers.IndexHandler)
	//router.GET("/:name", handlers.IndexHandler)
	router.Run()
}

