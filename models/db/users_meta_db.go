package db

import (
	"ADMgmtSystem/models"

	"gorm.io/gorm"
)

type UsersMeta struct {
	UsersID   string
	Name      string
	Email     string
	Avatar    []byte
	SendEmail bool
}

func (um *UsersMeta) InsertUsersMeta(db *gorm.DB) error {
	userMeta := &models.UsersMeta{UsersID: um.UsersID, Name: um.Name, Email: um.Email, Avatar: um.Avatar, SendEmail: um.SendEmail}

	if err := db.Create(userMeta).Error; err != nil {
		return err
	}

	return nil
}
