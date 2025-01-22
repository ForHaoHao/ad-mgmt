package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Users struct {
	ID           string `gorm:"primaryKey;size:255"`
	Account      string `gorm:"size:255;primarykey"`
	Password     string `gorm:"size:255"`
	PasswordSalt string `gorm:"size:255"`
	ErrorCount   uint   `gorm:"default:0"`
	Activated    bool   `gorm:"default:true"`
	Role         uint   `gorm:"default:1"`
}

func MigrateUsers() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250122_create_users_table",
		Migrate: func(g *gorm.DB) error {
			// 创建表结构
			if err := g.AutoMigrate(&Users{}); err != nil {
				return err
			}
			// 插入默认值
			users := []Users{
				{ID: "admin", Account: "admin", Password: "5a1d689fabfeefb613fbf4399f8795e9b54102bdc2ce85d13483dc3e2b97c003",
					PasswordSalt: `()#"(#!%+%`, ErrorCount: 0, Activated: true, Role: 1},
			}
			for _, user := range users {
				if err := g.Create(&user).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(g *gorm.DB) error {
			return g.Migrator().DropTable("users")
		},
	}
}
