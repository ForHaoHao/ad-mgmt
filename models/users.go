package models

type Users struct {
	ID           uint   `gorm:"primaryKey"`
	Account      string `gorm:"size:255;pirmarykey"`
	Password     string `gorm:"size:255"`
	PasswordSalt string `gorm:"size:10"`
	Activated    bool
	Role         uint
}
