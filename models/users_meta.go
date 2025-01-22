package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type UsersMeta struct {
	UsersID   string `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"uniqueIndex"`
	Avatar    []byte
	SendEmail bool
}

func MigrateUsersMeta() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250122_create_users_meta_table",
		Migrate: func(g *gorm.DB) error {
			// 创建表结构
			if err := g.AutoMigrate(&UsersMeta{}); err != nil {
				return err
			}
			// 插入默认值
			usersMeta := []UsersMeta{
				{UsersID: "admin", Name: "administrator", Email: "test@yahoo.com.tw", Avatar: nil, SendEmail: true},
			}
			for _, userMeta := range usersMeta {
				if err := g.Create(&userMeta).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(g *gorm.DB) error {
			return g.Migrator().DropTable("users_meta")
		},
	}
}
