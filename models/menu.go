package models

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Menu struct {
	ID     uint   `gorm:"primaryKey"`
	KeyId  string `gorm:"size:255"`
	Path   string `gorm:"size:255"`
	CanUse bool
}

func MigrateMenu() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: fmt.Sprintf("%s_create_menu_table", time.Now().Format("20060102")),
		Migrate: func(g *gorm.DB) error {

			if err := g.AutoMigrate(&Menu{}); err != nil {
				return err
			}

			menus := []Menu{}
			for _, menu := range menus {
				if err := g.FirstOrCreate(&menu).Error; err != nil {
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
