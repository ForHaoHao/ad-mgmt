package controllers

import (
	"ADMgmtSystem/database"
	"ADMgmtSystem/library"
	"ADMgmtSystem/models/db"
	"ADMgmtSystem/models/vo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	userID, err := library.GenerateRandom(255)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &db.User{ID: userID, Account: AddUserVO.Account, Activated: AddUserVO.Activated, Role: AddUserVO.Role}
	userMeta := &db.UsersMeta{UsersID: userID, Name: AddUserVO.Name, Email: AddUserVO.Email, Avatar: nil, SendEmail: AddUserVO.SendEmail}

	err = database.PgConn.Transaction(func(db *gorm.DB) error {
		if err := user.InsertUser(db); err != nil {
			return err
		}

		if err := userMeta.InsertUsersMeta(db); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"data": "Ok",
	})
}
