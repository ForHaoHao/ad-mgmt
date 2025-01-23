package controllers

import (
	"ADMgmtSystem/library"
	"ADMgmtSystem/models/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddUser Insert user to database
//
//	@Summary		AddUser
//	@Description	Insert user to database
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			addUserVO	body		vo.AddUserVO	true	"User data"
//	@Success		200			{string}	OK
//	@Failure		400			{string}	ErrorResponse
//	@Router			/user [post]
func AddUser(ginContext *gin.Context) {
	library.Log.Info("Start insert user info")
	var AddUserVO vo.AddUserVO

	if err := ginContext.ShouldBindJSON(&AddUserVO); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
}
