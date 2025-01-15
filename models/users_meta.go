package models

type UsersMeta struct {
	UsersID   string `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"uniqueIndex"`
	Avatar    []byte
	SendEmail bool
}
