package controllers

import (
	"ADMgmtSystem/library"
	"ADMgmtSystem/services"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Login
//
//	@Summary		User login function
//	@Description	login
//	@Tags			login
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	ok
//	@Failure		401	{string}	error
//	@Router			/login [get]
func Login(ginContext *gin.Context) {
	//headerAuth := context.GetHeader("Authorization")
	headerAuth := "ZHM3QHE6SrV4SHk0RXQ1"
	if headerAuth == "" {
		ginContext.JSON(401, gin.H{
			"error": "Authorization header is required",
		})
		return
	}
	decodeAuth, err := library.CaesarEncrypted(headerAuth, 0, 4)

	if err != nil {
		library.Log.Error(err)
		ginContext.JSON(401, gin.H{
			"error": "Caesar decode error",
		})
		return
	}

	originCode, err := base64.StdEncoding.DecodeString(decodeAuth)

	if err != nil {
		library.Log.Error(err)
		ginContext.JSON(401, gin.H{
			"error": "base64 decode error",
		})
		return
	}
	userInfo := strings.Split(string(originCode), ":")

	err = services.CheckUser(userInfo[0], userInfo[1])
	if err != nil {
		library.Log.Error(err)

		ginContext.JSON(401, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	ginContext.SetCookie("SystemID", string(originCode), 0, "/", ".haocess.com", false, true)
}
