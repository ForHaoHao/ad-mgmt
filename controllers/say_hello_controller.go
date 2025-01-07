package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	fmt.Printf("Hello World")
	c.JSON(200, gin.H{
		"message": "Fuck World",
	})
}
