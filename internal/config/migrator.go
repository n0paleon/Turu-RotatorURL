package config

import (
	"TuruSMM/internal/entity"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB, config *viper.Viper) {
	if config.GetBool("database.migrate") {
		db.AutoMigrate(
			entity.UrlList{},
			entity.RotatedUrl{},
		)
	}
}
