package models

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type MenuMeta struct {
	MenuID uint
	Name   string `gorm:"size:255"`
	Data   string `gorm:"size:510"`
	Type   string `gorm:"size:255`
}

func MigrateMenuMeta() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: fmt.Sprintf("%s_create_users_meta_table", time.Now().Format("20060102")),
		Migrate: func(g *gorm.DB) error {

			if err := g.AutoMigrate(&MenuMeta{}); err != nil {
				return err
			}

			menuMetas := []MenuMeta{}
			for _, menuMeta := range menuMetas {
				if err := g.FirstOrCreate(&menuMeta).Error; err != nil {
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
