package db

import (
	"ADMgmtSystem/database"
	"ADMgmtSystem/library"
	"ADMgmtSystem/models"
)

type UsersRepository interface {
	GetUserByAccount(userAccount string) (models.Users, error)
	UpdateUserErrorById(userId string) error
}
type User struct{}

func (u *User) GetUserByAccount(userAccount string) (models.Users, error) {
	var user models.Users

	library.Log.Info("Start get the user by user account")
	library.Log.Debugf("userAccount=%s", userAccount)
	result := database.PgConn.Where("account = ?", userAccount).First(&user)

	if result.Error != nil {
		library.Log.Error(result.Error)
		return user, result.Error
	}

	library.Log.Info("Finish get the user by user account")
	return user, nil
}

func (u *User) UpdateUserErrorById(userId string) error {
	library.Log.Info("Start update the error count by user id")
	library.Log.Debugf("userAccount=%s", userId)

	result := database.PgConn.Model(&models.Users{}).Where("id = ?", userId).Update("error_count + ?", 1)
	if result.Error != nil {
		library.Log.Error(result.Error)
		return result.Error
	}
	return nil
}
