package migrations

import (
	"github.com/GabrielLuizSF/cockroach-db/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(models.Photos{})
}