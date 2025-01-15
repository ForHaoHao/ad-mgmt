package models

type Users struct {
	ID           string `gorm:"primaryKey;size:255"`
	Account      string `gorm:"size:255;primarykey"`
	Password     string `gorm:"size:255"`
	PasswordSalt string `gorm:"size:255"`
	ErrorCount   uint   `gorm:"default:0"`
	Activated    bool   `gorm:"default:true"`
	Role         uint   `gorm:"default:1"`
}
