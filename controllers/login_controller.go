package controllers

import (
	"ADMgmtSystem/library"
	"ADMgmtSystem/services"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	//headerAuth := context.GetHeader("Authorization")
	headerAuth := "ZHM3QHE6SrV4SHk0RXQ1"
	if headerAuth == "" {
		context.JSON(401, gin.H{
			"error": "Authorization header is required",
		})
		return
	}
	decodeString, err := library.CaesarEncrypted(headerAuth, 0, 4)

	if err != nil {
		library.Log.Error(err)
		context.JSON(401, gin.H{
			"error": "Caesar decode error",
		})
		return
	}

	originCode, err := base64.StdEncoding.DecodeString(decodeString)

	if err != nil {
		library.Log.Error(err)
		context.JSON(401, gin.H{
			"error": "base64 decode error",
		})
		return
	}
	userInfo := strings.Split(string(originCode), ":")

	err = services.CheckUser(userInfo[0], userInfo[1])
	if err != nil {
		library.Log.Error(err)
		context.JSON(401, gin.H{
			"error": fmt.Sprintf("%c", err),
		})
		return
	}
	context.JSON(200, gin.H{
		"token": string(originCode),
	})
}
