package models

type UsersMeta struct {
	UsersID    uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100"`
	Email      string `gorm:"uniqueIndex"`
	Age        int32
	Avatar     []byte
	ErrorCount uint
	SendEmail  bool
}
