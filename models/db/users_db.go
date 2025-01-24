package db

import (
	"ADMgmtSystem/database"
	"ADMgmtSystem/library"
	"ADMgmtSystem/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUserByAccount(userAccount string) (models.Users, error)
	UpdateUserErrorById(userId string) error
}
type User struct {
	ID           string
	Account      string
	Password     string
	PasswordSalt string
	Activated    bool
	Role         uint
}

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

func (u *User) InsertUser(db *gorm.DB) error {
	var err error

	if u.Password, err = library.GenerateRandom(10); err != nil {
		return err
	}

	if u.PasswordSalt, err = library.GenerateRandomSymbols(10); err != nil {
		return err
	}

	user := &models.Users{ID: u.ID, Account: u.Account, Password: u.Password, PasswordSalt: u.PasswordSalt, ErrorCount: 0, Activated: u.Activated, Role: u.Role}

	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
