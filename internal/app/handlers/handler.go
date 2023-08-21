package handlers

import (

	"github.com/gin-gonic/gin"
	"goapi1/internal/app/models"
)

func IndexHandler(c *gin.Context){
	
	c.XML(200, models.Person{
		FirstName :"Mathias",
		LastName :"AHO"} )				
}